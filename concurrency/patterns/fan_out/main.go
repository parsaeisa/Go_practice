package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	ch1, err := readCSV("words.csv")
	if err != nil {
		log.Fatalf("error in opening file %v", err.Error())
	}

	br1 := breakup("1", ch1)
	br2 := breakup("2", ch1)
	br3 := breakup("3", ch1)

	for {
		if br1 == nil && br2 == nil && br3 == nil {
			break
		}

		select {
		case _, ok := <-br1:
			if !ok {
				br1 = nil
			}
		case _, ok := <-br2:
			if !ok {
				br2 = nil
			}
		case _, ok := <-br3:
			if !ok {
				br3 = nil
			}
		}
	}

}

func breakup(worker string, ch <-chan []string) chan struct{} {
	chE := make(chan struct{})

	go func() {
		for v := range ch {
			fmt.Println(worker, v)
		}
		close(chE)
	}()

	return chE
}

func readCSV(path string) (<-chan []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	ch := make(chan []string)
	reader := csv.NewReader(file)

	go func() {
		for {
			record, err := reader.Read()

			if err == io.EOF {
				close(ch)
				break
			}
			//if err != nil {
			//
			//}

			ch <- record
		}
	}()

	return ch, nil

}
