package interaction

import (
	"net/http"

	"douyin/internal/logic/interaction"
	"douyin/internal/svc"
	"douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := interaction.NewListCommentLogic(r.Context(), svcCtx)
		resp, err := l.ListComment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
