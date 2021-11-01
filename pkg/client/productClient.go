package client

import (
	"context"
	"github.com/reactivex/rxgo/v2"
)

// ProductClient - ProductClient interface
type ProductClient interface {
	GetProductByID(ctx context.Context, id int) rxgo.Observable
}
