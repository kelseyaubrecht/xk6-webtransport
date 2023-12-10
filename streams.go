package wt

import (
	"log"
)

func (c *Connection) OpenStream() int64 {
	str, err := c.Session.OpenStream()
	if err != nil {
		log.Println("Stream error: " + err.Error())
	}
	c.activeStream = str
	c.streams[int64(str.StreamID())] = str

	return int64(str.StreamID())
}

func (c *Connection) CloseStream() {
	if c.activeStream != nil {
		c.activeStream.Close()
		delete(c.streams, int64(c.activeStream.StreamID()))
		c.activeStream = nil
	}
}

func (c *Connection) CloseStreamById(id int64) {
	if c.streams[id] != nil {
		c.streams[id].Close()
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
