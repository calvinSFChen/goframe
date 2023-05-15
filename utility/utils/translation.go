package utils

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

func T(ctx context.Context, content string) string {
	return gi18n.T(ctx, content)
}
