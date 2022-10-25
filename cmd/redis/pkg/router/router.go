package router

import (
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/pkg/interfaces"
)

type Router struct {
	store interfaces.EventStore[[]byte]
}

func (r *Router) Publish(e interfaces.Event) error {
	switch e.(type) {
	case event.PixelDrawnEvent:

	}
}

func New(store interfaces.EventStore[[]byte]) *Router {
	return &Router{store: store}
}
