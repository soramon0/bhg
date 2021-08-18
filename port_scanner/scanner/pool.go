package scanner

import (
	"fmt"
	"net"
	"sort"
)

func worker(host string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}
		
		conn.Close()
		results <- p
	}
}

func WorkerScanner(host string, ports, workerCount int) {
	workerPorts := make(chan int, workerCount)
	workerResults := make(chan int)
	var openPorts []int

	for i:= 0; i < cap(workerPorts); i++ {
		go worker(host, workerPorts, workerResults)
	}

	go func() {
		for i := 1; i <= ports; i++ {
			workerPorts <- i
		}
	}()

	for i := 1; i <= ports; i++ {
		port := <- workerResults
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(workerPorts)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}