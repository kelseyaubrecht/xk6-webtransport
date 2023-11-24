package wt

func (c *Connection) ReadStream(maxReadLength int) {
	for {
		rsp := make([]byte, maxReadLength)

		n, err := c.Stream.Read(rsp)
		if err != nil {
			break
		}

		if n > 0 {
			defer c.logReadMetrics(n)
			subArray := rsp[0:n]
			c.readBuffer = append(c.readBuffer, subArray)
		}
	}
}

func (c *Connection) StartReadToBuffer(maxReadLength int) {
	go c.ReadStream(maxReadLength)
}

func (c *Connection) ReadBuffer() [][]byte {
	return c.readBuffer
}
