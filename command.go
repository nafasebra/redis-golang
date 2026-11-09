package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"redis-golang/store"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatcher(store store.Storer, c Command) {
	switch strings.ToUpper(c.Op) {
	case "GET":
		value, err := store.Get(c.Key)
		log.Printf("dispatcher operation=GET key=%q value=%q error=%v", c.Key, value, err)
	case "SET":
		err := store.Set(c.Key, c.Value)
		log.Printf("dispatcher operation=SET key=%q error=%v", c.Key, err)
	case "UPDATE":
		if _, err := store.Get(c.Key); err != nil {
			log.Printf("dispatcher operation=UPDATE key=%q error=%v", c.Key, err)
		}
		err := store.Set(c.Key, c.Value)
		log.Printf("dispatcher operation=UPDATE key=%q error=%v", c.Key, err)
	case "DELETE":
		err := store.Delete(c.Key)
		log.Printf("dispatcher operation=DELETE key=%q error=%v", c.Key, err)
	default:
		err := fmt.Errorf("unknown command: %s", c.Op)
		log.Printf("dispatcher operation=%q error=%v", c.Op, err)
	}
}

func RestoreOnBoot(st store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	wg.Add(len(cmds))

	for _, c := range cmds {
		go func(c Command) {
			defer wg.Done()
			dispatcher(st, c)
		}(c)
	}

	wg.Wait()
}
