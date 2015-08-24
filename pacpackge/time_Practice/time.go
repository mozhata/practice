package main

import (
	"fmt"
	"time"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	// basic()
	format()
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
	// 按照基本现有格式format
	// func (t Time) Format(layout string) string
	P(t.Format(time.RFC3339))
	//func Parse(layout, value string) (Time, error)
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	P(t1)
}
