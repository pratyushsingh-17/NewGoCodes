package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)
var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	weblist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range weblist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

	// go greeter("Hello")
	// greeter("World")
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string)  {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err!=nil {
		fmt.Println("OOPS in endpoint")
	}	else{
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}