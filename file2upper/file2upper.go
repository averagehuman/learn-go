package main

import "bufio"
import "fmt"
import "io"
import "os"
import "strings"

// Convert file (or stdin) to uppercase
//
// $ echo "abc" | ./file2upper
// ABC
//
func main() {

	var r io.Reader

	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// data is being piped to stdin
		r = os.Stdin
	} else {
		if len(os.Args) < 2 {
			fmt.Println("no input file")
			return
		}
		if _, err := os.Stat(os.Args[1]); err != nil {
			fmt.Println("ERROR: file not found")
			return
		}
		r, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer r.Close()
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println(strings.ToUpper(scanner.Text()))
	}
}
