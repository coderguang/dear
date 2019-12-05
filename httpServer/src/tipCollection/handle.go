package tipCollection

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
		sglog.Info("SUCCESS to deal xlsx file for tip collections,filename is ", resultfile)
		fileData, err := ioutil.ReadFile(resultfile)
		if err != nil {
			sglog.Error("Read tipcollections data File Err:", err)
		} else {
			sglog.Info("Send File:", resultfile)
			w.Write(fileData)
		}
		sglog.Info("deal tip collection ok")
	}
}
