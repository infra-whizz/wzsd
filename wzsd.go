package wzsd

import (
	"log"
	"time"

	"github.com/infra-whizz/wzlib"
	wzlib_transport "github.com/infra-whizz/wzlib/transport"
	"github.com/nats-io/nats.go"
)

type WzChannels struct {
	response *nats.Subscription
	console  *nats.Subscription
}

type WzStateDaemon struct {
	WzStateDaemonEvents
	transport *wzlib_transport.WzdPubSub
	channels  *WzChannels
}

func NewWzStateDaemon() *WzStateDaemon {
	wsd := new(WzStateDaemon)
	wsd.transport = wzlib_transport.NewWizPubSub()
	wsd.channels = new(WzChannels)

	return wsd
}

func (wsd *WzStateDaemon) GetTransport() *wzlib_transport.WzdPubSub {
	return wsd.transport
}

// Run the daemon, prior setting it up.
func (wsd *WzStateDaemon) Run() *WzStateDaemon {
	var err error

	wsd.GetTransport().Start()
	wsd.channels.console, err = wsd.GetTransport().
		GetSubscriber().
		Subscribe(wzlib.CHANNEL_CONSOLE, wsd.onConsoleEvent)
	if err != nil {
		log.Panicf("Unable to subscribe to a console channel: %s\n", err.Error())
	}

	wsd.channels.response, err = wsd.GetTransport().
		GetSubscriber().
		Subscribe(wzlib.CHANNEL_RESPONSE, wsd.onResponseEvent)
	if err != nil {
		log.Panicf("Unable to subscribe to a response channel: %s\n", err.Error())
	}

	return wsd
}

func (wsd *WzStateDaemon) AppLoop() {
	for {
		time.Sleep(10 * time.Second)
	}
}
