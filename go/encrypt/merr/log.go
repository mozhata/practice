package merr

import "fmt"

func ErrDetail(err error) string {
	e := WrapErr(err)
	return fmt.Sprintf("E%d: err: %s\nraw err: %s\ncall stack: %s\n",
		e.Code,
		e.Error(),
		e.RawErr(),
		e.CallStack(),
	)
}
