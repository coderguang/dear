package main

import (
	"httpServer/src/config"
	"httpServer/src/httpHandle"
	"log"
	"strconv"

	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func main() {

	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)

	sglog.Info("start dear http server...")

	config.InitCfg("./../../globalConfig/dear/httpServer/config.json")

	listenPort := strconv.Itoa(config.GlobalCfg.Port)
	go httpHandle.NewWebServer(listenPort)

	sgcmd.StartCmdWaitInputLoop()

}
