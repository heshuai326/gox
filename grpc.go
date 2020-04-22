package gox

import (
	"context"
	"fmt"
	"time"

	"github.com/gopub/gox/geo"

	"github.com/gopub/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

var KeepAliveClientParames = keepalive.ClientParameters{
	Time:                time.Second * 10,
	Timeout:             time.Second * 5,
	PermitWithoutStream: true,
}

func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return invoker(outgoingContext(ctx), method, req, reply, cc, opts...)
}

func StreamClientInterceptor(ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(outgoingContext(ctx), desc, cc, method, opts...)
}

func outgoingContext(ctx context.Context) context.Context {
	md := metadata.New(nil)

	if deviceID := GetDeviceID(ctx); len(deviceID) > 0 {
		md.Set(CKDeviceID.String(), deviceID)
	}

	if coordinate := GetLocation(ctx); coordinate != nil {
		md.Set(CKLocation.String(), fmt.Sprintf("%f,%f", coordinate.Y, coordinate.X))
	}

	if token := GetAccessToken(ctx); len(token) > 0 {
		md.Set(CKAccessToken.String(), token)
	}

	if ClientMetadataFromContext != nil {
		m := ClientMetadataFromContext(ctx)
		for k, v := range m {
			md.Append(k, v...)
		}
	}

	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

var ClientMetadataFromContext func(ctx context.Context) metadata.MD

func NewGRPCClient(url string) *grpc.ClientConn {
	conn, err := grpc.Dial(url,
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second*2),
		grpc.WithKeepaliveParams(KeepAliveClientParames),
		grpc.WithUnaryInterceptor(UnaryClientInterceptor),
		grpc.WithStreamInterceptor(StreamClientInterceptor))
	if err != nil {
		log.Panicf("Dial failed: %s %v", url, err)
	}
	return conn
}

var TokenAuthenticator func(ctx context.Context, token string) context.Context

type serverStreamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *serverStreamWrapper) Context() context.Context {
	return s.ctx
}

func handleIncomingMetadata(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	coordinateVal := md.Get(CKLocation.String())
	if len(coordinateVal) > 0 && len(coordinateVal[0]) > 0 {
		var lat, lng float64
		if _, err := fmt.Sscanf(coordinateVal[0], "%f,%f", &lat, &lng); err == nil {
			coordinate := &geo.Point{
				X: lng,
				Y: lat,
			}
			ctx = WithLocation(ctx, coordinate)
		}
	}

	deviceID := md.Get(CKDeviceID.String())
	if len(deviceID) > 0 && len(deviceID[0]) > 0 {
		ctx = WithDeviceID(ctx, deviceID[0])
	}

	accessToken := md.Get(CKAccessToken.String())
	if len(accessToken) > 0 && len(accessToken[0]) > 0 {
		ctx = WithAccessToken(ctx, accessToken[0])
		if TokenAuthenticator != nil {
			ctx = TokenAuthenticator(ctx, accessToken[0])
		}
	}

	if WithMetadata != nil {
		ctx = WithMetadata(md)
	}

	if id := GetUserID(ctx); id > 0 {
		logger := log.With("login", id)
		ctx = log.BuildContext(ctx, logger)
	}
	return ctx
}

var WithMetadata func(md metadata.MD) context.Context

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx = handleIncomingMetadata(ctx)
	return handler(ctx, req)
}

func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := handleIncomingMetadata(ss.Context())
	ss = &serverStreamWrapper{
		ServerStream: ss,
		ctx:          ctx,
	}
	return handler(srv, ss)
}

var (
	KeepAliveServerParams = keepalive.ServerParameters{
		MaxConnectionIdle: time.Minute * 5,
		Time:              time.Second * 30,
		Timeout:           time.Second * 10,
	}

	EnforcementPolicy = keepalive.EnforcementPolicy{
		MinTime:             time.Second * 10,
		PermitWithoutStream: true,
	}
)

func NewGRPCServer(opt ...grpc.ServerOption) *grpc.Server {
	opt = append(opt, grpc.UnaryInterceptor(UnaryServerInterceptor),
		grpc.StreamInterceptor(StreamServerInterceptor),
		grpc.KeepaliveParams(KeepAliveServerParams),
		grpc.KeepaliveEnforcementPolicy(EnforcementPolicy),
		grpc.StatsHandler(&StatsHandler{}))
	return grpc.NewServer(opt...)
}

type StatsHandler struct {
}

// TagConn prepares context for HandleConn and the suffix handlers
func (s *StatsHandler) TagConn(ctx context.Context, tag *stats.ConnTagInfo) context.Context {
	remoteAddr := tag.RemoteAddr.String()
	ctx = WithRemoteAddr(ctx, remoteAddr)
	return ctx
}

// HandleConn is called when new conn established or ended
func (s *StatsHandler) HandleConn(ctx context.Context, st stats.ConnStats) {}

// TagRPC return context which will affect context in HandleRPC and the suffix handlers
func (s *StatsHandler) TagRPC(ctx context.Context, tag *stats.RPCTagInfo) context.Context {
	return ctx
}

// HandleRPC allows prefix process before UnaryServerInterceptor/StreamServerInterceptor
// it's called for each RPC
func (s *StatsHandler) HandleRPC(ctx context.Context, st stats.RPCStats) {
}
