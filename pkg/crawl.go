package crawl

import (
	"log"
	"net/http"
	neturl "net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Options struct {
	Domain    *string // target domain
	Whitelist *string // domains to be whitelisted (Comma separated)
	Ext       *string // filter results on url extension
	UserAgent *string // the user agent
	Outfile   *string // file to write results
	Regexp    *string // regular expression to match against domains
	Parallels *int    // max parallels (default: 10)
	Depth     *int    // maximum depth to crawl
	Delay     *int    // seconds to wait before creating a new request to the matching domains
	Delay2    *int    // seconds to be randomized prior to each new request
	Async     *bool   // asynchronous network request
	JSON      *bool   // write to file as JSON
	Silent    *bool   // suppress output from console
	Version   *bool   // print version
	Help      *bool   // print help
}

func Go(options Options) (results []string) {
	var url string
	var whitelist []string

	resp, _ := http.Head("http://" + *options.Domain)
	url = resp.Request.URL.Scheme + "://" + resp.Request.URL.Host

	whitelist = strings.Split(*options.Whitelist, ",")
	whitelist = append(whitelist, *options.Domain)
	whitelist = append(whitelist, resp.Request.URL.Host)

	c := colly.NewCollector(
		colly.Async(*options.Async),
		colly.MaxDepth(*options.Depth),
		colly.AllowedDomains(whitelist...),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: *options.Parallels,
		Delay:       time.Duration(*options.Delay) * time.Second,
		RandomDelay: time.Duration(*options.Delay2) * time.Second,
	})

	if *options.UserAgent != "" {
		c.UserAgent = *options.UserAgent
	}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("src"))
	})

	if *options.Ext != "" {
		c.OnHTML("script[src]", func(e *colly.HTMLElement) {
			e.Request.Visit(e.Attr("src"))
		})
	}

	c.OnRequest(func(r *colly.Request) {
		p, _ := neturl.ParseRequestURI(r.URL.String())

		if p.Host != "" {
			if *options.Ext != "" {
				for _, ext := range strings.Split(*options.Ext, ",") {
					if strings.HasSuffix(r.URL.String(), "."+ext) {
						results = append(results, r.URL.String())
						if !*options.Silent {
							log.Println(r.URL)
						}
					}
				}
			} else if *options.Regexp != "" {
				re := regexp.MustCompile(*options.Regexp)
				if re.MatchString(r.URL.String()) {
					if !*options.Silent {
						log.Println(r.URL)
					} else {
						results = append(results, r.URL.String())
					}
				}
			} else {
				results = append(results, r.URL.String())
				if !*options.Silent {
					log.Println(r.URL)
				}
			}
		}
	})

	c.Visit(url)
	c.Wait()
	return results
}
