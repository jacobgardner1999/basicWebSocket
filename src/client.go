// client.go
package main

import (
    "log"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/websocket"
)

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
    defer close(done)
    for { 
