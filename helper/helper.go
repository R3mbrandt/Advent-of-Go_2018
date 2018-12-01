package helper

import (
	"bufio"
	"log"
	"os"
)

//ReadInput reads line by line from filename and returns a map of ints
func ReadInput(filename string) (out []string) {
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		log.Fatalln(err)
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return
}
