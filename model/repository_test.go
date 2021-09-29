package model

import (
	"context"
	"testing"

	"github.com/traPtitech/Jomon/ent"
)

func setup(t *testing.T, ctx context.Context) (*ent.Client, error) {
	t.Helper()
	client, err := SetupTestEntClient(t)
	if err != nil {
		return nil, err
	}
	err = dropAll(t, ctx, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func dropAll(t *testing.T, ctx context.Context, client *ent.Client) error {
	t.Helper()
	_, err := client.Comment.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.File.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.GroupBudget.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.Request.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.RequestStatus.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.RequestTarget.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.Tag.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.Transaction.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.TransactionDetail.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.User.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
