package client

import (
	service "bitbucket.org/chattigodev/chattigo-golang-library/pkg/deserializer"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"context"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// MessageChannelClientImpl - MessageChannelClientImpl interface
type MessageChannelClientImpl struct {
	log                 zerolog.Logger
	url                 string
	restyClient         *resty.Client
	deserializerService service.DeserializerService
}

func NewMessageChannelClientImpl(restyClient *resty.Client, deserializerService service.DeserializerService) MessageChannelClient {
	return &MessageChannelClientImpl{
		log:                 log.With().Str(utils.Struct, messageChannelClientImpl).Logger(),
		url:                 "http://localhost:2000/api/rest/v1/message",
		deserializerService: deserializerService,
		restyClient:         restyClient,
	}
}

const (
	getMessageChannel        = "GetMessageChannel"
	saveMessageChannel       = "SaveMessageChannel"
	messageChannelClientImpl = "MessageChannelClientImpl"
)

func (m *MessageChannelClientImpl) SaveMessageChannel(ctx context.Context, messageChannel *dto.MessageChannel) rxgo.Observable {
	ctxString := middleware.GetReqID(ctx)
	subLogger := m.log.With().Str("ctx", ctxString).Str(utils.Method, saveMessageChannel).Logger()
	subLogger.Info().Msg(utils.InitStr)

	subLogger.Info().Str("url", m.url).Msg(utils.Data)

	var message dto.MessageChannel

	response, err := m.restyClient.R().
		SetBody(messageChannel).
		Post(m.url)

	if err != nil {
		subLogger.Error().Err(err).Msg(utils.EndExceptionStr)
		return rxgo.Just(err)()
	}

	return m.deserializerService.BodyFromClient(ctxString, response.RawResponse, &message).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msg(utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		return rxgo.Just(item.V)()
	})

}

func (m *MessageChannelClientImpl) GetMessageChannel(ctx context.Context) rxgo.Observable {
	ctxString := middleware.GetReqID(ctx)
	subLogger := m.log.With().Str("ctx", ctxString).Str(utils.Method, getMessageChannel).Logger()
	subLogger.Info().Msg(utils.InitStr)

	var messageChannels []dto.MessageChannel

	response, err := m.restyClient.R().
		Get(m.url)

	if err != nil {
		subLogger.Error().Err(err).Msg(utils.EndExceptionStr)
		return rxgo.Just(err)()
	}

	return m.deserializerService.BodyFromClient(ctxString, response.RawResponse, &messageChannels).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msg(utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		return rxgo.Just(item.V)()
	})

}
