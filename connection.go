package wt

import (
	"context"
	"log"
	"time"

	"github.com/quic-go/webtransport-go"
	"go.k6.io/k6/js/modules"
)

type Connection struct {
	vu           modules.VU
	Session      *webtransport.Session
	Stream       webtransport.Stream
	metrics      WTMetrics
	requestTimes []time.Time
	readBuffer   [][]byte
}

func (c *Connection) Connect(url string) {
	var dialer webtransport.Dialer
	_, sess, err := dialer.Dial(context.Background(), url, nil)
	if err != nil {
		log.Println("Error: " + err.Error())
		return
	}
	c.Session = sess

	str, err := c.Session.OpenStream()
	if err != nil {
		log.Println("Stream error: " + err.Error())
	}
	c.Stream = str
}

func (c *Connection) Close() {
	if c.Stream != nil {
		c.Stream.Close()
	}
	c.Session.CloseWithError(0, "")
}
