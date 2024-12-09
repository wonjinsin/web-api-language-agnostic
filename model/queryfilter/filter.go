package queryfilter

import (
	"context"
)

// Filter is for query filter
type Filter struct {
	Offset *int `query:"offset"`
	Limit  *int `query:"limit"`
}

// QueryFilter for query filter, support gorm
type QueryFilter interface {
	MakeQuery(ctx context.Context, scope interface{}) (interface{}, error)
}
