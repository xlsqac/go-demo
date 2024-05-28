package logic

import (
    "context"
    "github.com/zeromicro/go-zero/core/stores/sqlx"

    "xlsqac/go-zeor-demo/demo/internal/svc"
    "xlsqac/go-zeor-demo/demo/internal/types"
    "xlsqac/go-zeor-demo/model/mysql"

    "github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
    return &DemoLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
    // todo: add your logic here and delete this line
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    conn := sqlx.NewMysql(dsn)
    mysql.NewUserModel(conn)

    resp = new(types.Response)
    resp.Message = req.Name
    resp.Code = "0"

    return
}
