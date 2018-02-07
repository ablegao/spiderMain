// Able Gao @
// ablegao@gmail.com
// description：
//
//

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	yaml "gopkg.in/yaml.v2"
)

func ReadConfigFile(p string) (conf Configure, err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(p)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &conf)
	return
}

type ConfigureTask struct {
	InputType string            `yaml:"input-type"`
	Value     string            `yaml:"value"`
	Header    map[string]string `yaml:"header"`
	Find      []string          `yaml:"find"`
	Attr      string            `yaml:"attr"`
	Output    string            `yaml:"output"`
	conf      Configure
}

func (task ConfigureTask) SetConfigure(conf Configure) {
	task.conf = conf
}
func (task ConfigureTask) Exec(in *bytes.Buffer) (*bytes.Buffer, error) {
	var doc *goquery.Document
	var buf = bytes.NewBuffer(nil)
	switch task.InputType {
	case "GET":
		req, err := http.NewRequest("GET", task.Value, nil)
		if err != nil {
			return buf, err
		}
		for x, v := range task.conf.Header {
			req.Header.Set(x, v)
		}
		for x, v := range task.Header {
			req.Header.Set(x, v)
		}

		if req, err := http.DefaultClient.Do(req); err == nil {
			doc, err = goquery.NewDocumentFromResponse(req)

		} else {
			return buf, err
		}
	case "html":
		buf1 := bytes.NewBuffer(in.Bytes())
		var err error
		doc, err = goquery.NewDocumentFromReader(buf1)
		if err != nil {
			return buf, err
		}

	}
	var sel *goquery.Selection

	for i, x := range task.Find {
		if i == 0 {
			sel = doc.Find(x)

		} else {
			sel = sel.Find(x)
		}
	}

	sel.Each(func(i int, e *goquery.Selection) {
		var h string

		if task.Attr != "" {
			//h,ok = e.Attr(task.Attr)
			if s, ok := e.Attr(task.Attr); ok {
				h = s + "\n"
			}
		} else {
			switch task.Output {
			case "html":
				h, _ = goquery.OuterHtml(e)
			case "text":
				h = e.Text() + "\n"
			}
		}
		buf.WriteString(h)

	})
	return buf, nil

}

type Configure struct {
	Header map[string]string `yaml:"header"`
	Task   []ConfigureTask   `yaml:"task"`
}
