package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	apiv1 "github.com/TBXark/sphere/layout/api/api/v1"
	"github.com/TBXark/sphere/layout/internal/pkg/dao"
	"github.com/TBXark/sphere/layout/internal/pkg/database/ent"
	"github.com/TBXark/sphere/layout/internal/pkg/database/ent/user"
	"github.com/TBXark/sphere/server/statuserr"
	"github.com/TBXark/sphere/storage"
	"github.com/TBXark/sphere/utils/safe"
	"github.com/TBXark/sphere/wechat"
)

var _ apiv1.UserServiceHTTPServer = (*Service)(nil)

var wechatAvatarDomains = map[string]struct{}{
	"thirdwx.qlogo.cn": {},
}

const RemoteImageMaxSize = 1024 * 1024 * 2

var (
	ErrImageSizeExceed     = fmt.Errorf("image size exceed")
	ErrImageHostNotAllowed = fmt.Errorf("image host not allowed")
)

func (s *Service) BindPhoneWxMini(ctx context.Context, req *apiv1.BindPhoneWxMiniRequest) (*apiv1.BindPhoneWxMiniResponse, error) {
	userId, err := s.GetCurrentID(ctx)
	if err != nil {
		return nil, err
	}
	number, err := s.wechat.GetUserPhoneNumber(ctx, req.Code, wechat.WithRetryable(true))
	if err != nil {
		return nil, err
	}
	if number.PhoneInfo.CountryCode != "86" {
		return nil, statuserr.NewError(400, "只支持中国大陆手机号")
	}
	err = dao.WithTxEx(ctx, s.db.Client, func(ctx context.Context, client *ent.Client) error {
		exist, e := client.User.Query().Where(user.PhoneEQ(number.PhoneInfo.PhoneNumber)).Only(ctx)
		if e != nil {
			if ent.IsNotFound(e) {
				_, ue := client.User.UpdateOneID(userId).SetPhone(number.PhoneInfo.PhoneNumber).Save(ctx)
				return ue
			}
			return e
		}
		if exist.ID != userId {
			return statuserr.NewError(400, "手机号已被绑定")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &apiv1.BindPhoneWxMiniResponse{}, nil
}

func (s *Service) Me(ctx context.Context, req *apiv1.MeRequest) (*apiv1.MeResponse, error) {
	id, err := s.GetCurrentID(ctx)
	if err != nil {
		return nil, err
	}
	me, err := s.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &apiv1.MeResponse{
		User: s.render.Me(me),
	}, nil
}

func (s *Service) Update(ctx context.Context, req *apiv1.UpdateRequest) (*apiv1.UpdateResponse, error) {
	id, err := s.GetCurrentID(ctx)
	if err != nil {
		return nil, err
	}
	req.Avatar, err = s.uploadRemoteImage(ctx, req.Avatar)
	if err != nil {
		return nil, err
	}
	req.Avatar = s.storage.ExtractKeyFromURL(req.Avatar)
	up, err := s.db.User.UpdateOneID(id).
		SetUsername(req.Username).
		SetAvatar(req.Avatar).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateResponse{
		User: s.render.Me(up),
	}, nil
}

func (s *Service) uploadRemoteImage(ctx context.Context, uri string) (string, error) {
	key, err := s.storage.ExtractKeyFromURLWithMode(uri, true)
	if key != "" && err == nil {
		return key, nil
	}
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	isValidHost := false
	for domain := range wechatAvatarDomains {
		if strings.HasSuffix(u.Host, domain) {
			isValidHost = true
			break
		}
	}
	if !isValidHost {
		return "", ErrImageHostNotAllowed
	}
	id, err := s.GetCurrentID(ctx)
	if err != nil {
		return "", err
	}
	key = storage.DefaultKeyBuilder(strconv.Itoa(int(id)))(uri, "user")
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.ContentLength > RemoteImageMaxSize {
		return "", ErrImageSizeExceed
	}
	defer safe.IfErrorPresent("close response body", resp.Body.Close)
	ret, err := s.storage.UploadFile(ctx, resp.Body, key)
	if err != nil {
		return "", err
	}
	return ret, nil
}
