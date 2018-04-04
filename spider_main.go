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

	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

var (
	config = flag.String("config", "./config.yaml", "Task Config")
	tpl    = flag.String("tpl", "./simple.html", "GUI Theme")
	nw     = flag.Bool("nw", false, "Run With GUI")
)
var (
	client = http.Client{}
)

func runTask() {
	conf, _ := ReadConfigFile(*config)

	var err error
	out := bytes.NewBuffer(nil)
	for _, task := range conf.Task {
		task.SetConfigure(conf)
		err = task.Exec(out)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	flag.Parse()
	if *nw == false {

		rect := sciter.NewRect(100, 100, 640, 640)
		w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, rect)
		if err != nil {
			log.Fatal(err)
		}
		// log.Printf("handle: %v", w.Handle)
		w.LoadFile(*tpl)
		//w.DefineFunction("debug",func(args ...*sciter.Value) *sciter.Value){})
		w.DefineFunction("debugLog", func(args ...*sciter.Value) *sciter.Value {
			log.Println(args)
			return nil
		})

		w.Show()

		w.Run()
	} else {
		runTask()
	}

}
