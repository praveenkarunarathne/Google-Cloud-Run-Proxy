package main

import (
    "io"
    "net"
    "os"
)

func handleClient(clientConn net.Conn, targetAddr string) {
    defer clientConn.Close()

    remoteConn, _ := net.Dial("tcp", targetAddr)
    defer remoteConn.Close()

    go func() {
        io.Copy(remoteConn, clientConn)
    }()

    io.Copy(clientConn, remoteConn)
}

func main() {
    listenAddr := ":" + os.Getenv("PORT")
    targetAddr := os.Getenv("V2RAY_SERVER_IP") + ":8081"
    listener, _ := net.Listen("tcp", listenAddr)
    defer listener.Close()

    for {
        clientConn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(clientConn, targetAddr)
    }
}
