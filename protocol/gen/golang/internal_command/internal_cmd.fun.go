// Code generated by protokitgo. DO NOT EDIT.
// source: internal_command/internal_cmd.proto

package internal_command

import (
	// fixed import by protokitgo
	"context"
	"fmt"
	"math"
	"reflect"
	"strings"

	"bitbucket.org/funplus/ark"
	"bitbucket.org/funplus/sandwich/base/net/link"
	"bitbucket.org/funplus/sandwich/base/serror"
	"bitbucket.org/funplus/sandwich/base/validator"
	"bitbucket.org/funplus/sandwich/client"
	"bitbucket.org/funplus/sandwich/current"
	"bitbucket.org/funplus/sandwich/message"
	"bitbucket.org/funplus/sandwich/metadata"
	"bitbucket.org/funplus/sandwich/protocol/netutils"
	"bitbucket.org/funplus/sandwich/router"
	"google.golang.org/protobuf/proto"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protobufEmpty "google.golang.org/protobuf/types/known/emptypb"
	// dynamic import
)

// Reference imports to suppress errors if they are not otherwise used.
var _ ark.Context
var _ client.Client
var _ router.Router
var _ = current.NewContext(context.Background())
var _ serror.Error
var _ link.Session
var _ message.Message
var _ netutils.RawPacket
var _ ark.Context
var _ context.Context
var _ = fmt.Errorf
var _ = math.Inf
var _ proto.MarshalOptions
var _ reflect.Type
var _ strings.Builder
var _ metadata.MD
var _ protobufEmpty.Empty
var _ validator.Validator

// ServerHandlerInternalCmdForRPC  server handler interface for RPC
type ServerHandlerInternalCmdForRPC interface {
	CmdPing(ctx context.Context, in *netutils.CmdPing) (*netutils.CmdPingAck, error)
	CmdCheckup(ctx context.Context, in *netutils.CmdCheckup) (*netutils.CmdCheckup, error)
	CmdBuildStream(ctx context.Context, in *CmdStream) (*CmdStream, error)
}

// ServerHandlerInternalCmdForActor  server handler interface for Actor system
type ServerHandlerInternalCmdForActor interface {
}

// ServerHandlerInternalCmd server handler interface
type ServerHandlerInternalCmd interface {
	ServerHandlerInternalCmdForRPC
	ServerHandlerInternalCmdForActor
}

// UnimplementedServerHandlerInternalCmdForRPC  can be embedded to have forward compatible implementations.
type UnimplementedServerHandlerInternalCmdForRPC struct{}

func (*UnimplementedServerHandlerInternalCmd) CmdPing(ctx context.Context, in *netutils.CmdPing) (*netutils.CmdPingAck, error) {
	return nil, serror.NewProtoEnum(netutils.ErrorCode_NotImplement)
}
func (*UnimplementedServerHandlerInternalCmd) CmdCheckup(ctx context.Context, in *netutils.CmdCheckup) (*netutils.CmdCheckup, error) {
	return nil, serror.NewProtoEnum(netutils.ErrorCode_NotImplement)
}
func (*UnimplementedServerHandlerInternalCmd) CmdBuildStream(ctx context.Context, in *CmdStream) (*CmdStream, error) {
	return nil, serror.NewProtoEnum(netutils.ErrorCode_NotImplement)
}

// UnimplementedServerHandlerInternalCmdForActor  can be embedded to have forward compatible implementations.
type UnimplementedServerHandlerInternalCmdForActor struct {
}

// UnimplementedServerHandlerInternalCmd  can be embedded to have forward compatible implementations.
type UnimplementedServerHandlerInternalCmd struct {
	UnimplementedServerHandlerInternalCmdForRPC
	UnimplementedServerHandlerInternalCmdForActor
}

// gen server side proxy
type proxyServerHandlerInternalCmdForRPC struct {
	handler ServerHandlerInternalCmdForRPC
}

// gen server side proxy
type proxyServerHandlerInternalCmdForActor struct {
	handler ServerHandlerInternalCmdForActor
}

