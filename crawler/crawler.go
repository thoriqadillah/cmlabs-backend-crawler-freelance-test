package crawler

import (
	"context"
)

type Crawler interface {
	Crawl() error
}

type CrawlerImpl func(ctx context.Context, url string) Crawler

var instances = make(map[string]CrawlerImpl)

func registerCrawler(name string, impl CrawlerImpl) {
	instances[name] = impl
}

type CrawlerDriver = string

const (
	CollyDriver CrawlerDriver = "colly"
)

type option struct {
	driver CrawlerDriver
}

type CrawlerOption func(o *option)

func UseDriver(driver CrawlerDriver) CrawlerOption {
	return func(o *option) {
		o.driver = driver
	}
}

func New(url string, options ...CrawlerOption) Crawler {
	return NewWithContext(context.Background(), url, options...)
}

func NewWithContext(ctx context.Context, url string, options ...CrawlerOption) Crawler {
	defaultOption := &option{
		driver: CollyDriver,
	}

	for _, option := range options {
		option(defaultOption)
	}

	crawler := instances[defaultOption.driver]
	return crawler(ctx, url)
}
