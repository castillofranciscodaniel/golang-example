package client

import (
	service "bitbucket.org/chattigodev/chattigo-golang-library/pkg/deserializer"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"bitbucket.org/chattigodev/chattigo-golang-logging-library/pkg/log"
	"context"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"
	"github.com/reactivex/rxgo/v2"
)

// MessageChannelClientImpl - MessageChannelClientImpl interface
type MessageChannelClientImpl struct {
	url                 string
	restyClient         *resty.Client
	deserializerService service.DeserializerService
}

func NewMessageChannelClientImpl(restyClient *resty.Client, deserializerService service.DeserializerService) MessageChannelClient {
	return &MessageChannelClientImpl{
		url:                 "http://localhost:2000/api/rest/v1/message",
		deserializerService: deserializerService,
		restyClient:         restyClient,
	}
}

const (
	structName = "MessageChannelClientImpl"
	getMessageChannel        = "GetMessageChannel"
	saveMessageChannel       = "SaveMessageChannel"
)

func (m *MessageChannelClientImpl) SaveMessageChannel(ctx context.Context, messageChannel *dto.MessageChannel) rxgo.Observable {
	ctxString := middleware.GetReqID(ctx)
	log.GetInstance().Info().
		Base(middleware.GetReqID(ctx), structName, saveMessageChannel).
		Did(messageChannel.Did).
		CampaignId(messageChannel.CampaignId).
		UserId(messageChannel.IdUser).
		Msisdn(messageChannel.Msisdn).
		Msgf("url: %v. %v", m.url, utils.InitStr)

	response, err := m.restyClient.R().
		SetBody(messageChannel).
		Post(m.url)

	if err != nil {
		log.GetInstance().Error().
			Err(err).
			Base(middleware.GetReqID(ctx), structName, saveMessageChannel).
			Did(messageChannel.Did).
			CampaignId(messageChannel.CampaignId).
			UserId(messageChannel.IdUser).
			Msisdn(messageChannel.Msisdn).
			Msgf("url: %v. %v", m.url, utils.EndExceptionStr)
		return rxgo.Just(err)()
	}

	var message dto.MessageChannel
	return m.deserializerService.BodyFromClient(ctxString, response.RawResponse, &message).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			log.GetInstance().Error().
				Err(item.E).
				Base(middleware.GetReqID(ctx), structName, saveMessageChannel).
				Did(messageChannel.Did).
				CampaignId(messageChannel.CampaignId).
				UserId(messageChannel.IdUser).
				Msisdn(messageChannel.Msisdn).
				Msgf("url: %v. %v", m.url, utils.EndExceptionStr)

			return rxgo.Just(item.E)()
		}

		return rxgo.Just(item.V)()
	})

}

func (m *MessageChannelClientImpl) GetMessageChannel(ctx context.Context) rxgo.Observable {
	ctxString := middleware.GetReqID(ctx)
	log.GetInstance().Info().
		Base(middleware.GetReqID(ctx), structName, getMessageChannel).
		Msgf("url: %v. %v", m.url, utils.InitStr)

	var messageChannels []dto.MessageChannel

	response, err := m.restyClient.R().
		Get(m.url)

	if err != nil {
		log.GetInstance().Error().
			Err(err).
			Base(middleware.GetReqID(ctx), structName, getMessageChannel).
			Msgf("url: %v. %v", m.url, utils.EndExceptionStr)
		return rxgo.Just(err)()
	}

	return m.deserializerService.BodyFromClient(ctxString, response.RawResponse, &messageChannels).FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			log.GetInstance().Error().
				Err(item.E).
				Base(middleware.GetReqID(ctx), structName, getMessageChannel).
				Msgf("url: %v. %v", m.url, utils.EndExceptionStr)
			return rxgo.Just(item.E)()
		}

		return rxgo.Just(item.V)()
	})

}
