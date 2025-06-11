package wt

import (
	"fmt"
	"log"
)

func (c *Connection) OpenStream() int64 {
	str, err := c.Session.OpenStream()
	if err != nil {
		log.Println("Stream error:", err.Error())
	}
	c.activeStream = str
	c.streams[int64(str.StreamID())] = str

	return int64(str.StreamID())
}

func (c *Connection) CloseStream() {
	if c.activeStream != nil {
		err := c.activeStream.Close()
		if err != nil {
			fmt.Printf("Encountered error while closing stream: %v\n", err.Error())
			return
		}
		delete(c.streams, int64(c.activeStream.StreamID()))
		c.activeStream = nil
	}
}

func (c *Connection) CloseStreamById(id int64) {
	if c.streams[id] != nil {
		err := c.streams[id].Close()
		if err != nil {
			fmt.Printf("Encountered error closing stream id %d, %v\n", id, err)
			return
		}
		delete(c.streams, id)
		if c.activeStream != nil && int64(c.activeStream.StreamID()) == id {
			c.activeStream = nil
		}
	}
}

func (c *Connection) CloseAllStreams() {
	for id := range c.streams {
		c.CloseStreamById(id)
	}
}

func (c *Connection) SetActiveStream(id int64) {
	c.activeStream = c.streams[id]
}
