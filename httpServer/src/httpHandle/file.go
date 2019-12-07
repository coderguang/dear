package httpHandle

import "httpServer/src/config"

import "github.com/coderguang/GameEngine_go/sgtime"

import "strconv"

func getPathAndFileName(index int) (string, string) {
	path := config.GlobalCfg.UploadPath + "/" + GlobalTypeList[index]
	filename := path + "/" + getNumString(index) + GLobalFileSuffix[index]

	return path, filename
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
