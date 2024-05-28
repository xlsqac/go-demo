package handler

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest/httpx"
    "xlsqac/go-zeor-demo/demo/internal/logic"
    "xlsqac/go-zeor-demo/demo/internal/svc"
    "xlsqac/go-zeor-demo/demo/internal/types"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.Request
        if err := httpx.Parse(r, &req); err != nil {
            //httpx.ErrorCtx(r.Context(), w, err)
            httpx.OkJsonCtx(r.Context(), w, err)
            return
        }

        l := logic.NewDemoLogic(r.Context(), svcCtx)
        resp, err := l.Demo(&req)
        if err != nil {
            httpx.OkJsonCtx(r.Context(), w, resp)
        } else {
            httpx.OkJsonCtx(r.Context(), w, resp)
        }
    }
}
