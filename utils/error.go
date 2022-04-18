package utils

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrNotImplemented      = errors.New("not implemented")
	ErrNilParameter        = errors.New("nil parameter")
	ErrNilOrWrongParameter = errors.New("nil or wrong parameter")
	ErrWrongParameter      = errors.New("wrong parameter")
	ErrShortRead           = errors.New("short read")
	ErrInvalidData         = errors.New("invalid data")
	ErrHandled             = errors.New("handled")
	ErrFailed              = errors.New("failed")
)

//没啥特殊的
type NumErr struct {
	N      int
	Prefix string
}

func (ne NumErr) Error() string {

	return ne.Prefix + strconv.Itoa(ne.N)
}

//就是带个buffer的普通ErrInErr，没啥特殊的
type ErrFirstBuffer struct {
	Err   error
	First *bytes.Buffer
}

func (ef ErrFirstBuffer) Unwarp() error {

	return ef.Err
}

func (ef ErrFirstBuffer) Error() string {

	return ef.Err.Error()
}

// ErrInErr 很适合一个err包含另一个err，并且提供附带数据的情况.
type ErrInErr struct {
	ErrDesc   string
	ErrDetail error
	Data      any
}

func (e ErrInErr) Error() string {
	return e.String()
}

func (e ErrInErr) Unwarp() error {
	return e.ErrDetail
}

func (e ErrInErr) Is(err error) bool {
	if e.ErrDetail == err {
		return true
	} else if errors.Is(e.ErrDetail, err) {
		return true
	}
	return false
}

func (e ErrInErr) String() string {

	if e.Data != nil {

		if e.ErrDetail != nil {
			return fmt.Sprintf("%s : %s, Data: %v", e.ErrDesc, e.ErrDetail.Error(), e.Data)

		}

		return fmt.Sprintf("%s , Data: %v", e.ErrDesc, e.Data)

	}
	if e.ErrDetail != nil {
		return fmt.Sprintf("%s : %s", e.ErrDesc, e.ErrDetail.Error())

	}
	return e.ErrDesc

}
