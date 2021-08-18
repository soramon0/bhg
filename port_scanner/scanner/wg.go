package scanner

import (
	"fmt"
	"net"
	"sync"
)

func WgScanner(host string, ports int) {
	var wg sync.WaitGroup

	for i := 1; i <= ports; i++ {
		wg.Add(1)

		go func (j int) {
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", host, j)
			conn, err := net.Dial("tcp", address)

			if err != nil {
				return
			}
			
			fmt.Printf("%d open \n", j)
			conn.Close()
		}(i)
	}

	wg.Wait()
}