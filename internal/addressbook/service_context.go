package addressbook

import (
	"context"

	"github.com/rs/zerolog/log"
)

type contextKey string

var (
	ctxKeyTenantID      = contextKey("tenantID")
	ctxKeyCorrelationID = contextKey("correlationID")
	ctxKeyBearerToken   = contextKey("bearerToken")
)

func CtxSetTenantID(ctx context.Context, value int64) context.Context {
	return context.WithValue(ctx, ctxKeyTenantID, value)
}

func CtxGetTenantID(ctx context.Context) int64 {
	tenantID, ok := ctx.Value(ctxKeyTenantID).(int64)
	if !ok {
		log.Warn().Msgf("Failed to retrieve TenantID from context.  Expected type int64, got %T", ctx.Value(ctxKeyTenantID))
		return 0
	}
	return tenantID
}

func CtxSetCorrelationID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ctxKeyCorrelationID, value)
}

func CtxGetCorrelationID(ctx context.Context) string {
	correlationID, ok := ctx.Value(ctxKeyCorrelationID).(string)
	if !ok {
		log.Warn().Msgf("Failed to retrieve CorrelationID from context.  Expected type string, got %T", ctx.Value(ctxKeyCorrelationID))
		return ""
	}
	return correlationID
}
