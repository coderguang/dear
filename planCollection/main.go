package main

import (
	"log"
	"os"
	dearDataCollectionScanner "planCollection/src/scanner"

	"github.com/coderguang/GameEngine_go/sgcmd"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func main() {

	sgserver.StartLogServer("debug", "./log", log.LstdFlags, true)

	arg_num := len(os.Args) - 1
	if arg_num < 1 {
		sglog.Error("please input data source file ")
		return
	}

	sglog.Info("welcome to data collection !any question can ask royalchen@royalchen.com")

	fileList := []string{}
	for index, v := range os.Args {
		if index == 0 {
			continue
		}

		fileList = append(fileList, v)
	}

	sglog.Info("file list %s", fileList)

	dearDataCollectionScanner.StartParseFileList(fileList)

	sgcmd.StartCmdWaitInputLoop()

}
