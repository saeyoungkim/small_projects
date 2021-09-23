package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("word and file name are needed...")
	}

	word, paths := os.Args[1], os.Args[2:]

	wg := sync.WaitGroup{}
	ch := make(chan []string, len(paths))

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go searchWord(&wg, file, word, ch)
	}

	wg.Wait()
	close(ch)

	for lines := range ch {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

func searchWord(wg *sync.WaitGroup, file *os.File, word string, ch chan []string) {
	rd := bufio.NewReader(file)

	fmt.Printf("start to read %s ...\n", file.Name())

	ret := []string{"====================", file.Name(), "--------------------"}

	tmp := ""
	row := 1
	for line, isPrefix, err := rd.ReadLine(); line != nil && err == nil; line, isPrefix, err = rd.ReadLine() {
		tmp += string(line)
		if !isPrefix {
			// insert string
			if strings.Contains(tmp, word) {
				ret = append(ret, fmt.Sprintf("'%12s'    %s", strconv.Itoa(row), tmp))
			}
			// increment row number
			tmp = ""
			row++
		}
	}

	ret = append(ret, "--------------------\n")

	ch <- ret
	wg.Done()
}
