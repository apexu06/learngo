package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func check_error(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// grape file.txt Hallo
// go run . -- file.txt Hallo
// Output: "5: Hallo wie geht es dir?"
func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Must provide 2 arguments")
		println("Must provide 2 arguments")
		os.Exit(1)
	}

	file_name := os.Args[1]
	search_string := os.Args[2]

	current_dir, err := os.Getwd()
	check_error(err)

	path := filepath.Join(current_dir, file_name)
	data, err := os.ReadFile(path)
	check_error(err)
	lines := strings.Split(string(data), "\n")

	found_lines := make(map[int]string)

	for line_number, line := range lines {
		if strings.Contains(line, search_string) {
			found_lines[line_number+1] = line
		}
	}

	keys := make([]int, 0, len(found_lines))
	for k := range found_lines {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Fprintf(os.Stdout, "Found %d lines with %s: \n", len(found_lines), search_string)

	for _, k := range keys {
		fmt.Fprintf(os.Stdout, "%d: %s\n", k, found_lines[k])
	}

}
