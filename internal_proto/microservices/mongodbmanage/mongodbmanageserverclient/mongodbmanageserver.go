// Code generated by goctl. DO NOT EDIT.
// Source: mongodbmanage.proto

package mongodbmanageserverclient

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mongodbmanage/types/mongodbmanageserver"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest     = mongodbmanageserver.IdRequest
	MinioResponse = mongodbmanageserver.MinioResponse

	MongodbManageServer interface {
		GetMinio(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MinioResponse, error)
	}

	defaultMongodbManageServer struct {
		cli zrpc.Client
	}
)

func NewMongodbManageServer(cli zrpc.Client) MongodbManageServer {
	return &defaultMongodbManageServer{
		cli: cli,
	}
}

func (m *defaultMongodbManageServer) GetMinio(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MinioResponse, error) {
	client := mongodbmanageserver.NewMongodbManageServerClient(m.cli.Conn())
	return client.GetMinio(ctx, in, opts...)
}
