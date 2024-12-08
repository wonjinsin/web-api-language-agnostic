package filter

import (
	"context"

	"gorm.io/gorm"
)

// Filter is for query filter
type Filter struct {
	Offset *int
	Limit  *int
}

// IGormFilter for query filter
type IGormFilter interface {
	MakeQuery(ctx context.Context, scope *gorm.DB) *gorm.DB
}
