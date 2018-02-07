// Able Gao @
// ablegao@gmail.com
// description：
//
//

package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// http.Get
func (task ConfigureTask) RunnerHttpGet(buf *bytes.Buffer) error {
	if task.Value == "" {
		task.Value = buf.String()
	}
	req, err := http.NewRequest("GET", task.Value, nil)

	if err != nil {
		return err
	}
	for x, v := range task.conf.Header {
		req.Header.Set(x, v)
	}
	for x, v := range task.Header {
		req.Header.Set(x, v)
	}

	if req, err := http.DefaultClient.Do(req); err == nil {
		io.Copy(buf, req.Body)
		req.Body.Close()
		return nil
	} else {
		return err
	}
}

// HTML
func (task ConfigureTask) RunnerHTML(buf *bytes.Buffer) error {

	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return err
	}

	buf.Reset()
	var sel *goquery.Selection

	for i, x := range task.Find {
		if i == 0 {
			sel = doc.Find(x)

		} else {
			sel = sel.Find(x)
		}
	}
	if sel == nil {
		return errors.New("Not Found")
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
			default:
				h = e.Text() + "\n"
			}
		}
		buf.WriteString(h)

	})

	return nil
}

func (task ConfigureTask) RunnerWriteToFile(buf *bytes.Buffer) error {
	if fi, err := os.OpenFile(task.OutputPath, os.O_WRONLY|os.O_CREATE, 0644); err == nil {
		io.Copy(fi, buf)
		buf.Reset()
		fi.Close()
		log.Println("Write to file :", task.OutputPath)
		return nil
	} else {
		return err
	}
}

func (task ConfigureTask) RunnerEachDownload(buf *bytes.Buffer) error {

	scanner := bufio.NewReader(buf)
	downloads := task.OutputPath
	for {
		line, e := scanner.ReadString('\n')
		line = strings.Trim(line, "\n")
		if e != nil {
			break
		}
		u, err := url.Parse(line)
		if err != nil {
			continue
		}
		baseName := path.Base(u.Path)
		if len(baseName) < 2 {
			continue
		}

		// make file path
		d := downloads + path.Dir(u.Path)
		err = os.MkdirAll(d, 0775)
		if err != nil {
			log.Println("ERR", err)
			continue
		}
		// get file body
		out := bytes.NewBuffer(nil)
		task.Value = line
		err = task.RunnerHttpGet(out)
		if err != nil {
			log.Println("ERR", err)
			continue
		}
		fi := d + "/" + baseName
		// write to file
		task.OutputPath = fi
		err = task.RunnerWriteToFile(out)
		if err != nil {
			log.Println("ERR", err)
			continue
		}

	}
	return nil

}

func (task ConfigureTask) RunnerEachHTML(buf *bytes.Buffer) error {
	in := bytes.NewBuffer(nil)
	io.Copy(in, buf)
	buf.Reset()
	scanner := bufio.NewReader(in)

	linkOut := bytes.NewBuffer(nil)
	for {
		line, e := scanner.ReadString('\n')
		line = strings.Trim(line, "\n")
		if e != nil {
			break
		}

		task.Value = line
		task.RunnerHttpGet(linkOut)
		task.RunnerHTML(linkOut)
		buf.Write(linkOut.Bytes())
		linkOut.Reset()
	}
	return nil
}
