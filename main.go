package main

import (
	"io"
	"net"
	"os"
	"sync"
	"time"
)

// Config holds the proxy configuration
type Config struct {
	ListenAddr string
	TargetAddr string
	BufferSize int
	Timeout    time.Duration
}

// TCPProxy manages the TCP proxy operations
type TCPProxy struct {
	config     Config
	bufferPool sync.Pool
}

// NewTCPProxy creates a new TCP proxy with the given configuration
func NewTCPProxy(config Config) *TCPProxy {
	return &TCPProxy{
		config: config,
		bufferPool: sync.Pool{
			New: func() interface{} {
				return make([]byte, config.BufferSize)
			},
		},
	}
}

// Start begins listening for connections and proxying them
func (p *TCPProxy) Start() error {
	listener, err := net.Listen("tcp", p.config.ListenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			continue
		}

		go p.handleConnection(clientConn)
	}
}

// handleConnection processes a single client connection
func (p *TCPProxy) handleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	if p.config.Timeout > 0 {
		clientConn.SetDeadline(time.Now().Add(p.config.Timeout))
	}

	targetConn, err := net.DialTimeout("tcp", p.config.TargetAddr, 10*time.Second)
	if err != nil {
		return
	}
	defer targetConn.Close()

	if p.config.Timeout > 0 {
		targetConn.SetDeadline(time.Now().Add(p.config.Timeout))
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Client -> Target
	go func() {
		defer wg.Done()
		buf := p.bufferPool.Get().([]byte)
		defer p.bufferPool.Put(buf)
		
		io.CopyBuffer(targetConn, clientConn, buf)
		targetConn.(*net.TCPConn).CloseWrite()
	}()

	// Target -> Client
	go func() {
		defer wg.Done()
		buf := p.bufferPool.Get().([]byte)
		defer p.bufferPool.Put(buf)
		
		io.CopyBuffer(clientConn, targetConn, buf)
		clientConn.(*net.TCPConn).CloseWrite()
	}()

	wg.Wait()
}

func main() {
	listenAddr := ":" + os.Getenv("PORT")
	targetAddr := os.Getenv("V2RAY_SERVER_IP") + ":80"

	config := Config{
		ListenAddr: listenAddr,
		TargetAddr: targetAddr,
		BufferSize: 32 * 1024,       // 32KB buffer
		Timeout:    5 * time.Minute, // 5 minute timeout
	}

	proxy := NewTCPProxy(config)
	proxy.Start()
}
