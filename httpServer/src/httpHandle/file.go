package httpHandle

import (
	"httpServer/src/config"
	"strconv"

	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgstring"
	"github.com/coderguang/GameEngine_go/sgtime"
)

func getPathAndFileName(index int) (string, string, string) {
	path := config.GlobalCfg.UploadPath + "/" + GlobalTypeList[index]
	rangdomStr := sgstring.RandNumStringRunes(5)
	filename := path + "/" + getNumString(index) + rangdomStr + "." + GlobalFileSuffix[index]
	resultfile := path + "/" + getNumString(index) + rangdomStr + "_result." + GlobalFileSuffix[index]
	return path, filename, resultfile
}

func getNumString(index int) string {
	now := sgtime.New()
	monthDay := sgtime.YMDString(now)
	cur := GlobalFileNum[index]
	GlobalFileNum[index] = cur + 1
	curStr := ""
	if cur < 10 {
		curStr = "00" + strconv.Itoa(cur)
	} else if cur < 100 {
		curStr = "0" + strconv.Itoa(cur)
	} else {
		curStr = strconv.Itoa(cur)
	}
	return monthDay + curStr
}

func writefile(filename string, fileBytes []byte) error {
	newFile, err := sgfile.Create(filename)
	if err != nil {
		sglog.Error("crete file error,", err)
		return err
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		sglog.Error("write to file error")
		return err
	}
	return nil
}
