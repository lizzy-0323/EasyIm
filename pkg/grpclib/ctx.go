package grpclib

import (
	"context"
	"go-im/pkg/gerrors"
	"go-im/pkg/logger"
	"strconv"

	"google.golang.org/grpc/metadata"
)

const (
	CtxUserId    = "user_id"
	CtxDeviceId  = "device_id"
	CtxToken     = "token"
	CtxRequestId = "request_id"
)

func ContextWithRequestId(ctx context.Context, requestId int64) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(CtxRequestId, strconv.FormatInt(requestId, 10)))
}

func GetCtxRequestId(ctx context.Context) int64 {
	requestIdStr := Get(ctx, CtxRequestId)
	requestId, err := strconv.ParseInt(requestIdStr, 10, 64)
	if err != nil {
		return 0
	}
	return requestId
}

func Get(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values, ok := md[key]
	if !ok || len(values) == 0 {
		return ""
	}
	return values[0]
}

// GetCtxData 获取ctx的用户数据，依次返回user_id,device_id, 可用于鉴权
func GetCtxData(ctx context.Context) (int64, int64, error) {
	var (
		userId   int64
		deviceId int64
		err      error
	)
	userIdStr := Get(ctx, CtxUserId)
	userId, err = strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, 0, gerrors.ErrUnauthorized
	}

	deviceIdStr := Get(ctx, CtxDeviceId)
	deviceId, err = strconv.ParseInt(deviceIdStr, 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, 0, gerrors.ErrUnauthorized
	}
	return userId, deviceId, nil
}

func NewAndCopyRequestId(ctx context.Context) context.Context {
	newCtx := context.TODO()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return newCtx
	}

	requestIds, ok := md[CtxRequestId]
	if !ok && len(requestIds) == 0 {
		return newCtx
	}
	return metadata.NewOutgoingContext(newCtx, metadata.Pairs(CtxRequestId, requestIds[0]))
}

// GetCtxToken 获取ctx的token
func GetCtxToken(ctx context.Context) string {
	return Get(ctx, CtxToken)
}
