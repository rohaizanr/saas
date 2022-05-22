package ent

import (
	"context"
	"database/sql"
	"entgo.io/ent"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/data"
)

const DbKind = "ent"

func Saas(next ent.Mutator) ent.Mutator {
	type hasTenant interface {
		SetTenantID(ss *sql.NullString)
	}
	return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
		if hf, ok := mutation.(hasTenant); ok {
			ct, _ := common.FromCurrentTenant(ctx)
			at := data.FromAutoSetTenantId(ctx)
			if ok && at {
				if ct.GetId() != "" {
					//normalize tenant side only
					hf.SetTenantID(&sql.NullString{
						String: ct.GetId(),
						Valid:  true,
					})
				}
			}
		}
		return next.Mutate(ctx, mutation)
	})
}