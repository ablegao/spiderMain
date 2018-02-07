// Able Gao @
// ablegao@gmail.com
// description：
//
//

package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_loadConfig(t *testing.T) {
	conf, err := ReadConfigFile("./config.yaml")
	if err != nil {
		t.Error(err)
	}
	log.Println(conf)

}
func Test_request(t *testing.T) {
	conf, _ := ReadConfigFile("./config.yaml")
	t.Log(conf.Task[0])

	var out *bytes.Buffer
	var err error
	for _, task := range conf.Task {
		task.SetConfigure(conf)
		out, err = task.Exec(out)
		if err != nil {
			t.Error(err)
		}
	}

	t.Log(out)

}
