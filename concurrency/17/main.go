package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// read test files
// put custom objects in channel
// add it to channel
// use waitgroups + close channel
// read all from channel
// do total size
var dirName = "testfiles"

type CustomFile struct {
	name string
	size int64
}

func PanicIfError(err error, msg string) {
	if err != nil {
		panic(err)
	}
	fmt.Println("**", msg)
}

func main() {
	files, err := ioutil.ReadDir(dirName)
	PanicIfError(err, "Reading DIR is complete")
	ch := make(chan *CustomFile, len(files))
	out := make(chan int64)
	var wg sync.WaitGroup
	for _, fInfo := range files {
		wg.Add(1)
		go func(fInfo os.FileInfo) {
			ch <- &CustomFile{fInfo.Name(), fInfo.Size()}
			wg.Done()
		}(fInfo)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		var total int64
		for cf := range ch {
			total += cf.size
		}
		out <- total
		close(out)
	}()
	fmt.Println(<-out)
}
