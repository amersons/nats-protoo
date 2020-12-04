package nprotoo

import (
	"encoding/json"

	"github.com/amersons/nats-protoo/logger"
	"github.com/chuckpreslar/emission"
	nats "github.com/nats-io/nats.go"
)

// Broadcaster .
type Broadcaster struct {
	emission.Emitter
	subj string
	np   *NatsProtoo
}

func newBroadcaster(subj string, np *NatsProtoo, nc *nats.Conn) *Broadcaster {
	var bc Broadcaster
	bc.Emitter = *emission.NewEmitter()
	bc.subj = subj
	bc.np = np
	bc.np.On("close", func(code int, err string) {
		logger.Infof("Transport closed [%d] %s", code, err)
		bc.Emit("close", code, err)
	})
	bc.np.On("error", func(code int, err string) {
		logger.Warnf("Transport got error (%d, %s)", code, err)
		bc.Emit("error", code, err)
	})
	return &bc
}

// Say .
func (bc *Broadcaster) Say(method string, data interface{}) {
	dataStr, err := json.Marshal(data)
	if err != nil {
		logger.Errorf("Marshal data %v", err)
		return
	}
	notification := &Notification{
		NotificationData: NotificationData{
			Notification: true,
		},
		CommonData: CommonData{
			Method: method,
			Data:   dataStr,
		},
	}
	str, err := json.Marshal(notification)
	if err != nil {
		logger.Errorf("Marshal %v", err)
		return
	}
	logger.Debugf("Send notification [%s] queue[%s]", method, bc.np.queue)
	if bc.np.queue != "" {
		bc.np.Send(str, bc.subj, bc.np.queue)
		return
	}
	bc.np.Send(str, bc.subj, _EMPTY_)
}
