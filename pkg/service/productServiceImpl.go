package service

import (
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"context"
	"github.com/castillofranciscodaniel/golang-example/pkg/client"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ProductServiceIml -
type ProductServiceIml struct {
	log           zerolog.Logger
	productClient client.ProductClient
}

func NewProductServiceIml(productClient client.ProductClient) ProductService {
	return &ProductServiceIml{
		log:           log.With().Str(utils.Struct, "ProductServiceIml").Logger(),
		productClient: productClient,
	}
}

// HandlerProductByID -
func (m *ProductServiceIml) HandlerProductByID(ctx context.Context, product dto.Product) rxgo.Observable {
	subLogger := m.log.With().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, "GetStateAgentCarer").Logger()
	subLogger.Info().Msgf(utils.InitStr)

	return m.productClient.GetProductByID(ctx, product).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msgf("[An error from webClient][%v]", utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		return rxgo.Just(item.V)()
	})
}
