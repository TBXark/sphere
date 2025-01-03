package dash

import (
	"context"
	dashv1 "github.com/TBXark/sphere/api/dash/v1"
)

var _ dashv1.SystemServiceHTTPServer = (*Service)(nil)

func (s *Service) CacheReset(ctx context.Context, req *dashv1.CacheResetRequest) (*dashv1.CacheResetResponse, error) {
	err := s.Cache.DelAll(ctx)
	if err != nil {
		return nil, err
	}
	return &dashv1.CacheResetResponse{}, nil
}

func (s *Service) MenuAll(ctx context.Context, request *dashv1.MenuAllRequest) (*dashv1.MenuAllResponse, error) {
	return &dashv1.MenuAllResponse{}, nil
}
