package main
import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(ip string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err == nil {
		fmt.Printf("Port %d is open\n", port)
		conn.Close()
	}
}

func main() {
	var wg sync.WaitGroup
	ip := "10.1.11.31"
	startPort, endPort := 1, 9999

	fmt.Printf("Scanning ports on %s.....\n", ip)
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort(ip, port, &wg)
	}

	wg.Wait()
	fmt.Println("Scan Complete!")
}