package main

import (
	"encoding/json"

	nprotoo "github.com/amersons/nats-protoo"
	"github.com/amersons/nats-protoo/logger"
)

func JsonEncode(str string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}
	return data
}

func main() {
	logger.Init("info")
	npc := nprotoo.NewNatsProtoo(nprotoo.DefaultNatsURL)
	npc.OnRequest("channel", func(request nprotoo.Request, accept nprotoo.RespondFunc, reject nprotoo.RejectFunc) {
		method := request.Method
		data := request.Data
		logger.Infof("channel method => %s, data => %v", method, string(data))
		accept(JsonEncode(`{ "sdp": "channel.*"}`))
	})
	npc.OnBroadcast("OnBroadcast", func(dat nprotoo.Notification, subj string) {
		logger.Infof("OnBroadcast  method => %s, data => %v", dat.Method, string(dat.Data))
	})
	select {}
}
