package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) *client {
	return &client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *client) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	fmt.Fprintf(os.Stderr, "...Connected to %s\n", c.address)
	c.conn = conn

	return nil
}

func (c *client) Close() error {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("closer: %w", err)
		}
	}
	return nil
}

func (c *client) Send() error {
	if _, err := io.Copy(c.conn, c.in); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	fmt.Fprintln(os.Stderr, "...EOF")
	return nil
}

func (c *client) Receive() error {
	if _, err := io.Copy(c.out, c.conn); err != nil {
		return fmt.Errorf("receive: %w", err)
	}
	fmt.Fprintln(os.Stderr, "...Connection was closed")
	return nil
}

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalln("Invalid flag")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	client := NewTelnetClient(net.JoinHostPort(host, port), *timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Fatalln("Failed connection: ", err)
	}
	defer client.Close()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	go func() {
		client.Send()
		cancel()
	}()
	go func() {
		client.Receive()
		cancel()
	}()

	<-ctx.Done()
}
