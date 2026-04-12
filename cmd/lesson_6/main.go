package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var result = []string{}
var wg = sync.WaitGroup{}
var mutex = sync.RWMutex{}

func main() {
	t0 := time.Now()
	var i int
	for i = range dbData {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Printf("Total time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	time.Sleep(time.Second * 2)
	fmt.Printf("Value of db: %v\n", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}


func save(data string) {
	mutex.Lock()
	result = append(result, data)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Printf("Current result: %v\n", result)
	mutex.RUnlock()
}
