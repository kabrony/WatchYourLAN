package main

import (
	"flag"
	// "net/http"

	// _ "net/http/pprof"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/routines"
	"github.com/aceberg/WatchYourLAN/internal/web"
)

const dirPath = "/data/WatchYourLAN"
const nodePath = ""

func main() {
	dirPtr := flag.String("d", dirPath, "Path to config dir")
	nodePtr := flag.String("n", nodePath, "Path to node modules")
	flag.Parse()

	// pprof - memory leak detect
	// go tool pprof -alloc_space http://localhost:8085/debug/pprof/heap
	// (pprof) web
	// (pprof) list db.Select
	//
	// go func() {
	// 	http.ListenAndServe("localhost:8085", nil)
	// }()

	// Generate AppConfig
	conf.Generate(*dirPtr, *nodePtr)

	gdb.Start()

	routines.ScanRestart()
	routines.HistoryTrim()

	web.Gui()
}
