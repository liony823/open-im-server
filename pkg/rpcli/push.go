package rpcli

import (
	"github.com/liony823/protocol/push"
	"google.golang.org/grpc"
)

func NewPushMsgServiceClient(cc grpc.ClientConnInterface) *PushMsgServiceClient {
	return &PushMsgServiceClient{push.NewPushMsgServiceClient(cc)}
}

type PushMsgServiceClient struct {
	push.PushMsgServiceClient
}