func (h *proxyServerHandlerInternalCmdForRPC) proxyCmdPing(ctx context.Context, in interface{}) (interface{}, error) {
	if req, ok := in.(*netutils.CmdPing); ok {
		return h.handler.CmdPing(ctx, req)
	} else {
		return nil, serror.NewProtoEnum(netutils.ErrorCode_MessageCastError)
	}
}
func (h *proxyServerHandlerInternalCmdForRPC) proxyCmdCheckup(ctx context.Context, in interface{}) (interface{}, error) {
	if req, ok := in.(*netutils.CmdCheckup); ok {
		return h.handler.CmdCheckup(ctx, req)
	} else {
		return nil, serror.NewProtoEnum(netutils.ErrorCode_MessageCastError)
	}
}
func (h *proxyServerHandlerInternalCmdForRPC) proxyCmdBuildStream(ctx context.Context, in interface{}) (interface{}, error) {
	if req, ok := in.(*CmdStream); ok {
		return h.handler.CmdBuildStream(ctx, req)
	} else {
		return nil, serror.NewProtoEnum(netutils.ErrorCode_MessageCastError)
	}
}

// tcp handler
func RegisterServerHandlerInternalCmdForRPCForCommonRouter(r router.Router, s ServerHandlerInternalCmdForRPC) {
	h := &proxyServerHandlerInternalCmdForRPC{handler: s}
	{
		rsp := new(netutils.CmdPingAck)
		req := new(netutils.CmdPing)
		r.SetMessageHandler(req, h.proxyCmdPing)
		// uri mapping
		if s, ok := r.(interface {
			SetURIMapping(interface{}, interface{})
		}); ok {
			s.SetURIMapping(req, rsp)
		}
	}
	{
		rsp := new(netutils.CmdCheckup)
		req := new(netutils.CmdCheckup)
		r.SetMessageHandler(req, h.proxyCmdCheckup)
		// uri mapping
		if s, ok := r.(interface {
			SetURIMapping(interface{}, interface{})
		}); ok {
			s.SetURIMapping(req, rsp)
		}
	}
	{
		rsp := new(CmdStream)
		req := new(CmdStream)
		r.SetMessageHandler(req, h.proxyCmdBuildStream)
		// uri mapping
		if s, ok := r.(interface {
			SetURIMapping(interface{}, interface{})
		}); ok {
			s.SetURIMapping(req, rsp)
		}
	}
}

// tcp handler
func RegisterServerHandlerInternalCmdForActorForCommonRouter(r router.Router, s ServerHandlerInternalCmdForActor) {
}

func RegisterServerHandlerInternalCmdForCommonRouter(r router.Router, s ServerHandlerInternalCmd) {
	RegisterServerHandlerInternalCmdForRPCForCommonRouter(r, s)
	RegisterServerHandlerInternalCmdForActorForCommonRouter(r, s)
}

// http handler
func RegisterServerHandlerInternalCmdForRPCForArk(r ark.Router, s ServerHandlerInternalCmdForRPC) {
	h := &proxyServerHandlerInternalCmdForRPC{handler: s}
	{
		t := new(netutils.CmdPing)
		m := h.proxyCmdPing
		err := router.RegisterMessageUnder(r, t, m)

		if err != nil {
			panic(fmt.Errorf("service register failed with RegisterMessageUnder, with message:%v, err:%w", t, err))
		}
		// auto generate by ProtoKitGo
		path := strings.TrimPrefix("/auto/netutils.CmdPing", r.BasePath())
		r.GetPost(path, router.ArkCommonHandle(t, m))
	}
	{
		t := new(netutils.CmdCheckup)
		m := h.proxyCmdCheckup
		err := router.RegisterMessageUnder(r, t, m)

		if err != nil {
			panic(fmt.Errorf("service register failed with RegisterMessageUnder, with message:%v, err:%w", t, err))
		}
		// auto generate by ProtoKitGo
		path := strings.TrimPrefix("/auto/netutils.CmdCheckup", r.BasePath())
		r.GetPost(path, router.ArkCommonHandle(t, m))
	}
	{
		t := new(CmdStream)
		m := h.proxyCmdBuildStream
		err := router.RegisterMessageUnder(r, t, m)

		if err != nil {
			panic(fmt.Errorf("service register failed with RegisterMessageUnder, with message:%v, err:%w", t, err))
		}
		// auto generate by ProtoKitGo
		path := strings.TrimPrefix("/auto/internal_command.CmdStream", r.BasePath())
		r.GetPost(path, router.ArkCommonHandle(t, m))
	}
}

