package service

import (
	"context"
	"github.com/reactivex/rxgo/v2"
)

// ProductService -
type ProductService interface {
	HandlerProductByID(ctx context.Context, id int) rxgo.Observable
}
