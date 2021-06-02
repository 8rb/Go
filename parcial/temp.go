package main

import (
	"fmt"
	"sync"
)

func writer(roomEmpty sync.Mutex) {
	for {
		roomEmpty.Lock()
		fmt.Println("Writing in CS")
		roomEmpty.Unlock()
	}
}

func reader(readers int, mutex, roomEmpty sync.Mutex) {
	for {
		mutex.Lock()
		readers++
		if readers == 1 {
			roomEmpty.Lock()
		}
		mutex.Unlock()
		fmt.Println("Reading CS")
		mutex.Lock()
		readers--
		if readers == 0 {
			roomEmpty.Unlock()
		}
		mutex.Unlock()
	}
}

func main() {
	readers := 0
	mutex := &sync.Mutex{}
	roomEmpty := &sync.Mutex{}

	go writer(*roomEmpty)
	go reader(readers, *mutex, *roomEmpty)
	reader(readers, *mutex, *roomEmpty)
}
