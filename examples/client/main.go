package main

import (
	"encoding/json"
	nprotoo "github.com/amersons/nats-protoo"
	"github.com/amersons/nats-protoo/logger"
	"time"
)

func JsonEncode(str string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}
	return data
}

func main() {
	logger.Init("debug")
	npc := nprotoo.NewNatsQueueProtoo(nprotoo.DefaultNatsURL, "debug2")
	//npc := nprotoo.NewNatsProtoo(/*nprotoo.DefaultNatsURL*/nats)
	req := npc.NewRequestor("channel.aaa.test1")
	req.AsyncRequest("offer", JsonEncode(`{ "sdp": "dummy-sdp1"}`)).Then(
		func(result nprotoo.RawMessage) {
			logger.Infof("AsyncRequest.Then: offer success: =>  %s", result)
		},
		func(err *nprotoo.Error) {
			logger.Warnf("AsyncRequest.Then: offer reject: %d => %s", err.Code, err.Reason)
		})

	result, err := req.SyncRequest("offer", JsonEncode(`{ "sdp": "dummy-sdp3"}`))
	if err != nil {
		logger.Warnf("offer reject: %d => %s", err.Code, err.Reason)
	} else {
		logger.Infof("offer success: =>  %s", result)
	}

	req.AsyncRequest("offer", JsonEncode(`{ "sdp": "dummy-sdp2"}`))

	bc := npc.NewBroadcaster("even1")
	bc.Say("hello", JsonEncode(`{"key": "value"}`))

	time.Sleep(5 * time.Second)
	req.Close()
}
