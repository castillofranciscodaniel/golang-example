package client

import (
	"context"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/reactivex/rxgo/v2"
)

// ProductClient - ProductClient interface
type ProductClient interface {
	GetProductByID(ctx context.Context, product dto.Product) rxgo.Observable
}
