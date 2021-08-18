package main

import (
	"flag"
	"soramon0/port_scanner/scanner"
)

func main() {
	host := flag.String("host", "scanme.nmap.org", "host to scan")
	ports := flag.Int("ports", 1024, "ports to scan")
	scanMode := flag.String("mode", "worker", "running mode: either worker or wg")
	workersCount := flag.Int("workers", 100, "number of works to start")
	flag.Parse()

	if (*scanMode == "wg") {
		scanner.WgScanner(*host, *ports)
	} else {
		scanner.WorkerScanner(*host, *ports, *workersCount)
	}
}
