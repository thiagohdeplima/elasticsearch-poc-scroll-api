package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

var items int

var wg sync.WaitGroup

func main() {
	flag.IntVar(&items, "n", 20000, "total items to save in elastic")
	flag.Parse()

	log.Print("waiting for elasticsearch startup")
	time.Sleep(time.Second * 30)

	for i := 1; i <= items; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			person, _ := GeneratePerson()

			log.Printf("saving person %d: %s", i, person.IDNumber)

			if err := SaveInElastic(person); err != nil {
				log.Fatalf("aborting: `%s`", err)
			}
		}()

		wg.Wait()
	}
}
