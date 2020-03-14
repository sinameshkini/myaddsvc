package addservice

import (
	"context"
	"errors"
)

type Service interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

var (
	ErrTwoZeroes = errors.New("can't sum two zeroes")
	ErrIntOverflow = errors.New("integer overflow")
	ErrMaxSizeExceeded = errors.New("result exceed max size")
)

type basicService struct {}

func NewBasicService() Service {
	return basicService{}
}

func New() Service { // TODO: add required param for midllewares to method input
	var svc Service
	{
		svc = NewBasicService()
	//	TODO: add middlewares
	}
	return svc
}

const (
	intMax = 1<<31 - 1
	intMin = -(intMax + 1)
	maxLen = 10
)

func (s basicService) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (intMax-b)) || (b < 0 && a < (intMin-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

func (s basicService) Concat(_ context.Context, a, b string) (string, error) {
	if len(a)+len(b) > maxLen {
		return "", ErrMaxSizeExceeded
	}
	return a + b, nil
}
