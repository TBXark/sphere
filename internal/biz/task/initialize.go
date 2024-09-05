package task

import (
	"context"
	"github.com/tbxark/go-base-api/internal/pkg/dao"
	"github.com/tbxark/go-base-api/internal/pkg/encrypt"
	"github.com/tbxark/go-base-api/pkg/dao/ent"
	"github.com/tbxark/go-base-api/pkg/dao/ent/keyvaluestore"
	"strconv"
	"time"
)

type Initialize struct {
	db *dao.Dao
}

func NewInitialize(db *dao.Dao) *Initialize {
	return &Initialize{db: db}
}

func initAdminIfNeed(ctx context.Context, client *ent.Client) error {
	count, err := client.Admin.Query().Count(context.Background())
	if err != nil || count > 0 {
		return nil
	}
	return client.Admin.Create().
		SetUsername("admin").
		SetPassword(encrypt.CryptPassword("aA1234567")).
		SetRoles([]string{"all"}).
		OnConflict().
		Ignore().
		Exec(ctx)
}

func (i *Initialize) Identifier() string {
	return "initialize"
}

func (i *Initialize) Run() error {
	ctx := context.Background()
	key := "did_init"
	return dao.WithTxEx(ctx, i.db.Client, func(ctx context.Context, client *ent.Client) error {
		exist, err := client.KeyValueStore.Query().Where(keyvaluestore.KeyEQ(key)).Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return nil
		}
		_, err = client.KeyValueStore.Create().
			SetKey(key).
			SetValue([]byte(strconv.Itoa(int(time.Now().Unix())))).
			Save(ctx)
		if err != nil {
			return err
		}
		_ = initAdminIfNeed(ctx, client)
		return nil

	})
}
