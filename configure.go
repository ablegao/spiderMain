// Able Gao @
// ablegao@gmail.com
// description：
//
//

package main

import (
	"bytes"
	"io/ioutil"

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
	RunType    string            `yaml:"run"`
	Value      string            `yaml:"value"`
	Header     map[string]string `yaml:"header"`
	Find       []string          `yaml:"find"`
	Attr       string            `yaml:"attr"`
	Output     string            `yaml:"out-type"`
	OutputPath string            `yaml:"out-path"`
	conf       Configure
}

func (task ConfigureTask) SetConfigure(conf Configure) {
	task.conf = conf
}
func (task ConfigureTask) Exec(buf *bytes.Buffer) error {
	switch task.RunType {
	case "get":
		return task.RunnerHttpGet(buf)
	case "html":
		return task.RunnerHTML(buf)
	case "write-to-file":
		return task.RunnerWriteToFile(buf)
	case "each-download":
		return task.RunnerEachDownload(buf)
	case "each-html":
		return task.RunnerEachHTML(buf)
	case "stdout":
		return task.RunnerStdout(buf)
	}

	return nil

}

type Configure struct {
	Header map[string]string `yaml:"header"`
	Task   []ConfigureTask   `yaml:"workflows"`
}
