package client

import (
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/http"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"context"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/go-chi/chi/v5/middleware"
	jsoniter "github.com/json-iterator/go"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// MessageChannelClientImpl - MessageChannelClientImpl interface
type MessageChannelClientImpl struct {
	log       zerolog.Logger
	url       string
	webClient http.WebClient
}

func NewMessageChannelClientImpl(webClient http.WebClient) MessageChannelClient {
	return &MessageChannelClientImpl{
		log:       log.With().Str(utils.Struct, messageChannelClientImpl).Logger(),
		url:       "http://localhost:2000/api/rest/v1/message",
		webClient: webClient,
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

	var response dto.MessageChannel
	return rxgo.Just(messageChannel)().Marshal(jsoniter.Marshal).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msg(utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		return m.webClient.HTTPDoSimpleReq(ctxString, m.url, item.V.([]byte), http.POST, &response).FlatMap(func(item rxgo.Item) rxgo.Observable {
			if item.Error() {
				subLogger.Error().Err(item.E).Msg(utils.EndExceptionStr)
				return rxgo.Just(item.E)()
			}

			return rxgo.Just(item.V)()
		})
	})

}

func (m *MessageChannelClientImpl) GetMessageChannel(ctx context.Context) rxgo.Observable {
	ctxString := middleware.GetReqID(ctx)
	subLogger := m.log.With().Str("ctx", ctxString).Str(utils.Method, getMessageChannel).Logger()
	subLogger.Info().Msg(utils.InitStr)

	var messageChannels []dto.MessageChannel

	return m.webClient.HTTPDoSimpleReq(ctxString, m.url, nil, http.GET, &messageChannels).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msg(utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		return rxgo.Just(item.V)()
	})

}
