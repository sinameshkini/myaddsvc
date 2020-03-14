package addendpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"myaddsvc/pkg/addservice"
)

type Set struct {
	SumEndpoint		endpoint.Endpoint
	ConcatEndpoint	endpoint.Endpoint
}

func New(svc addservice.Service) Set {
	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = MakeSumEndpoint(svc)
	//	TODO: another some code...
	}

	var concatEndpoint endpoint.Endpoint
	{
		concatEndpoint = MakeConcatEndpoint(svc)
	//	TODO: same up...
	}
	return Set{
		SumEndpoint: sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}
}

func MakeSumEndpoint(s addservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SumRequest)
		v, err := s.Sum(ctx, req.A, req.B)
		return SumResponse{v, err}, nil
	}
}

func MakeConcatEndpoint(s addservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ConcatRequest)
		v, err := s.Concat(ctx, req.A, req.B)
		return ConcatResponse{v, err}, nil
	}
}

// TODO: What is this???
// compile time assertions for our response types implementing endpoint.Failer.
//var (
//	_ endpoint.Failer = SumResponse{}
//	_ endpoint.Failer = ConcatResponse{}
//)


type SumRequest struct {
	A, B int
}

type SumResponse struct {
	V	int		`json:"v"`
	Err error 	`json:"err"`
}

func (r SumResponse) Failed() error {return r.Err}

type ConcatRequest struct {
	A, B string
}

type ConcatResponse struct {
	V	string	`json:"v"`
	Err error	`json:"err"`
}

func (r ConcatResponse) Failed() error {return r.Err}
