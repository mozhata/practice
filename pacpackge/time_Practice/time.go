package main

import (
	"fmt"
	"time"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	// basic()
	// format()
	// duration()
	// tick()
	ticker()
	// Ttimer()
	// getBeforeTime()
}
func basic() {
	now := time.Now()
	P("now: ", now)
	then := time.Date(2014, 11, 22, 88, 39, 25, 651387237, time.UTC)
	// time.Date(year, month, day, hour, min, sec, nsec, loc)
	P("time: ", then)

	// 是否是time类型的零值
	P(then.IsZero())

	// func (t Time) ISOWeek() (year, week int)
	// 返回ISO标准:年,和第几个week
	P(then.ISOWeek())

	P(then.Year())
	P("month: ", then.Month())
	P(then.Day())
	P(then.Hour())
	P(then.Nanosecond())
	P(then.Location())

	P("weekday: ", then.Weekday())

	// true or false
	P(then.Before(now))
	P(then.After(now))
	P(then.Equal(now))

	diff := now.Sub(then)
	idff := then.Sub(now)
	P(diff, "\t\t", idff)

	// 时间间隔以hour,minutes,second等方式表示
	P(diff.Hours())
	P(diff.Minutes())
	P(diff.Seconds())
	P(diff.Nanoseconds())
	P(diff.String())

	// 加减时间
	P(then.Add(diff))
	P(then.Add(-diff))
}
func format() {
	t := time.Now()
	P(t, t.UTC())
	// 按照基本现有格式format
	// func (t Time) Format(layout string) string
	P(t.Format(time.RFC3339))
	//func Parse(layout, value string) (Time, error)
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	P(t1)
}
func duration() {
	second := time.Second
	P(int64(second / time.Microsecond))

	seconds := 10
	P(time.Duration(seconds) * time.Second)

	t0 := time.Now()
	time.Sleep(1 * time.Second)
	t1 := time.Now()
	P("sleep took times: ", t1.Sub(t0))
}

// 每隔 time.Tick秒,做点什么
// 循环执行,不停止
func tick() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		P(now, "do sth..")
	}
}

// type Ticker
// do something repeatedly at regular intervals

// 	func NewTicker(d Duration) *Ticker
// 	func (t *Ticker) Stop()
func ticker() {
	P("now: ", time.Now())
	ticker := time.NewTicker(time.Millisecond * 5000)
	go func() {
		for t := range ticker.C {
			P("tick at: ", t)
		}
	}()
	time.Sleep(time.Millisecond * 15000)
	ticker.Stop()
	P("ticker stopped")
}

func Ttimer() {
	done := make(chan bool, 2)

	// 代表在将来发生的单一事件,需要设定等待时间
	// <The <-timer1.C blocks on the timer’s channel C
	// until it sends a value indicating that the timer expired.
	timer1 := time.NewTimer(time.Second * 4)

	go func() {
		<-timer1.C
		P("timer1 expired")
		done <- true
	}()
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		P("timer2 expired")
		done <- true
	}()

	//you can cancel the timer before it expires.
	if stop2 := timer2.Stop(); stop2 {
		P("timer2 stopped")
	}

	// for i := 0; i < 2; i++ {
	<-done
	// }
}
func getBeforeTime() {
	now := time.Now()
	// Add 参数可为负数
	before := now.Add(time.Hour * -24)
	P("now: ", now)
	P("before: ", before)

}
