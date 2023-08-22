package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRelationClient)

// type RelationConn stdgrpc.ClientConnInterface

// // NewRelationClient 创建一个Relation服务客户端，接收Relation服务数据
// func NewRelationClient(c *conf.Client, logger log.Logger) RelationConn {
// 	conn, err := grpc.DialInsecure(
// 		context.Background(),
// 		grpc.WithEndpoint(c.Relation.To),
// 		grpc.WithMiddleware(
// 			recovery.Recovery(),
// 			logging.Client(logger),
// 		),
// 	)
// 	if err != nil {
// 		log.Fatalf("Error connecting to Relation Services, err : %w", err)
// 	}
// 	log.Info("Relation Services connected.")
// 	return conn
// }
