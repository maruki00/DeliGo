package pkgTenant

import "context"

type TenantCtx struct {
	Ctx      context.Context
	TenantId string
	UserInfo any
	PG       any
}
