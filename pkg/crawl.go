package pkg

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	colly "github.com/gocolly/colly/v2"
)

func RunCrawler() (string, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.tradingview.com", "tradingview.com"),
		colly.UserAgent("Mozilla/5.0 (compatible; JakeBot/1.0)"),
		colly.IgnoreRobotsTxt(),
	)

	var dxy string
	c.OnHTML("section[data-an-section-id=symbol-overview-page-section]", func(e *colly.HTMLElement) {
		substringToReplace := "The current value of U.S. Dollar Index is"
		expression := `(\d+\.\d+)`
		full := fmt.Sprintf("%s %s", substringToReplace, expression)
		target := regexp.MustCompile(full)

		match := target.FindString(e.Text)
		extracted := strings.Replace(match, substringToReplace, "", 1)
		log.Println("match: ", match, "extracted: ", extracted)

		dxy = extracted
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})
	return dxy, c.Visit("https://www.tradingview.com/symbols/TVC-DXY/")
}
