package data

import (
	_ "embed"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

type userAgentInfo struct {
	Percent   float64 `json:"percent"`
	Useragent string  `json:"useragent"`
	System    string  `json:"system"`
	Browser   string  `json:"browser"`
	Version   float64 `json:"version"`
	Os        string  `json:"os"`
}

//go:embed browsers_20231117.json
var contents string

func Test_GenBrowsersFile(t *testing.T) {
	var m = make(map[string][]string)
	rows := strings.Split(contents, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		info := &userAgentInfo{}
		if err := json.Unmarshal([]byte(row), info); err != nil {
			t.Fatal(err)
		}
		if info.Browser == "" {
			t.Fatal("browser is empty")
		}
		m[info.Browser] = append(m[info.Browser], info.Useragent)
	}
	t.Log("total load", len(rows))
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		t.Fatalf("Marshal err=%v", err)
	}
	if err = os.WriteFile("./browsers.json", data, 0666); err != nil {
		t.Fatalf("WriteFile err=%v", err)
	}
}
