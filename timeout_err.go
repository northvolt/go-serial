package serial

import (
	"fmt"
	"time"
)

type timeoutErr struct{
	msg string
}

func newReadTimeoutErr(timeout time.Duration) error {
	return &timeoutErr{
		msg: fmt.Sprintf("timeout (%s) reading from port", timeout),
	}
}

func newWriteTimeoutErr(timeout time.Duration) error {
	return &timeoutErr{
		msg: fmt.Sprintf("timeout (%s) writing to port", timeout),
	}
}

func (t *timeoutErr) Error() string {
	return t.msg
}

func (t *timeoutErr) Timeout() bool {
	return true
}
