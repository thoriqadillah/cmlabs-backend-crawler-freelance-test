package crawler

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly/v2"
)

type collyCrawler struct {
	ctx context.Context
	url string
}

func newColly(ctx context.Context, url string) Crawler {
	return &collyCrawler{ctx, url}
}

func (c *collyCrawler) Crawl() error {
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	var content string
	var err error
	collector.OnHTML("html", func(h *colly.HTMLElement) {
		html, err := h.DOM.Html()
		if err != nil {
			fmt.Println("error fetching page html:", err)
			return
		}

		content = html
	})

	err = collector.Visit(c.url)
	if err != nil {
		fmt.Println("error navigating to url:", err)
		return err
	}

	url, _ := url.Parse(c.url)
	filename := fmt.Sprintf("./out/%s.html", url.Host)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	fmt.Printf("> Crawling finished. Result written on %s\n", filename)

	return nil
}

func init() {
	registerCrawler(CollyDriver, newColly)
}
