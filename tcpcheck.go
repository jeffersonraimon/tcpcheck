package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: tcpcheck <IP> <Port>")
		os.Exit(1)
	}

	ip := os.Args[1]
	portStr := os.Args[2]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port number")
		os.Exit(1)
	}

	for {
		start := time.Now()
		addr := fmt.Sprintf("[%s]:%d", ip, port)
		conn, err := net.DialTimeout("tcp", addr, time.Second)
		elapsed := time.Since(start)
		pingMs := elapsed.Seconds() * 1000

		pingColor := color.New(color.FgHiWhite).SprintFunc()
		if pingMs < 40 {
			pingColor = color.New(color.FgHiGreen).SprintFunc()
		} else if pingMs < 90 {
			pingColor = color.New(color.FgHiYellow).SprintFunc()
		} else {
			pingColor = color.New(color.FgHiRed).SprintFunc()
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05.000")

		if err == nil {
			conn.Close()
			fmt.Printf("[%s] Status: %s | Ping: %s | Port: %s | Protocol: %s\n",
				timestamp,
				color.HiGreenString("Connected"),
				pingColor(fmt.Sprintf("%.0fms", pingMs)),
				color.HiGreenString("%d", port),
				color.HiGreenString("TCP"),
			)
		} else {
			fmt.Printf("[%s] Status: %s - %s\n",
				timestamp,
				color.HiRedString("Could not connect"),
				color.HiRedString(err.Error()),
			)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
