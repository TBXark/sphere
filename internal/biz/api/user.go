package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tbxark/go-base-api/internal/pkg/dao"
	"github.com/tbxark/go-base-api/pkg/dao/ent"
	"github.com/tbxark/go-base-api/pkg/dao/ent/user"
	"github.com/tbxark/go-base-api/pkg/web"
	"github.com/tbxark/go-base-api/pkg/web/model"
)

type UserInfoMePlatform struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Join bool   `json:"join"`
}

type UserInfoMeResponse struct {
	Info    *ent.User `json:"info"`
	Inviter *ent.User `json:"inviter"`
}

// UserInfoMe
// @Summary 获取当前用户信息
// @Tags api
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} UserInfoMeResponse
// @Router /api/user/me [get]
func (w *Web) UserInfoMe(ctx *gin.Context) (*UserInfoMeResponse, error) {
	id, err := w.auth.GetCurrentID(ctx)
	if err != nil {
		return nil, err
	}
	me, err := w.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	res := UserInfoMeResponse{
		Info: w.render.Me(me),
	}
	return &res, nil
}

type UpdateUserInfoRequest struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// UpdateUserInfo
// @Summary 更新用户信息
// @Tags api
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body UpdateUserInfoRequest true "用户信息"
// @Success 200 {object} ent.User
// @Router /api/user/update [post]
func (w *Web) UpdateUserInfo(ctx *gin.Context) (gin.H, error) {
	id, err := w.auth.GetCurrentID(ctx)
	if err != nil {
		return nil, err
	}
	var info UpdateUserInfoRequest
	if e := ctx.BindJSON(&info); e != nil {
		return nil, e
	}
	info.Avatar, err = w.uploadRemoteImage(ctx, info.Avatar)
	if err != nil {
		return nil, err
	}
	info.Avatar = w.cdn.KeyFromURL(info.Avatar)
	up, err := w.db.User.UpdateOneID(id).
		SetUsername(info.Username).
		SetAvatar(info.Avatar).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"info": w.render.Me(up),
	}, nil
}

type WxMiniBindPhoneRequest struct {
	Code string `json:"code"`
}

// WxMiniBindPhone
// @Summary 绑定手机号
// @Tags api
// @Accept json
// @Produce json
// @Param request body WxMiniBindPhoneRequest true "绑定信息"
// @Success 200 {object} model.MessageResponse
// @Router /api/wx/mini/bind/phone [post]
func (w *Web) WxMiniBindPhone(ctx *gin.Context) (*model.MessageResponse, error) {
	var req WxMiniBindPhoneRequest
	if err := ctx.BindJSON(&req); err != nil {
		return nil, err
	}
	userId, err := w.auth.GetCurrentID(ctx)
	if err != nil {
		return nil, err
	}
	number, err := w.wx.GetUserPhoneNumber(req.Code, true)
	if err != nil {
		return nil, err
	}
	if number.PhoneInfo.CountryCode != "86" {
		return nil, model.NewHTTPError(400, "只支持中国大陆手机号")
	}
	err = dao.WithTxEx(ctx, w.db.Client, func(ctx context.Context, client *ent.Client) error {
		exist, e := client.User.Query().Where(user.PhoneEQ(number.PhoneInfo.PhoneNumber)).Only(ctx)
		if e != nil {
			if ent.IsNotFound(e) {
				_, ue := client.User.UpdateOneID(userId).SetPhone(number.PhoneInfo.PhoneNumber).Save(ctx)
				return ue
			}
			return e
		}
		if exist.ID != userId {
			return model.NewHTTPError(400, "手机号已被绑定")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return model.NewSuccessResponse(), nil
}

func (w *Web) bindUserRoute(r gin.IRouter) {
	r.GET("/api/user/me", web.WithJson(w.UserInfoMe))
	r.POST("/api/user/update", web.WithJson(w.UpdateUserInfo))
	r.POST("/api/wx/mini/bind/phone", web.WithJson(w.WxMiniBindPhone))
}
