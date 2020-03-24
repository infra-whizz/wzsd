package wzsd

import (
	"log"

	"github.com/nats-io/nats.go"
)

type WzStateDaemonEvents struct {
}

func (wz *WzStateDaemonEvents) onConsoleEvent(m *nats.Msg) {
	log.Println("received from console", len(m.Data), "bytes")
}

func (wz *WzStateDaemonEvents) onResponseEvent(m *nats.Msg) {
	log.Println("received from response channel", len(m.Data), "bytes")
}
