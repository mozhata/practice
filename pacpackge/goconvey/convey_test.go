package goconvey

/*
命令行下使用 go test -v
使用webUI需要事先执行 `go get github.com/smartystreets/goconvey`
然后在测试目录下执行 `goconvey`

*/
import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("testing Add: add two number", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSubstract(t *testing.T) {
	Convey("testing Substruct: 4 -6", t, func() {
		So(Substract(4, 6), ShouldEqual, -2)
	})
}

func TestMultiply(t *testing.T) {
	Convey("testing Multiply: 4 * 5", t, func() {
		So(Multiply(4, 5), ShouldEqual, 20)
	})
}

func TestDivision(t *testing.T) {
	Convey("test Division", t, func() {
		Convey("divisor is 0", func() {
			_, err := Division(2, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("divisor is not 0", func() {
			num, err := Division(5, 3)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 1)
		})
	})
}
