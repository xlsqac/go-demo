// Code generated by goctl. DO NOT EDIT.
// Source: rpcdemo.proto

package server

import (
	"context"

	"xlsqac/go-zeor-demo/rpcdemo/internal/logic"
	"xlsqac/go-zeor-demo/rpcdemo/internal/svc"
	"xlsqac/go-zeor-demo/rpcdemo/rpcdemo"
)

type RpcdemoServer struct {
	svcCtx *svc.ServiceContext
	rpcdemo.UnimplementedRpcdemoServer
}

func NewRpcdemoServer(svcCtx *svc.ServiceContext) *RpcdemoServer {
	return &RpcdemoServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcdemoServer) Ping(ctx context.Context, in *rpcdemo.Request) (*rpcdemo.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
