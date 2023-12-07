package main

import (
	"fmt"
	"net"
	"time"

	"github.com/sllt/arpc/extension/protocol/websocket"
	"github.com/sllt/arpc/extension/pubsub"
	"github.com/sllt/arpc/log"
)

var (
	password = "123qwe"

	topicName = "Broadcast"
)

func onTopic(topic *pubsub.Topic) {
	log.Info("[OnTopic] [%v] \"%v\", [%v]",
		topic.Name,
		string(topic.Data),
		time.Unix(topic.Timestamp/1000000000, topic.Timestamp%1000000000).Format("2006-01-02 15:04:05.000"))
}

func consumer(c *pubsub.Client) {
	if err := c.Subscribe(topicName, onTopic, time.Second); err != nil {
		panic(err)
	}
}

func producer(c *pubsub.Client) {
	ticker := time.NewTicker(time.Second)
	for i := 0; true; i++ {
		_, ok := <-ticker.C
		if !ok {
			break
		}
		if i%5 == 0 {
			c.Publish(topicName, fmt.Sprintf("message from client %d", i), time.Second)
		} else {
			c.PublishToOne(topicName, fmt.Sprintf("message from client %d", i), time.Second)
		}
	}
}

func dialer() (net.Conn, error) {
	return websocket.Dial("ws://localhost:8888/ws")
}

func newClient() *pubsub.Client {
	client, err := pubsub.NewClient(dialer)
	if err != nil {
		panic(err)
	}
	// client.Password = password

	// authentication
	err = client.Authenticate()
	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	{
		for i := 0; i < 5; i++ {
			client := newClient()
			defer client.Stop()
			consumer(client)
		}
	}

	{
		client := newClient()
		defer client.Stop()
		go producer(client)
	}

	<-make(chan int)
}
