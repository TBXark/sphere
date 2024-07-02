package dao

import (
	"context"
	"github.com/tbxark/go-base-api/pkg/dao/ent"
	"github.com/tbxark/go-base-api/pkg/dao/ent/admin"
)

func (d *Dao) GetAdmins(ctx context.Context, ids []int) (map[int]*ent.Admin, error) {
	admins, err := d.Client.Admin.Query().Where(admin.IDIn(RemoveDuplicateAndZero(ids)...)).All(ctx)
	if err != nil {
		return nil, err
	}
	adminMap := make(map[int]*ent.Admin)
	for _, a := range admins {
		adminMap[a.ID] = a
	}
	return adminMap, nil
}
