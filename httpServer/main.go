package main

import (
	"httpServer/src/config"
	"httpServer/src/httpHandle"
	"log"
	"net/http"
	"strconv"

	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func main() {

	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)

	sglog.Info("start dear http server...")

	config.InitCfg("./../../globalConfig/dear/httpServer/")

	http.HandleFunc(config.GlobalCfg.Upload, httpHandle.UploadFileHandler())

	// fs := http.FileServer(http.Dir(config.GlobalCfg.UploadPath))
	// http.Handle(config.GlobalCfg.DownloadPath, http.StripPrefix(config.GlobalCfg.DownloadPath, fs))

	serverUrl := "0.0.0.0:" + strconv.Itoa(config.GlobalCfg.Port)
	sglog.Info("Server started on ", serverUrl, ", use ", config.GlobalCfg.Upload, " for uploading files and ", config.GlobalCfg.Download, "/{fileName} for downloading")
	go http.ListenAndServe(serverUrl, nil)

	sgcmd.StartCmdWaitInputLoop()

}
