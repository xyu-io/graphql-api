package auth_tool

import (
	"context"
	"testing"
)

func TestUserFromContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ActiveUserContextKey, &Identity{Id: "12", OrgId: "13"})
	uc, err := UserFromContext(ctx)
	if err != nil {
		t.Error(err)
	}
	if uc.IdInt() != 12 || uc.OrgIdInt() != 13 {
		t.Error()
	}
}
