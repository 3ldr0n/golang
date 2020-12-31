package channels

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func writeFiles() {
	os.Mkdir("files", os.ModePerm)
	for i := 0; i < 100; i++ {
		data := []byte("data")
		ioutil.WriteFile("files/"+strconv.Itoa(i)+".txt", data, os.ModePerm)
	}
}

func readFiles(queue chan string) {
	files, _ := ioutil.ReadDir("files")
	for _, file := range files {
		queue <- file.Name()
	}
}

func RunChannels() {
	writeFiles()

	files := make(chan string)

	go func() {
		readFiles(files)
	}()

	for {
		select {
		case file := <-files:
			fmt.Println(file)
		default:
			os.Exit(0)
		}
	}
}
