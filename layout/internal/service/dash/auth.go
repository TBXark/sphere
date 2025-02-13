package dash

import (
	"context"
	dashv1 "github.com/TBXark/sphere/layout/api/dash/v1"
	"github.com/TBXark/sphere/layout/internal/pkg/database/ent"
	"github.com/TBXark/sphere/layout/internal/pkg/database/ent/admin"
	"github.com/TBXark/sphere/server/auth/authorizer"
	"github.com/TBXark/sphere/server/ginx"
	"github.com/TBXark/sphere/server/statuserr"
	"github.com/TBXark/sphere/utils/secure"
	"time"
)

var _ dashv1.AuthServiceHTTPServer = (*Service)(nil)

const (
	AuthTokenValidDuration    = time.Hour * 24
	RefreshTokenValidDuration = time.Hour * 24 * 30
	AuthExpiresTimeFormat     = "2006/01/02 15:04:05"
)

var (
	ErrPasswordNotMatch = statuserr.NewError(400, "password not match")
)

type AdminToken struct {
	Admin        *ent.Admin
	AccessToken  string
	RefreshToken string
	Expires      string
}

type AdminLoginResponseWrapper = ginx.DataResponse[AdminToken]

func renderClaims(admin *ent.Admin, duration time.Duration) *authorizer.RBACClaims[int64] {
	return authorizer.NewRBACClaims(admin.ID, admin.Username, admin.Roles, time.Now().Add(duration))
}

func (s *Service) createToken(u *ent.Admin) (*AdminToken, error) {
	claims := renderClaims(u, AuthTokenValidDuration)
	token, err := s.Authorizer.GenerateToken(claims)
	if err != nil {
		return nil, err
	}
	refresh, err := s.AuthRefresher.GenerateToken(renderClaims(u, RefreshTokenValidDuration))
	if err != nil {
		return nil, err
	}
	u.Avatar = s.Storage.GenerateImageURL(u.Avatar, 512)
	return &AdminToken{
		Admin:        u,
		AccessToken:  token,
		RefreshToken: refresh,
		Expires:      claims.ExpiresAt.Format(AuthExpiresTimeFormat),
	}, nil
}

func (s *Service) AuthCodes(ctx context.Context, request *dashv1.AuthCodesRequest) (*dashv1.AuthCodesResponse, error) {
	return &dashv1.AuthCodesResponse{
		Data: make([]string, 0),
	}, nil
}

func (s *Service) AuthLogin(ctx context.Context, req *dashv1.AuthLoginRequest) (*dashv1.AuthLoginResponse, error) {
	u, err := s.DB.Admin.Query().Where(admin.UsernameEQ(req.Username)).Only(ctx)
	if err != nil {
		return nil, ErrPasswordNotMatch // 隐藏错误信息
	}
	if !secure.IsPasswordMatch(req.Password, u.Password) {
		return nil, ErrPasswordNotMatch
	}
	token, err := s.createToken(u)
	if err != nil {
		return nil, err
	}
	return &dashv1.AuthLoginResponse{
		AccessToken: token.AccessToken,
	}, nil
}

func (s *Service) AuthLogout(ctx context.Context, request *dashv1.AuthLogoutRequest) (*dashv1.AuthLogoutResponse, error) {
	return &dashv1.AuthLogoutResponse{}, nil
}

func (s *Service) AuthRefresh(ctx context.Context, request *dashv1.AuthRefreshRequest) (*dashv1.AuthRefreshResponse, error) {
	id, err := s.GetCurrentID(ctx)
	if err != nil {
		return nil, err
	}
	u, err := s.DB.Admin.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	token, err := s.createToken(u)
	if err != nil {
		return nil, err
	}
	return &dashv1.AuthRefreshResponse{
		AccessToken: token.AccessToken,
	}, nil
}
