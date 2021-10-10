package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	c    net.Conn
	mute chan struct{}
}

func NewClient(ip string, port int, timeout time.Duration) *Client {
	d := net.Dialer{
		Timeout: timeout,
	}
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	c, err := d.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	return &Client{c: c}
}

func (c *Client) reader(r io.Reader) {
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			println(err)
			return
		}
		println("Client got:\"", string(buf[0:n]), "\"")
		(*c).mute <- struct{}{}
	}
}

func (c *Client) ListenAndServe() {
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Println(err)
		}
	}((*c).c)

	go (*c).reader((*c).c)
	for {
		fmt.Println("Enter your message:")
		var text string
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			text = scanner.Text()
		}

		_, err := (*c).c.Write([]byte(text))
		if err != nil {
			log.Fatal("write error:", err)
		}
		<-(*c).mute
	}
}

func main() {
	var (
		host    = flag.String("host", "", "host to connect")
		port    = flag.Int("port", 80, "current port to listen")
		timeout = flag.Duration("timeout", time.Second*15, "timeout to connect")
	)

	if *host == "" {
		log.Fatal("please, use flag host to describe source host")
	}

	client := NewClient(*host, *port, *timeout)
	client.ListenAndServe()
}
