package fake_useragent

import "github.com/ForrestSu/fake-useragent/useragent"

// major browsers
const (
	CHROME  = "chrome"
	EDGE    = "edge"
	FIREFOX = "firefox"
	SAFARI  = "safari"
)

// Chrome useragent
func Chrome() string {
	return defaultBrowser.Chrome()
}

// Random useragent
func Random() string {
	return defaultBrowser.Random()
}

// Edge useragent
func Edge() string {
	return defaultBrowser.Edge()
}

// Safari useragent
func Safari() string {
	return defaultBrowser.Safari()
}

// Firefox useragent
func Firefox() string {
	return defaultBrowser.Firefox()
}

type browser struct{}

// defaultBrowser default browser
var defaultBrowser = &browser{}

func (b *browser) Random() string {
	return useragent.UA.GetAllRandom()
}

func (b *browser) Chrome() string {
	return useragent.UA.GetRandom(CHROME)
}

func (b *browser) Edge() string {
	return useragent.UA.GetRandom(EDGE)
}

func (b *browser) Firefox() string {
	return useragent.UA.GetRandom(FIREFOX)
}

func (b *browser) Safari() string {
	return useragent.UA.GetRandom(SAFARI)
}
