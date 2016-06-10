package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/centrifugal/centrifuge-go"
	"github.com/centrifugal/centrifugo/libcentrifugo"
	"github.com/centrifugal/centrifugo/libcentrifugo/auth"
	"github.com/centrifugal/gocent"
)

type message struct {
	From      string `json:"userId"`
	ToChannel string `json:"toChannel"`
	Text      string `json:"text"`
}

var stopListening chan bool

func listenChatRoom(userId string, channel string, secret string) {
	timestamp := centrifuge.Timestamp()
	info := ""
	token := auth.GenerateClientToken(secret, userId, timestamp, info)
	creds := &centrifuge.Credentials{
		User:      userId,
		Timestamp: timestamp,
		Info:      info,
		Token:     token,
	}
	wsURL := "ws://localhost:8000/connection/websocket"
	c := centrifuge.NewCentrifuge(wsURL, creds, nil, centrifuge.DefaultConfig)
	defer c.Close()

	err := c.Connect()
	if err != nil {
		fmt.Println(err)
	}

	onMessage := func(sub *centrifuge.Sub, msg libcentrifugo.Message) error {
		m := &message{}
		b, err := msg.Data.MarshalJSON()
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, m)
		if err != nil {
			return err
		}

		if m.From != userId {
			fmt.Printf("\tFrom %s: %s\n", m.From, m.Text)
		}
		return nil
	}

	onJoin := func(sub *centrifuge.Sub, msg libcentrifugo.ClientInfo) error {
		fmt.Printf("%v\n", msg)
		return nil
	}

	onLeave := func(sub *centrifuge.Sub, msg libcentrifugo.ClientInfo) error {
		return nil
	}

	events := &centrifuge.SubEventHandler{
		OnMessage: onMessage,
		OnJoin:    onJoin,
		OnLeave:   onLeave,
	}

	sub, err := c.Subscribe(channel, events)
	if err != nil {
		log.Fatal(err)
	}

	<-stopListening

	err = sub.Unsubscribe()
	if err != nil {
		log.Fatalln(err)
	}
}

// this will be called by service backend, not directy by client as in here
func broadcastToChatRoom(userId string, channel string, secret string) {
	c := gocent.NewClient("http://localhost:8000", secret, 5*time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := message{
			From:      userId,
			ToChannel: channel,
			Text:      scanner.Text(),
		}

		j, _ := json.Marshal(msg)

		_, err := c.Publish(channel, j)
		if err != nil {
			fmt.Printf("msg: %s, error: %s\n", msg, err.Error())
			return
		}
	}
}

func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	stopListening <- true
}

func main() {
	userId := "1002"
	roomName := "hede"
	//get this from server
	roomUsers := "1001,1002,1003"
	channel := fmt.Sprintf("%s#%s", roomName, roomUsers)
	secret := "secret"
	go broadcastToChatRoom(userId, channel, secret)
	go listenChatRoom(userId, channel, secret)
	waitForSignal()
}
