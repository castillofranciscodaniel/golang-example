package client

import (
	error2 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/error"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/http"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"context"
	"fmt"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ProductClientImpl - ProductClientImpl interface
type ProductClientImpl struct {
	log       zerolog.Logger
	url       string
	webClient http.WebClient
}

func NewProductClientImpl(webClient http.WebClient) ProductClient {
	return &ProductClientImpl{
		log:       log.With().Str(utils.Struct, "ProductClientImpl").Logger(),
		url:       "http://localhost:2000/api/rest/v1/products",
		webClient: webClient,
	}
}

func (p *ProductClientImpl) GetProductByID(ctx context.Context, id int) rxgo.Observable {
	ctxString := middleware.GetReqID(ctx)
	subLogger := p.log.With().Str("ctx", ctxString).Str(utils.Method, "GetProductByID").Logger()
	subLogger.Info().Int("id", id).Msg(utils.InitStr)

	url := fmt.Sprintf("%v/id/2", p.url)
	subLogger.Info().Str("url", url).Msg(utils.Data)

	var product dto.Product
	return p.webClient.HTTPDoSimpleReq(ctxString, url, nil, http.GET, &product).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().
				Err(item.E).
				Msg(utils.EndExceptionStr)
			error2.DeleteTrace(item.E)
			return rxgo.Just(item.E)()
		}

		subLogger.Info().Int("id", product.Id).Msgf("httpRequest OK - %v", utils.EndStr)
		return rxgo.Just(item.V)()
	})

}