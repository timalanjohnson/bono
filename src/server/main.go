package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	go handleMessages()

	http.HandleFunc("/ws", handleConnections)

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServe error:", err)
		}
	}()

	address := ":20777"
	fmt.Printf("Attempting to resolve %s\n", address)
	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("UDP server listening on %s\n", address)
	fmt.Println("WebSocket server listening on port :8080")

	buffer := make([]byte, (1024 * 4))
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
			continue
		}

		go parsePacket(buffer[:n])
	}
}

func parsePacket(packetBytes []byte) {
	header := parsePacketHeader(packetBytes)
	switch header.PacketID {
	case 0:
		motionData := parsePacketMotionData(packetBytes)
		message, _ := json.Marshal(motionData)
		broadcast <- message
	}
}

func parsePacketHeader(packetBytes []byte) PacketHeader {
	header := PacketHeader{}
	err := binary.Read(bytes.NewReader(packetBytes), binary.LittleEndian, &header)
	if err != nil {
		log.Printf("Failed to decode PacketHeader: %s", err)
	}
	return header
}

func parsePacketMotionData(packetBytes []byte) PacketMotionData {
	motionData := PacketMotionData{}
	err := binary.Read(bytes.NewReader(packetBytes), binary.LittleEndian, &motionData)
	if err != nil {
		log.Printf("Failed to decode PacketMotionData: %s", err)
	}
	return motionData
}

func handleMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Printf("write error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
	}
}
