package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err != nil {
			return
		}

		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
		c.room.forward <- msg
	}
}
