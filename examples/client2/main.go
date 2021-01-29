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
	clann := make(chan int, 100)
	go func() {
		count := 0
		for {
			clann <- count
			count++
			go sendTest(clann)
		}
	}()
	select {}
}

func sendTest(clann chan int) {
	count := 0
	npc := nprotoo.NewNatsProtoo(nprotoo.DefaultNatsURL)
Loop:
	for {
		req := npc.NewRequestor("channel")
		dat, err := req.SyncRequest("offer", JsonEncode(`{ "sdp": "channel.aaa.test"}`))
		logger.Infof("data[%s] err[%v] count[%d]", string(dat), err, count)

		req2 := npc.NewBroadcaster("OnBroadcast")
		req2.Say("offer", JsonEncode(`{ "sdp": "channel"}`))
		req.Close()
		if count > 1000 {
			break Loop
		}
		count++
	}
	npc.Close()
	<-clann
}
