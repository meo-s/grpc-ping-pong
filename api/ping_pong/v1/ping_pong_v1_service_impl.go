package v1

import (
	context "context"

	pb_empty "github.com/golang/protobuf/ptypes/empty"
	pb_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

type PingPongRpcServerImpl struct {
	UnimplementedPingPongRpcServer
}

func (s *PingPongRpcServerImpl) Call(cx context.Context, _ *pb_empty.Empty) (*pb_wrappers.StringValue, error) {
	resp := &pb_wrappers.StringValue{
		Value: "pong-!",
	}
	return resp, nil
}
