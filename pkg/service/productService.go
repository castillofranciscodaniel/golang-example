package service

import (
	"context"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/reactivex/rxgo/v2"
)

// ProductService -
type ProductService interface {
	HandlerProductByID(ctx context.Context, product dto.Product) rxgo.Observable

	HandlerProductByIDPointer(ctx context.Context, product *dto.Product) rxgo.Observable

	GetProducts(ctx context.Context) rxgo.Observable

}