// http handler
func RegisterServerHandlerInternalCmdForActorForArk(r ark.Router, s ServerHandlerInternalCmdForActor) {
}

func RegisterServerHandlerInternalCmdForArk(r ark.Router, s ServerHandlerInternalCmd) {
	RegisterServerHandlerInternalCmdForRPCForArk(r, s)
	RegisterServerHandlerInternalCmdForActorForArk(r, s)
}

func RegisterServerHandlerInternalCmdForRPC(dst interface{}, handler ServerHandlerInternalCmdForRPC) {
	found := false
	if commonRouter, ok := dst.(router.Router); ok {
		found = true
		RegisterServerHandlerInternalCmdForRPCForCommonRouter(commonRouter, handler)
	}
	if arkRouter, ok := dst.(ark.Router); ok {
		found = true
		RegisterServerHandlerInternalCmdForRPCForArk(arkRouter, handler)
	}
	if !found {
		panic(fmt.Sprintf("got invalid router type:%v", reflect.TypeOf(dst)))
	}
}

func RegisterServerHandlerInternalCmdForActor(dst interface{}, handler ServerHandlerInternalCmdForActor) {
	found := false
	if commonRouter, ok := dst.(router.Router); ok {
		found = true
		RegisterServerHandlerInternalCmdForActorForCommonRouter(commonRouter, handler)
	}
	if arkRouter, ok := dst.(ark.Router); ok {
		found = true
		RegisterServerHandlerInternalCmdForActorForArk(arkRouter, handler)
	}
	if !found {
		panic(fmt.Sprintf("got invalid router type:%v", reflect.TypeOf(dst)))
	}
}

func RegisterServerHandlerInternalCmd(dst interface{}, handler ServerHandlerInternalCmd) {
	found := false
	if commonRouter, ok := dst.(router.Router); ok {
		found = true
		RegisterServerHandlerInternalCmdForCommonRouter(commonRouter, handler)
	}
	if arkRouter, ok := dst.(ark.Router); ok {
		found = true
		RegisterServerHandlerInternalCmdForArk(arkRouter, handler)
	}
	if !found {
		panic(fmt.Sprintf("got invalid router type:%v", reflect.TypeOf(dst)))
	}
}

// Deprecated: use ServerHandlerInternalCmd
type InternalCmdService = ServerHandlerInternalCmd

// Deprecated: use RegisterServerHandlerInternalCmdForArk or RegisterServerHandlerInternalCmd
var RegisterInternalCmdServiceHttpHandler = RegisterServerHandlerInternalCmdForArk

// Deprecated: use RegisterServerHandlerInternalCmdForCommonRouter or RegisterServerHandlerInternalCmd
var RegisterInternalCmdServiceTcpHandler = RegisterServerHandlerInternalCmdForCommonRouter

// rpcProxy将服务内的rpc server转发到Client指定的服务，如Python RPC Server
type rpcProxyServerHandlerInternalCmd struct {
	rpcClient RpcClientInternalCmd
	opts      []client.CallOption
}

func NewServerHandlerInternalCmdWithClient(cc client.Client, opts ...client.CallOption) ServerHandlerInternalCmd {
	return &rpcProxyServerHandlerInternalCmd{rpcClient: NewRpcClientInternalCmd(cc), opts: opts}
}
func (r *rpcProxyServerHandlerInternalCmd) CmdPing(ctx context.Context, in *netutils.CmdPing) (*netutils.CmdPingAck, error) {
	return r.rpcClient.CmdPing(ctx, in, r.opts...)
}
func (r *rpcProxyServerHandlerInternalCmd) CmdCheckup(ctx context.Context, in *netutils.CmdCheckup) (*netutils.CmdCheckup, error) {
	return r.rpcClient.CmdCheckup(ctx, in, r.opts...)
}
func (r *rpcProxyServerHandlerInternalCmd) CmdBuildStream(ctx context.Context, in *CmdStream) (*CmdStream, error) {
	return r.rpcClient.CmdBuildStream(ctx, in, r.opts...)
}
