package handler

import (
	service3 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/deserializer"
	service2 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/serializer"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/castillofranciscodaniel/golang-example/pkg/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type MessageChannelHandler struct {
	log                   zerolog.Logger
	serializerService     service2.SerializerService
	deserializerService   service3.DeserializerService
	messageChannelService service.MessageChannelService
}

func NewMessageChannelHandler(serializerService service2.SerializerService, deserializerService service3.DeserializerService,
	messageChannelService service.MessageChannelService,
) MessageChannelHandler {
	return MessageChannelHandler{
		log:                   log.With().Str(utils.Struct, "MessageChannelHandler").Logger(),
		serializerService:     serializerService,
		deserializerService:   deserializerService,
		messageChannelService: messageChannelService,
	}
}

const (
	saveMessageChannel = "SaveMessageChannel"
	getMessageChannel  = "GetMessageChannel"
)

// SaveMessageChannel -
func (m *MessageChannelHandler) SaveMessageChannel(w http.ResponseWriter, r *http.Request) {
	ctxString := middleware.GetReqID(r.Context())
	subLogger := m.log.With().Str(utils.Thread, ctxString).Str(utils.Method, saveMessageChannel).Logger()
	subLogger.Info().Msgf(utils.InitStr)

	var messageChannel dto.MessageChannel

	result := m.deserializerService.BodyToObservable(ctxString, r, &messageChannel).FlatMap(func(item rxgo.Item) rxgo.Observable {

		return m.messageChannelService.SaveMessageChannel(r.Context(), &messageChannel).FlatMap(func(item rxgo.Item) rxgo.Observable {
			if item.Error() {
				subLogger.Error().Err(item.E).Msgf("[An error from webClient][%v]", utils.EndExceptionStr)
				return rxgo.Just(item.E)()
			}

			subLogger.Info().Msgf("httpRequest OK - %v", utils.EndStr)
			return rxgo.Just(item.V)()
		})
	})

	m.serializerService.ServerResponse(w, result.First(), http.StatusCreated)
}

// GetMessageChannel -
func (m *MessageChannelHandler) GetMessageChannel(w http.ResponseWriter, r *http.Request) {
	ctxString := middleware.GetReqID(r.Context())
	subLogger := m.log.With().Str(utils.Thread, ctxString).Str(utils.Method, getMessageChannel).Logger()
	subLogger.Info().Msgf(utils.InitStr)

	result := m.messageChannelService.GetMessageChannel(r.Context()).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msgf("[An error from webClient][%v]", utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		subLogger.Info().Msgf("httpRequest OK - %v", utils.EndStr)
		return rxgo.Just(item.V)()
	})

	m.serializerService.ServerResponse(w, result.First(), http.StatusOK)
}
