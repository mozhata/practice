package cryp

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"math/big"
	"sync"

	"github.com/mozhata/merr"

	"practice/go/util"
)

const base64Encoder = "zxcvbnmqwertyuiopasdfghjklASDFGHJKLZXCVBNMQWERTYUIOP-+1234567890"

var Encoding64 = base64.NewEncoding(base64Encoder)

var count = counter{
	num: 1000,
}

type counter struct {
	sync.Mutex
	num uint32
}

type Payload struct {
	IssueTime   uint32 // 时间戳，秒
	TTL         uint16 // 单位为分钟
	AccountId   string // 长度不定
	LoginSource string // 3B
}

func TryEncrypt() {
	pls := []*Payload{
		&Payload{
			// IssueTime:   uint32(time.Now().Unix()),
			IssueTime:   uint32(1517916231),
			TTL:         uint16(300),
			AccountId:   "1234567890qwertyuiopasdfghjklzxcvbnmm",
			LoginSource: "web",
		},
	}
	util.Debug("pls: %s", util.MarshalJSONOrDie(pls))

	str, err := Encrypt(pls[0])
	CheckErr(err)
	util.Debug("encrypted: %s", str)
	p, err := Decrypt(str)
	CheckErr(err)
	util.Debug("decrypt payload: %s", util.MarshalJSONOrDie(p))
}

func Encrypt(tk *Payload) (string, error) {

	count.Lock()
	seq := count.num
	count.num += 1
	count.Unlock()
	data, err := tk.encryptV1(seq)
	if err != nil {
		return "", err
	}
	// token := base64.URLEncoding.EncodeToString(data)
	token := Encoding64.EncodeToString(data)
	return token, nil
}

func (t *Payload) encryptV1(seq uint32) ([]byte, error) {

	var datas = make([]byte, 6)
	var sigs = make([]byte, 1)
	var seqs = make([]byte, 4)

	// SIG = VERSION(1B) + SEQ(3B)
	sigs[0] = 0x01
	binary.LittleEndian.PutUint32(seqs[0:4], seq)
	sigs = append(sigs, seqs[0:3]...)

	fixedLongSource := packLeadingZero16([]byte(t.LoginSource))

	// PAYLOAD = ISSUETIME(4B) + TTL(2B) + ACCOUNTID(32B) + SOURCE(>0B)
	binary.LittleEndian.PutUint32(datas[0:4], t.IssueTime)
	binary.LittleEndian.PutUint16(datas[4:6], uint16(t.TTL))
	datas = append(datas, fixedLongSource...)
	datas = append(datas, []byte(t.AccountId)...)

	// 对SIG+PAYLOAD做签名，SIGN = SIGN_R(32B)+SIGN_S(32B)
	h := md5.New()
	h.Write(append(sigs, datas...))
	sign := h.Sum(nil)[:16]

	// TOKEN = SIG{VERSION(1B)+SEQ(3B)} + SIGN{SIGN_R(32B)+SIGN_S(32B)} + PAYLOAD{ISSUETIME(4B) + TTL(2B) + ACCOUNTID(32B) + SOURCE(>0B)}
	sigs = append(sigs, sign...)
	// sigs = append(sigs, packLeadingZero32(s.Bytes())...)
	sigs = append(sigs, datas...)

	return sigs, nil
}

func Decrypt(token string) (*Payload, error) {
	// data, err := base64.URLEncoding.DecodeString(token)
	data, err := Encoding64.DecodeString(token)
	if err != nil {
		return nil, errors.New("invalid base64 token")
	}
	pl := &Payload{}
	err = pl.decryptV1(data)
	if err != nil {
		return nil, err
	}
	return pl, nil
}

func (t *Payload) decryptV1(data []byte) error {

	// SIG = VERSION(1B) + SEQ(3B)
	sigs := data[0:4]

	// PAYLOAD = ISSUETIME(4B) + TTL(2B) + ACCOUNTID(>0B) + SOURCE(>0B)
	payload := data[20:]

	t.IssueTime = uint32(binary.LittleEndian.Uint32(payload[:4]))
	t.TTL = uint16(binary.LittleEndian.Uint16(payload[4:6]))
	t.LoginSource = string(unpackLeadingZero(payload[6:22]))
	t.AccountId = string(payload[22:])

	// SIGN = SIGN(32B)
	s := big.NewInt(0)
	s = s.SetBytes(unpackLeadingZero(data[4:20]))
	sign := s.Bytes()

	// 签名校验
	h := md5.New()
	h.Write(append(sigs, payload...))
	hashed := h.Sum(nil)[:16]
	if !bytes.Equal(sign, hashed) {
		return merr.WrapErr(nil, "sign verify failed")
	}

	return nil
}

func packLeadingZero16(bs []byte) []byte {
	if len(bs) >= 16 {
		return bs
	}

	packBytes := make([]byte, 16-len(bs))

	return append(packBytes, bs...)
}

func unpackLeadingZero(bs []byte) []byte {
	i := 0
	for i < len(bs) && bs[i] == 0 {
		i++
	}

	return bs[i:]
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
