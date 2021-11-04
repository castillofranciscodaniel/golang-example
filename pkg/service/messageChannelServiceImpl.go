package service

import (
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"context"
	"errors"
	"github.com/castillofranciscodaniel/golang-example/pkg/client"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// MessageChannelServiceIml -
type MessageChannelServiceIml struct {
	log                  zerolog.Logger
	messageChannelClient client.MessageChannelClient
}

func NewMessageChannelServiceIml(messageChannelClient client.MessageChannelClient) MessageChannelService {
	return &MessageChannelServiceIml{
		log:                  log.With().Str(utils.Struct, structName).Logger(),
		messageChannelClient: messageChannelClient,
	}
}

const (
	structName = "MessageChannelServiceIml"
	errCastMessageChannel = "item.V is not a MessageChannel"
	errCastArrayMessageChannel = "item.V is not a []MessageChannel"
)

// SaveMessageChannel -
func (m *MessageChannelServiceIml) SaveMessageChannel(ctx context.Context, messageChannel *dto.MessageChannel) rxgo.Observable {
	subLogger := m.log.With().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, "GetStateAgentCarer").Logger()
	subLogger.Info().Msgf(utils.InitStr)


	return m.messageChannelClient.SaveMessageChannel(ctx, messageChannel).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msgf("[An error from webClient][%v]", utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		message, ok := item.V.(*dto.MessageChannel)
		if !ok {
			err := errors.New(errCastMessageChannel)
			subLogger.Error().Err(item.E).Msg(utils.EndExceptionStr)
			return rxgo.Just(err)()
		}

		message.Did = "holi"

		return rxgo.Just(item.V)()
	})
}

// GetMessageChannel -
func (m *MessageChannelServiceIml) GetMessageChannel(ctx context.Context) rxgo.Observable {
	subLogger := m.log.With().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, "GetStateAgentCarer").Logger()
	subLogger.Info().Msgf(utils.InitStr)

	return m.messageChannelClient.GetMessageChannel(ctx).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			subLogger.Error().Err(item.E).Msgf("[An error from webClient][%v]", utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		messageChannel, ok := item.V.(*[]dto.MessageChannel)
		if !ok {
			err := errors.New(errCastArrayMessageChannel)
			subLogger.Error().Err(item.E).Msg(utils.EndExceptionStr)
			return rxgo.Just(err)()
		}
		for idx, _ := range *messageChannel  {
			(*(messageChannel))[idx].Did = "edit did"
		}
		return rxgo.Just(messageChannel)()
	})
}

