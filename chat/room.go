package main

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"
	"log"
	"net/http"
)

type room struct {
	forward chan *message
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	tracer  Tracer
	avatar  Avatar
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true

		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("클라이언트가 떠났습니다.")

		case msg := <-r.forward:
			r.tracer.Trace("메시지를 받았습니다: ", msg.Message)
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- 클라이언트에 전송")
			}
		}
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	socket, err := upgrader.Upgrade(w, request, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	authCookie, err := request.Cookie("auth")
	if err != nil {
		log.Fatal("Failed to get auth cookie:", err)
		return
	}

	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}

	r.join <- client

	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

func newRoom(avatar Avatar) *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  Off(),
		avatar:  avatar,
	}
}
