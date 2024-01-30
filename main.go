package main

import (
	"bufio"
	"crawler/crawler"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	url := ""

	for {
		fmt.Print("> Input url: ")
		text, _ := reader.ReadString('\n')

		url = strings.Replace(text, "\n", "", -1)
		break
	}

	fmt.Printf("> Crawling %s. This might take a while...\n", url)

	crawler := crawler.New(url)
	if err := crawler.Crawl(); err != nil {
		fmt.Println(err)
	}
}
