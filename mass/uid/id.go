package uid

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var (
	idChan = func() <-chan string {
		c := make(chan string)
		offset := time.Date(2017, time.April, 18, 0, 0, 0, 0, time.UTC)
		go func() {
			seq := int64(0)
			previousId := int64(0)
			for {
				inc, err := rand.Int(rand.Reader, big.NewInt(499))
				if err != nil {
					panic(err)
				}
				if seq > 9500 {
					seq = 0
				}
				seq += inc.Int64()
				// eg: 150706370050000 + seq, and seq < 10000, will failed 200years later
				id := int64(time.Now().Sub(offset).Seconds()*float64(time.Microsecond)*10)*10000 + seq
				if id < previousId {
					id += 1
				}
				previousId = id
				c <- base62encode(id)
			}
		}()
		return c
	}()

	idChars    = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	idCharsLen = int64(len(idChars))
)

func base62encode(n int64) string {
	if n < 0 {
		panic(fmt.Sprintf("%v", n))
	}
	if n < idCharsLen {
		return string(idChars[n])
	}
	ret := ""
	for n != 0 {
		ret = string(idChars[n%idCharsLen]) + ret
		n = n / idCharsLen
	}
	return ret
}

func NewUniqueId() string {
	return <-idChan
}

func TryUniqueID() {
	offset := time.Date(2017, time.April, 1, 0, 0, 0, 0, time.UTC)
	second := int64(time.Now().Sub(offset).Seconds())
	id1 := int64(time.Now().Sub(offset)) / int64(time.Microsecond) / 100 * 10000
	id2 := int64(time.Now().Sub(offset).Nanoseconds() / 10)
	id3 := int64(time.Now().Sub(offset).Seconds() * float64(time.Second) / 10)
	id4 := int64(time.Now().Sub(offset).Seconds()*float64(time.Microsecond)*10) * 10000
	fmt.Printf("sec:\t%-25d\nid1:\t%-25d\nid2:\t%-25d\nid3:\t%-25d\nid4:\t%-25d\n",
		second, id1, id2, id3, id4)
	for i := 0; i < 5; i++ {
		j := int64(60 + i)
		fmt.Printf("the %dst id: %s\tbase62encode: %s\tmode: %d\n",
			j, NewUniqueId(), base62encode(j), j%idCharsLen)
	}
}
