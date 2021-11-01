package handler

import (
	service3 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/deserializer"
	service2 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/serializer"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"errors"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/castillofranciscodaniel/golang-example/pkg/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ProductHandler struct {
	log                 zerolog.Logger
	serializerService   service2.SerializerService
	deserializerService service3.DeserializerService
	productService      service.ProductService
}

func NewProductHandler(serializerService service2.SerializerService, deserializerService service3.DeserializerService,
	productService service.ProductService,
) ProductHandler {
	return ProductHandler{
		log:                 log.With().Str(utils.Struct, "ProductHandler").Logger(),
		serializerService:   serializerService,
		deserializerService: deserializerService,
		productService:      productService,
	}
}

// HandlerProductByID -
func (p *ProductHandler) HandlerProductByID(w http.ResponseWriter, r *http.Request) {
	ctxString := middleware.GetReqID(r.Context())
	subLogger := p.log.With().Str(utils.Thread, ctxString).Str(utils.Method, "Inbound").Logger()
	subLogger.Info().Msgf(utils.InitStr)

	var product dto.Product

	result := p.deserializerService.BodyToObservable(ctxString, r, &product).FlatMap(func(item rxgo.Item) rxgo.Observable {

		return p.productService.HandlerProductByID(r.Context(), product).FlatMap(func(item rxgo.Item) rxgo.Observable {
			if item.Error() {
				subLogger.Error().Err(item.E).Msgf("[An error from webClient][%v]", utils.EndExceptionStr)
				return rxgo.Just(item.E)()
			}

			subLogger.Info().Int("id", product.Id).Msgf("httpRequest OK - %v", utils.EndStr)
			return rxgo.Just(item.V)()
		})
	})

	p.serializerService.ServerResponse(w, result.First(), http.StatusCreated)
}

func (p *ProductHandler) TestErrorDto(w http.ResponseWriter, r *http.Request) {
	ctxString := middleware.GetReqID(r.Context())
	subLogger := p.log.With().Str(utils.Thread, ctxString).Str(utils.Method, "Inbound").Logger()
	subLogger.Info().Msgf(utils.InitStr)

	//error := error2.ErrorDto{
	//	StatusCode: 514,
	//	Msg:        "Status 514",
	//	TraceDetail: map[string]interface{}{
	//		"error": "error 1",
	//	},
	//}

	//product := dto.Product{
	//			Id:    200,
	//			Name:  "holi",
	//			Price: 500,
	//		}

	//p.serializerService.ServerResponse(w, result.First(), http.StatusCreated)
	p.serializerService.ServerEmptyResponse(w, rxgo.Just(errors.New("sdsdf"))(), http.StatusCreated)
}
