// Able Gao @
// ablegao@gmail.com
// description：
//
//

package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
)

var config = flag.String("config", "./config.yaml", "Task Config")
var client = http.Client{}

func main() {
	flag.Parse()
	conf, _ := ReadConfigFile(*config)
	var out *bytes.Buffer
	var err error
	for _, task := range conf.Task {
		task.SetConfigure(conf)
		out, err = task.Exec(out)
		if err != nil {

			log.Println("ERROR", err)
			return
		}
	}
	log.Println(out.String())

}
