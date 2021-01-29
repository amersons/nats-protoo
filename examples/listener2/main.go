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
	logger.Init("debug")
	//npc := nprotoo.NewNatsProtoo(nats)
	npc := nprotoo.NewNatsQueueProtoo(nprotoo.DefaultNatsURL, "4222")
	npc.OnRequest("channel", func(request nprotoo.Request, accept nprotoo.RespondFunc, reject nprotoo.RejectFunc) {
		method := request.Method
		data := request.Data
		logger.Infof("method => %s, data => %v", method, data)

		//accept(JsonEncode(`{}`))
		reject(404, "Not found")
	})

	npc.OnBroadcast("even1", func(data nprotoo.Notification, subj string) {
		logger.Infof("Got Broadcast1 subj => %s, data => %v", subj, data)
	})

	npc.OnBroadcast("even1", func(data nprotoo.Notification, subj string) {
		logger.Infof("Got Broadcast2 subj => %s, data => %v", subj, data)
	})

	select {}
}
