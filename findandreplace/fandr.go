package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, err := os.Open("./L1.txt")
	if err != nil {
		fmt.Println("Unable to create file L1:", err)
		os.Exit(1)
	}

	var L1 []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		L1 = append(L1, scanner.Text())
	}

	file, err = os.Open("./L2.txt")
	if err != nil {
		fmt.Println("Unable to create file L2:", err)
		os.Exit(1)
	}

	var L2 []string
	var L3 []string
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		L2 = append(L2, scanner.Text())
		L3 = append(L3, scanner.Text())
	}

	for _, l1 := range L1 {
		fmt.Println(l1)
		regexpL1, err := regexp.Compile("(" + l1 + ")")
		if err != nil {
			log.Fatal(err)
		}
		for i, l2 := range L2 {
			if err == nil && regexpL1.MatchString(l2) {
				updatedText := strings.ReplaceAll(l2, l1, "")
				L3[i] = updatedText
			}
		}
	}

	file, err = os.Create("./L3.txt")
	if err != nil {
		fmt.Println("Unable to create file L3:", err)
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range L3 {
		fmt.Fprintln(w, line)
	}
	w.Flush()

}
