package handler

import (
	"context"
	"go-micro.dev/v4/logger"

	pb "github.com/go-micro/demo/helloservice/proto"
)

type Helloservice struct{}

func (e *Helloservice) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Helloservice.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
