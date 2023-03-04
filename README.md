# crwl
A simple/fast web crawler written in Go

## Installation

Requires [Go](<https://golang.org/>) 17+ and git

```
go install github.com/root4loot/crwl@latest
```

## ☛ Usage

```
                              888 
 e88~~\ 888-~\ Y88b    e    / 888 
d888    888     Y88b  d8b  /  888 
8888    888      Y888/Y88b/   888 
Y888    888       Y8/  Y8/    888 
 "88__/ 888        Y    Y     888 


Arguments:

    -domain     <string>    Domain to crawl
    -whitelist  <string>    Domains to be whitelisted (comma separated)
    -ext        <string>    Filter results on URL file extensions (comma separated)
    -outfile    <string>    File to write results
    -regex      <string>    Regular expression to match against domains
    -useragent  <string>    Set user-agent
    -para       <int>       Max parallelism (Default: 10)
    -depth      <int>       Maximum depth to crawl (Default: 999)
    -delay      <int>       Seconds to wait before creating a new request to the matching domains
    -delay2     <int>       Seconds to be randomized prior to each new request (Default: 0)
    -json                   Output as JSON (-outfile)
    -silent                 Suppress output from console
    -async                  Enable asynchronous network requests (Default: ON)
    -version                Display version
    -help                   Display help
```

## Example

Parallelism
```
➜ crwl -domain=hackerone.com -depth=5 -parallels=12
https://www.hackerone.com
https://www.hackerone.com/attack-resistance-assessment
https://www.hackerone.com/product/attack-surface-management
https://www.hackerone.com/6th-annual-hacker-powered-security-report
https://www.hackerone.com/events/rsa-conference-2023
https://www.hackerone.com/security-at-beyond

...
```

Filter result on file extension type
```
➜ crwl -domain=hackerone.com -depth=5 -ext=js,cjs,ts
https://www.hackerone.com/sites/default/files/js/js_Ikd9nsZ0AFAesOLgcgjc7F6CRoODbeqOn7SVbsXgALQ.js
https://www.hackerone.com/sites/default/files/js/js_C-5Xm0bH3IRZtqPDWPr8Ga4sby1ARHgF6iBlpL4UHao.js
https://www.hackerone.com/sites/default/files/js/js_qMhgwN8eW12xYBuYp72UCEgsf2VpdEL1oEsgYy-41uQ.js
https://www.hackerone.com/sites/default/files/js/js_frWAQWhKej-kLg0hDQJ34LJ4H4rz7j5907BPIOKGmU0.js
https://www.hackerone.com/sites/default/files/js/js_zApVJ5sm-YHSWP4O5K9MqZ_6q4nDR3MciTUC3Pr1ogA.js
https://www.hackerone.com/sites/default/files/js/js_2nrQvMeN7-UiDTpvJ3GUwcaoYW3QOwfPQPUnrZyRus8.js

...
```

Filter on regular expression
```
➜ crwl -domain=hackerone.com -whitelist=www.hackerone.com -regex=\/default\/files\/
https://www.hackerone.com/sites/default/files/styles/logo_band_white/public/zoom-logo.png.webp?itok=u6jMGMbY
https://www.hackerone.com/sites/default/files/styles/carousel_tabbed/public/HAC-ARM-Product-1-L1R1%402x.png.webp?itok=QulSSLb0
https://www.hackerone.com/sites/default/files/styles/logo_band_white/public/Nintendo2X_0.png.webp?itok=h2H_a61-
https://www.hackerone.com/sites/default/files/styles/hero_main/public/HERO_%20%281%29%202.png.webp?itok=eWJYJ0v_
https://www.hackerone.com/sites/default/files/styles/testimonial_company_2x/public/GoodRx2X.png.webp?itok=gTPC1ATy
https://www.hackerone.com/sites/default/files/styles/logo_band_white/public/GM2X_0.png.webp?itok=_mTFOxRw
https://www.hackerone.com/sites/default/files/styles/logo_band_white/public/PayPal2X.png.webp?itok=UJfM8O9a
https://www.hackerone.com/sites/default/files/styles/testimonial_person/public/Kevin_Pawloski_GoodRX2X.png.webp?itok=2DT_nE29
https://www.hackerone.com/sites/default/files/styles/carousel_tabbed/public/Carosel_PickSolution%20%281%29%202.png.webp?itok=9RY31jLo
https://www.hackerone.com/sites/default/files/styles/testimonial_company_2x/public/Hired2X_0.png.webp?itok=qBm6wfuq

...
```

Silence and write to file
```
➜ crwl -domain=hackerone.com -depth=5 --silent --outfile results.txt
➜ 
```

## As lib

```go
package main

import (
	"fmt"

	crwl "github.com/root4loot/crwl/pkg"
)

func main() {

	domain    := "hackerone.com"
	whitelist := "www.hackerone.com,example.com"
	ext       := ""
	regexp    := ""
	useragent := ""
	parallels := 10
	depth     := 3
	delay     := 0
	delay2    := 0
	async     := true
	silent    := true

	options := crwl.Options{
		Domain:    &domain,
		Whitelist: &whitelist,
		Ext:       &Ext,
		Parallels: &parallels,
		Regexp:    &regexp,
		Depth:     &depth,
		Delay:     &delay,
		Delay2:    &delay2,
		UserAgent: &useragent,
		Async:     &async,
		Silent:    &silent,
	}

	results := crwl.Go(options)
	fmt.Println(results)
}

```

---

Made possible with [Colly](https://github.com/gocolly)

[@danielantonsen](https://twitter.com/danielantonsen)


