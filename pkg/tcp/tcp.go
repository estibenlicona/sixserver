package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sixservergo/pkg/dispatcher"
	"sixservergo/pkg/protocols/pes6"
	"strconv"
	"sync/atomic"
)

const maxConcurrentConnections = 100

var semaphore = make(chan struct{}, maxConcurrentConnections)

// Modificar Start para recibir el dispatcher
func Start(port int, dispatcher *dispatcher.Dispatcher) error {
	log.Println("Starting Checking Service (TCP) on port:", port)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return fmt.Errorf("failed to start service on TCP Port %v: %v", port, err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		semaphore <- struct{}{}
		go func() {
			defer func() { <-semaphore }()
			handlerConnection(conn, dispatcher)
		}()
	}
}

func handlerConnection(conn net.Conn, dispatcher *dispatcher.Dispatcher) {
	defer conn.Close()
	var recvPacketCount uint32 = 1
	var sendPacketCount uint32 = 1
	dataReceived(conn, &recvPacketCount, &sendPacketCount, dispatcher)
}

func dataReceived(conn net.Conn, recvPacketCount *uint32, sendPacketCount *uint32, dispatcher *dispatcher.Dispatcher) {
	reader := bufio.NewReader(conn)
	buffer := make([]byte, 0)

	for {
		data := make([]byte, 1024)
		bytes, err := reader.Read(data)
		if err != nil {
			log.Println("Error reading data:", err)
			break
		}

		buffer = append(buffer, data[:bytes]...)
		buffer = processCheckingPackets(buffer, conn, recvPacketCount, sendPacketCount, dispatcher)
	}
}

func processCheckingPackets(buffer []byte, conn net.Conn, recvPacketCount *uint32, sendPacketCount *uint32, dispatcher *dispatcher.Dispatcher) []byte {
	return processPackets(buffer, conn, recvPacketCount, sendPacketCount, dispatcher)
}

const packetIdSize = 8

func processPackets(buffer []byte, conn net.Conn, recvPacketCount *uint32, sendPacketCount *uint32, dispatcher *dispatcher.Dispatcher) []byte {
	for len(buffer) >= packetIdSize {
		packetHeader, err := pes6.ProcessPacketHeader(buffer, packetIdSize, recvPacketCount)
		log.Println("Packet Recibido", packetHeader)
		pes6.HandlerError(err)

		totalPacketSize := pes6.CalculateTotalPacketSize(packetHeader)
		if len(buffer) < totalPacketSize {
			return buffer
		}

		packet, err := pes6.ProcessCompletePacket(buffer, totalPacketSize, sendPacketCount)
		if pes6.HandlerError(err) {
			return buffer
		}

		dispatcher.DispatchPacket(conn, packet)
		atomic.AddUint32(recvPacketCount, 1)
		buffer = removeProcessedPacket(buffer, totalPacketSize)
	}

	return buffer
}

func removeProcessedPacket(buffer []byte, totalPacketSize int) []byte {
	if totalPacketSize <= len(buffer) {
		return buffer[totalPacketSize:]
	}
	return buffer[:0]
}
