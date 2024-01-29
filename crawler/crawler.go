package crawler

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

type Crawler interface {
	Crawl() error
}

func isCSR(url string, timeout ...time.Duration) bool {
	duration := 30 * time.Second
	if len(timeout) > 0 {
		duration = timeout[0]
	}

	ctx, _ := context.WithTimeout(context.Background(), duration)
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("> CSR website detected. Crawling with CSR strategy")
			return true
		default:
			err := chromedp.Run(ctx,
				chromedp.Navigate(url),
				chromedp.WaitVisible("body"),
			)

			if err != nil {
				fmt.Println("> CSR website detected. Crawling with CSR strategy")
				return true
			}

			fmt.Println("> SSR website detected. Crawling with SSR strategy")
			return false
		}
	}
}

func New(ctx context.Context, url string) Crawler {
	if isCSR(url) {
		return &csrCrawler{ctx, url}
	}

	return &ssrCrawler{ctx, url}
}
