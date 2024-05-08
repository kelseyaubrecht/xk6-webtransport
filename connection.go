package wt

import (
	"context"
	"log"

	"github.com/quic-go/webtransport-go"
	"go.k6.io/k6/js/modules"
)

type Connection struct {
	vu           modules.VU
	Session      *webtransport.Session
	activeStream webtransport.Stream
	metrics      WTMetrics
	readBuffer   [][]byte
	streams      map[int64]webtransport.Stream
}

func (c *Connection) Connect(url string) {
	var dialer webtransport.Dialer
	_, sess, err := dialer.Dial(context.Background(), url, nil)
	if err != nil {
		log.Println("Error: " + err.Error())
		return
	}
	c.Session = sess

	c.streams = make(map[int64]webtransport.Stream)
}

func (c *Connection) Close() {
	if c.activeStream != nil {
		c.CloseAllStreams()
	}
	err := c.Session.CloseWithError(0, "")
	if err != nil {
		log.Println("Error: " + err.Error())
	}
}
