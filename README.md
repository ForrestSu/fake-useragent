# fake-useragent

Go library A wide variety of random useragent.

## Support

- All User-Agent Random
- Chrome
- Edge
- Firefox
- Safari

## Installation

```
$ go get github.com/ForrestSu/fake-useragent
```

## Usage

``` go
package main

import (
	"log"

	browser "github.com/ForrestSu/fake-useragent"
)

func main() {
	// recommend to use
	random := browser.Random()
	log.Printf("Random: %s", random)

	chrome := browser.Chrome()
	log.Printf("Chrome: %s", chrome)
}
```

## Output

``` sh
Random: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36

Chrome: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36
```