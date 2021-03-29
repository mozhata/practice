package mist

/*
* 薄雾算法
*
* 1      2                                                     48         56       64
* +------+-----------------------------------------------------+----------+----------+
* retain | increas                                             | saltA    | saltB    |
* +------+-----------------------------------------------------+----------+----------+
* 0      | 0000000000 0000000000 0000000000 0000000000 0000000 | 00000000 | 00000000 |
* +------+-----------------------------------------------------+------------+--------+
*
* 0. 最高位，占 1 位，保持为 0，使得值永远为正数；
* 1. 自增数，占 47 位，自增数在高位能保证结果值呈递增态势，遂低位可以为所欲为；
* 2. 随机因子一，占 8 位，上限数值 255，使结果值不可预测；
* 3. 随机因子二，占 8 位，上限数值 255，使结果值不可预测；
*
* 编号上限为百万亿级，上限值计算为 140737488355327 即 int64(1 << 47 - 1)，假设每天取值 10 亿，能使用 385+ 年
 */

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

const (
	saltShift     = 8
	increaseShift = saltShift + 8
)

type Mist struct {
	sync.Mutex
	increas int64
	saltA   int64
	saltB   int64
}

func NewMist() *Mist {
	return &Mist{}
}

func (c *Mist) Generate() int64 {
	c.Lock()
	c.increas++
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltA = randA.Int64()
	// 也可以直接用math/rand 里面的随机数
	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltB = randB.Int64()
	// increas 位移16位, saltA位移8位, saltB不位移, 通过"|"运算做merge
	mist := int64((c.increas << increaseShift) | (c.saltA << saltShift) | c.saltB)
	fmt.Printf("increas: %d, in bytes: %064b, .increas << %d: %064b\n", c.increas, c.increas, increaseShift, c.increas<<increaseShift)
	fmt.Printf("c.saltA: %d, in bytes: %064b, .saltA << saltShift: %064b\n", c.saltA, c.saltA, c.saltA<<saltShift)
	fmt.Printf("c.saltB: %d, in bytes: %064b, \nmist: %d, inbyte: %064b\n", c.saltB, c.saltB, mist, mist)

	c.Unlock()
	return mist
}
