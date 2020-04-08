package wordToVoice

import (
	"io/ioutil"
	"net/http"

	"github.com/coderguang/GameEngine_go/sglog"
)

func DoLogic(w http.ResponseWriter, filename string, resultfile string) {

	err := StartParse(filename, resultfile)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		sglog.Info("SUCCESS to deal world to voice,filename is ", resultfile)
		fileData, err := ioutil.ReadFile(resultfile)
		if err != nil {
			sglog.Error("Read mp3 data File Err:", err)
		} else {
			sglog.Info("Send File:", resultfile)
			w.Header().Set("Content-Disposition", "attachment; filename=data.mp3")
			w.Write(fileData)
		}
		sglog.Info("deal world to voice ok")
	}
}
