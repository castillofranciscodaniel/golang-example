package client

import (
	"context"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/reactivex/rxgo/v2"
)

// MessageChannel - MessageChannel interface
type MessageChannelClient interface {
	SaveMessageChannel(ctx context.Context, messageChannel *dto.MessageChannel) rxgo.Observable

	GetMessageChannel(ctx context.Context) rxgo.Observable
}
