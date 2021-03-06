// +build ignore
package main

// This program generates Go code via countdown.html and installs it into your program as a constant
// of the form: [*CountdownEntry, *CountdownEntry...].

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
	"time"
	"io/ioutil"
	"github.com/araddon/dateparse"
	"strings"
)

const TEMPLATE = `// Code generated by go generate; DO NOT EDIT.
package main

import "text/template"

type CountdownEntry struct {
	Text string
	Time int64
}

var PageTemplate = template.Must(template.New("").Parse({{ printf "%q" .PageContent }}))

var CountdownEntries = []CountdownEntry{
{{- range .Entries }}
	{ {{ printf "%q, %d" .Text .Time.Epoch }} },
{{- end }}
}
`

var now = time.Now()

type CountdownEntry struct {
	Text string `json:"text"`
	Time customTime `json:"time"`
}

type customTime struct {
	time.Time
	Epoch int64
}

func (self *customTime) UnmarshalJSON(buf []byte) error {
	var (
		str string = strings.Trim(strings.TrimSpace(string(buf)), `"'`)
		t time.Time
		d time.Duration
		err error
	)
	
	if strings.HasPrefix(str, "+") {
		if d, err = time.ParseDuration(str[1:]); err != nil {
			return err
		} else {
			t = time.Unix(now.Unix() + int64(d.Seconds()), 0)
		}
		// It's an offset, testing only usually
	} else if t, err = dateparse.ParseAny(str); err != nil {
		return err
	}

	self.Time = t
	self.Epoch = t.Unix()
	return nil
}

func main() {
	var (
		pagedata []byte
		tpl *template.Template
		outfile *os.File
		err error
		rawjson []byte
		jsondata []CountdownEntry
	)
	zone, _ := now.In(time.Local).Zone()
	fmt.Printf("Parsing dates in time zone %s\n",  zone)

	if tpl, err = template.New("").Parse(TEMPLATE); err != nil {
		panic(err)
	}
	
	if pagedata, err = ioutil.ReadFile("countdown.html"); err != nil {
		panic(err)
	}

	if _, err = template.New("").Parse(string(pagedata)); err != nil {
		panic(err)
	}

	if rawjson, err = ioutil.ReadFile("data.json"); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(rawjson, &jsondata); err != nil {
		panic(err)
	}

	if outfile, err = os.OpenFile("data_generated.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
		panic(err)
	}
	defer outfile.Close()

	for _, v := range jsondata {
		fmt.Printf("Message: %s\nDate: %s\n\n", v.Text, v.Time)
	}

	tpl.Execute(outfile, struct {
		Entries 	[]CountdownEntry
		PageContent   string
	}{
		jsondata,
		string(pagedata),
	})
}

