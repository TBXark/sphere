package dash

import (
	"context"

	"github.com/TBXark/sphere/database/mapper"
	{{.ServicePackage}} "{{.BizPackagePath}}/api/{{.PackagePath}}"
    "{{.BizPackagePath}}/internal/pkg/render"
)

var _ {{.ServicePackage}}.{{.ServiceName}}ServiceHTTPServer = (*Service)(nil)

func (s *Service) Create{{.ServiceName}}(ctx context.Context, request *{{.ServicePackage}}.Create{{.ServiceName}}Request) (*{{.ServicePackage}}.Create{{.ServiceName}}Response, error) {
	item, err := render.Create{{.ServiceName}}(s.db.{{.ServiceName}}.Create(), request.{{.ServiceName}}).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &{{.ServicePackage}}.Create{{.ServiceName}}Response{
		{{.ServiceName}}: s.render.{{.ServiceName}}(item),
	}, nil
}

func (s *Service) Delete{{.ServiceName}}(ctx context.Context, request *{{.ServicePackage}}.Delete{{.ServiceName}}Request) (*{{.ServicePackage}}.Delete{{.ServiceName}}Response, error) {
	err := s.db.{{.ServiceName}}.DeleteOneID(request.Id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &{{.ServicePackage}}.Delete{{.ServiceName}}Response{}, nil
}

func (s *Service) Get{{.ServiceName}}(ctx context.Context, request *{{.ServicePackage}}.Get{{.ServiceName}}Request) (*{{.ServicePackage}}.Get{{.ServiceName}}Response, error) {
	item, err := s.db.{{.ServiceName}}.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &{{.ServicePackage}}.Get{{.ServiceName}}Response{
		{{.ServiceName}}: s.render.{{.ServiceName}}(item),
	}, nil
}

func (s *Service) List{{plural .ServiceName}}(ctx context.Context, request *{{.ServicePackage}}.List{{plural .ServiceName}}Request) (*{{.ServicePackage}}.List{{plural .ServiceName}}Response, error) {
	query := s.db.{{.ServiceName}}.Query()
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	page, size := mapper.Page(count, int(request.PageSize), mapper.DefaultPageSize)
	all, err := query.Clone().Limit(size).Offset(size * int(request.Page)).All(ctx)
	if err != nil {
		return nil, err
	}
	return &{{.ServicePackage}}.List{{plural .ServiceName}}Response{
		{{plural .ServiceName}}: mapper.Map(all, s.render.{{.ServiceName}}),
		TotalSize: int64(count),
		TotalPage: int64(page),
	}, nil
}

func (s *Service) Update{{.ServiceName}}(ctx context.Context, request *{{.ServicePackage}}.Update{{.ServiceName}}Request) (*{{.ServicePackage}}.Update{{.ServiceName}}Response, error) {
	item, err := render.UpdateOne{{.ServiceName}}(
		s.db.{{.ServiceName}}.UpdateOneID(request.{{.ServiceName}}.Id),
		request.{{.ServiceName}},
	).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &{{.ServicePackage}}.Update{{.ServiceName}}Response{
		{{.ServiceName}}: s.render.{{.ServiceName}}(item),
	}, nil
}
