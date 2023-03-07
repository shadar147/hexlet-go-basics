package ex30

import (
	"errors"
	"fmt"
)

func Test() {
	fmt.Println(GetErrorMsg(ErrBadConnection))           // "bad connection"
	fmt.Println(GetErrorMsg(ErrBadRequest))              // "bad request"
	fmt.Println(GetErrorMsg(NonCriticalError{}))         // ""
	fmt.Println(GetErrorMsg(errors.New("random error"))) // "unknown error"
}
