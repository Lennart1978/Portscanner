package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/Lennart1978/Portscanner/misc"
)

func main() {

	var (
		ports_open int
		portso     []int
	)

	greeting()

	for {
		fmt.Printf("Hello %s please enter target host name (or q to quit):", misc.GetUserName())
		host := misc.Input()

		if host == "q" || host == "Q" {
			break
		}

		fmt.Printf("Enter number of ports to scan:")
		p := misc.Input()

		port, _ := strconv.Atoi(p)

		totalPorts := port + 1
		portsScanned := 0

		fmt.Print("Progress: 0.00%") // Initialanzeige

		for count := 0; count <= port; count++ {
			if Portscan(host, count) {
				fmt.Printf("\033[1;42m --> Port: %d is open ! ", count)
				fmt.Println("\033[0m")
				ports_open++
				portso = append(portso, count)
			}

			portsScanned++
			progress := float64(portsScanned) / float64(totalPorts) * 100
			fmt.Printf("\rProgress: %.2f%%", progress) // Zeile überschreiben
		}

		if ports_open > 1 {
			fmt.Printf("\n\033[31m%d ports are open:\n", ports_open)
		} else if ports_open == 1 {
			fmt.Printf("\n\033[31m%d port is open:\n", ports_open)
		} else {
			fmt.Printf("\n\033[31mNo ports are open !\n")
		}

		fmt.Println(portso)
		fmt.Println("\033[0m")
		portso = nil
		ports_open = 0
	}
}

func Portscan(host string, ip int) bool {
	target := host + ":" + strconv.Itoa(ip)
	_, err := net.DialTimeout("tcp", target, time.Millisecond*30)
	return err == nil
}

func greeting() {
	fmt.Println("\033[34mWelcome to Lennart's Portscanner V1.0")
	fmt.Println("\033[34m-------------------------------------")
	fmt.Println("\033[0m")
}
