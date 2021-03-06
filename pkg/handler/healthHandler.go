package handler

import (
	service2 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/serializer"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"github.com/castillofranciscodaniel/golang-example/pkg/service"
	"github.com/reactivex/rxgo/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type HealthHandler struct {
	log               zerolog.Logger
	serializerService     service2.SerializerService
	messageChannelService service.MessageChannelService
}

func NewHealthHandler(serializerServer service2.SerializerService, messageChannelService service.MessageChannelService) HealthHandler {
	return HealthHandler{
		log:                   log.With().Str(utils.Struct, "MessageChannelHandler").Logger(),
		messageChannelService: messageChannelService,
		serializerService:     serializerServer,
	}
}

// Health - Handles Healths 200
func (o *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	o.serializerService.ServerResponse(w, rxgo.Empty(), http.StatusOK)
}
