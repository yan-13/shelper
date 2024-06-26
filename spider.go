package shelper

import (
	"fmt"
	"time"
)

type Spider interface {
	GetTitle(body string) (title string, err error)
	GetArticle(body string) (title string, err error)
}

func GetUrlsAsTxt(spider Spider, urls []string, wait time.Duration) (txt string, err error) {
	for _, v := range urls {
		if v == "" {
			continue
		}

		var (
			httpRes string
			title   string
			article string
		)
		httpRes, err = HttpGet(v)
		if err != nil {
			return
		}

		//title
		title, err = spider.GetTitle(httpRes)
		if err != nil {
			return
		}

		//article
		article, err = spider.GetArticle(httpRes)
		if err != nil {
			return
		}
		txt += fmt.Sprintf("%s\n%s\n", title, article)
		if wait > 0 {
			time.Sleep(wait)
		}
	}
	return
}
