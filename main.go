package main

import (
	"fmt"
	"strconv"

	"github.com/Lennart1978/Portscanner/misc"
)

func main() {

	var Scanner Portscanner  // Create a Portscanner instance

	Scanner.Greet()  // Display greeting message

	for {  // Start an infinite loop to continually prompt the user for input
		fmt.Printf("Hello %s please enter target host name (or q to quit):", misc.GetUserName())  // Prompt for target host
		Scanner.Host = misc.Input()  // Get user input for target host

		if Scanner.Host == "q" || Scanner.Host == "Q" {  // Check if the user wants to quit
			break  // Exit the loop if the user enters 'q' or 'Q'
		}

		fmt.Printf("Enter number of ports to scan or 'all':")  // Prompt for number of ports to scan

		p := misc.Input()  // Get user input for number of ports

		if p == "all" {
			Scanner.Port = 65535  // Set to scan all ports if user enters 'all'
		} else {
			Scanner.Port, _ = strconv.Atoi(p)  // Convert user input to integer
			if Scanner.Port < 1 || Scanner.Port > 65535 {
				Scanner.Port = 65535  // Set to scan all ports if user input is out of range
			}
		}

		Scanner.totalPorts = Scanner.Port + 1  // Set the total number of ports to scan

		fmt.Print("Progress: 0.00%")  // Initial progress display

		for Scanner.scanningPort < Scanner.totalPorts {  // Loop through all ports to be scanned
			Scanner.Port = Scanner.scanningPort  // Set the current port to scan
			if Scanner.Scan() {
				Scanner.Print()  // Print the port if it's open
				Scanner.ports_open++  // Increment the count of open ports
				Scanner.portso = append(Scanner.portso, Scanner.scanningPort)  // Add the open port to the list
			}

			Scanner.Pprogress()  // Update and display the scanning progress
			Scanner.scanningPort++  // Move to the next port
			Scanner.portsScanned++  // Increment the count of scanned ports
		}

		// Display the results based on the number of open ports found
		if Scanner.ports_open > 1 {
			fmt.Printf("\n\033[31m✅ %d ports are open:\n", Scanner.ports_open)
		} else if Scanner.ports_open == 1 {
			fmt.Printf("\n\033[31m✅ %d port is open:\n", Scanner.ports_open)
		} else {
			fmt.Printf("\n\033[31m⛔ No ports are open !\n")
		}

		fmt.Println(Scanner.portso)  // Print the list of open ports
		fmt.Println("\033[0m")  // Reset terminal color

		Scanner.Reset()  // Reset the scanner for the next iteration
	}
}
