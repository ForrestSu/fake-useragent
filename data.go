package fake_useragent

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/ForrestSu/fake-useragent/useragent"
)

//go:embed data/browsers.json
var cacheData []byte

func init() {
	var m map[string][]string
	if err := json.Unmarshal(cacheData, &m); err != nil {
		fmt.Println(err)
	}
	useragent.UA.SetData(m)
}
