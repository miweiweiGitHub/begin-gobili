// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v0/Broadcast.proto

/*
Package v0 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v0/Broadcast.proto
*/
package v0

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ===================
// Broadcast Interface
// ===================

type BroadcastRPCClient interface {
	// * 发送 NOTICE_MSG 广播接口
	// 如果有定义默认样式的, 且如果样式字段没有传, 会被补全; 最终输出字段格式参考 http://info.bilibili.co/pages/viewpage.action?pageId=3693681#id-%E5%BC%B9%E5%B9%95%E9%80%9A%E7%9F%A5%E5%8D%8F%E8%AE%AE-%E6%96%B0%E7%89%88%E5%BC%B9%E5%B9%95%E9%80%9A%E7%9F%A5%E5%8D%8F%E8%AE%AE
	SendNoticeMsg(ctx context.Context, req *BroadcastSendNoticeMsgReq, opts ...liverpc.CallOption) (resp *BroadcastSendNoticeMsgResp, err error)
}

// =========================
// Broadcast Live Rpc Client
// =========================

type broadcastRPCClient struct {
	client *liverpc.Client
}

// NewBroadcastRPCClient creates a client that implements the BroadcastRPCClient interface.
func NewBroadcastRPCClient(client *liverpc.Client) BroadcastRPCClient {
	return &broadcastRPCClient{
		client: client,
	}
}

func (c *broadcastRPCClient) SendNoticeMsg(ctx context.Context, in *BroadcastSendNoticeMsgReq, opts ...liverpc.CallOption) (*BroadcastSendNoticeMsgResp, error) {
	out := new(BroadcastSendNoticeMsgResp)
	err := doRPCRequest(ctx, c.client, 0, "Broadcast.send_notice_msg", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRPCRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message, opts []liverpc.CallOption) (err error) {
	err = client.Call(ctx, version, method, in, out, opts...)
	return
}
