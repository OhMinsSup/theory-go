package scrapper

import (
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
)

type profile struct {
	id        string
	thumbnail string
}

type user struct {
	id       string
	username string
	profile  profile
}

type post struct {
	id               string
	title            string
	shortDescription string
	thumbnail        string
	user             user
	urlSlug          string
	likes            int
	commentsCount    string
	releasedAt       string
}

const nextPageURL = "https://v2.velog.io/graphql"
const nextPageVariable = `{"limit":"%s"}`

type trendingPageInfo struct {
	offset    int
	limit     int
	timeFrame string
}

type postsPageInfo struct {
	cursor   string
	limit    int
	username string
	tempOnly bool
	tag      string
}

// PostListScrape scraping data from posts
func PostListScrape(tranding bool) {
	if tranding {
		trandingScrape()
		return
	}
	postsScrape()
}

func trandingScrape() {

}

func postsScrape() {
	//outputDir := fmt.Sprintf("./velog/")

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"),
	)

	c.OnRequest(func(req *colly.Request) {
		req.Headers.Set("Referer", "https://velog.io/recent")
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	c.OnResponse(func(res *colly.Response) {
		log.Println("Write file:", ioutil.WriteFile("data.json", res.Body, 0600))
	})

	c.OnHTML("main > sc-Rmtcm", func(el *colly.HTMLElement) {

	})

	// 해당 사이트로 방문
	c.Visit("https://velog.io/recent")
}
