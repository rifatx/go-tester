package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

func broadcastToChatRoom(userId string, channel string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//url to chat service
		_, err := http.Get(fmt.Sprintf("http://localhost:1234/broadcastToChatRoom?userId=%s&channel=%s&msg=%s", userId, url.QueryEscape(channel), scanner.Text()))
		if err != nil {
			fmt.Printf("msg: %s, error: %s\n", scanner.Text(), err.Error())
			return
		}
	}
}

func mockChatServer() {
	http.HandleFunc("/broadcastToChatRoom", func(rw http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userId")
		channel := r.URL.Query().Get("channel")
		m := r.URL.Query().Get("msg")

		//get this from backend
		secret := "secret"
		if userId == "" || channel == "" || secret == "" || m == "" {
			fmt.Printf("something is null: '%s' '%s' '%s' '%s'\n", userId, channel, secret, m)
			return
		}

		c := gocent.NewClient("http://localhost:8000", secret, 5*time.Second)
		msg := message{
			From:      userId,
			ToChannel: channel,
			Text:      m,
		}
		j, _ := json.Marshal(msg)
		_, err := c.Publish(channel, j)
		if err != nil {
			fmt.Printf("error sending, msg: %v, err: %s\n", msg, err)
		}
	})

	http.ListenAndServe(":1234", http.DefaultServeMux)
}

func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	stopListening <- true
}

func main() {
	go mockChatServer()

	userId := "1002"
	roomName := "hede"
	//get these from server
	roomUsers := "1001,1002,1003"
	secret := "secret"
	channel := fmt.Sprintf("%s#%s", roomName, roomUsers)
	go broadcastToChatRoom(userId, channel)
	go listenChatRoom(userId, channel, secret)
	waitForSignal()
}
