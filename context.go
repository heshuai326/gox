package gox

import (
	"context"

	"github.com/gopub/gox/geo"
)

type ContextKey string

const (
	CKDeviceID    ContextKey = "gox_device_id"
	CKRemoteAddr  ContextKey = "gox_remote_addr"
	CKLocation    ContextKey = "gox_location"
	CKAccessToken ContextKey = "gox_access_token"
	CKUserID      ContextKey = "gox_user_id"
	CKTraceID     ContextKey = "gox_trace_id"
	CKUser        ContextKey = "gox_user"
)

func (k ContextKey) String() string {
	switch k {
	case CKDeviceID:
		return "device_id"
	case CKRemoteAddr:
		return "remote_addr"
	case CKLocation:
		return "location"
	case CKAccessToken:
		return "access_token"
	case CKUserID:
		return "user_id"
	case CKTraceID:
		return "trace_id"
	case CKUser:
		return "user"
	default:
		return "none"
	}
}

func GetUserID(ctx context.Context) int64 {
	id, _ := ctx.Value(CKUserID).(int64)
	return id
}

func WithUserID(ctx context.Context, id int64) context.Context {
	if id == 0 {
		return ctx
	}
	return context.WithValue(ctx, CKUserID, id)
}

func GetUser(ctx context.Context) interface{} {
	return ctx.Value(CKUser)
}

func WithUser(ctx context.Context, u interface{}) context.Context {
	return context.WithValue(ctx, CKUser, u)
}

func GetAccessToken(ctx context.Context) string {
	token, _ := ctx.Value(CKAccessToken).(string)
	return token
}

func WithAccessToken(ctx context.Context, token string) context.Context {
	if len(token) == 0 {
		return ctx
	}
	return context.WithValue(ctx, CKAccessToken, token)
}

func GetRemoteAddr(ctx context.Context) string {
	ip, _ := ctx.Value(CKRemoteAddr).(string)
	return ip
}

func WithRemoteAddr(ctx context.Context, addr string) context.Context {
	if len(addr) == 0 {
		return ctx
	}
	return context.WithValue(ctx, CKRemoteAddr, addr)
}

func GetDeviceID(ctx context.Context) string {
	id, _ := ctx.Value(CKDeviceID).(string)
	return id
}

func WithDeviceID(ctx context.Context, deviceID string) context.Context {
	if len(deviceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, CKDeviceID, deviceID)
}

func GetTraceID(ctx context.Context) string {
	id, _ := ctx.Value(CKTraceID).(string)
	return id
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	if len(traceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, CKTraceID, traceID)
}

func GetLocation(ctx context.Context) *geo.Point {
	id, _ := ctx.Value(CKLocation).(*geo.Point)
	return id
}

func WithLocation(ctx context.Context, location *geo.Point) context.Context {
	if location == nil {
		return ctx
	}
	return context.WithValue(ctx, CKLocation, location)
}

func DetachedContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	if token := GetAccessToken(ctx); len(token) > 0 {
		newCtx = WithAccessToken(newCtx, token)
	}
	if deviceID := GetDeviceID(ctx); len(deviceID) > 0 {
		newCtx = WithDeviceID(newCtx, deviceID)
	}
	if c := GetLocation(ctx); c != nil {
		newCtx = WithLocation(newCtx, c)
	}
	if addr := GetRemoteAddr(ctx); len(addr) > 0 {
		newCtx = WithRemoteAddr(newCtx, addr)
	}
	if traceID := GetTraceID(ctx); len(traceID) > 0 {
		newCtx = WithTraceID(newCtx, traceID)
	}
	if loginID := GetUserID(ctx); loginID > 0 {
		newCtx = WithUserID(newCtx, loginID)
	}
	return newCtx
}
