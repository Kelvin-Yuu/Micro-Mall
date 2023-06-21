package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"rpc-common/types"
	"userapi/internal/logic"
	"userapi/internal/svc"
)

type UserHandler struct {
	serverCtx *svc.ServiceContext
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var req types.IdRequest
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.serverCtx)
	resp, err := l.GetUser(&req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	//直接生成一个token字符串
	var req types.LoginRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.serverCtx)
	resp, err := l.Login(&req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

func NewUserHandler(serverCtx *svc.ServiceContext) *UserHandler {
	return &UserHandler{
		serverCtx: serverCtx,
	}
}
