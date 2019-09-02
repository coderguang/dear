package dearDataCollectionScanner

import (
	"bufio"
	"io"
	"os"
	dearDataCollectionDef "planCollection/src/def"
	dearDataCollectionXlsx "planCollection/src/xlsx"
	"strings"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgthread"
)

func StartParseFileList(fileList []string) {
	for _, v := range fileList {
		StartParseData(v)
		dearDataCollectionXlsx.WriteDataToXlsx(v)
		dearDataCollectionXlsx.ResetData()
	}
}

func StartParseData(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		sglog.Fatal("read file:%s error,err:=%s", filename, err)
		return
	}

	strlist := []string{}
	sum := 0
	rd := bufio.NewReader(file)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		str := string([]byte(line))
		strlist = append(strlist, str)
		sum++
	}
	transFromFileData(strlist)
}

func transFromFileData(strlist []string) {

	tmpData := new(dearDataCollectionDef.CollectionData)

	for _, v := range strlist {
		if v == "" {
			DealData(tmpData)
			tmpData = new(dearDataCollectionDef.CollectionData)
			continue
		}
		if strings.Contains(v, "计划主题") {
			contents := strings.Split(v, "：")
			if 2 == len(contents) {
				tmpData.Title = contents[1]
			} else {
				sglog.Fatal("计划主题 data formation error")
				sgthread.DelayExit(2)
			}
		} else if strings.Contains(v, "回购ROI") {
			contents := strings.Split(v, "\t")
			if 8 == len(contents) {
				tmpstr := contents[3]
				tmpData.BuyBack = strings.Replace(tmpstr, "个", "", -1)
				tmpMoney := contents[5]
				tmpData.BuyMoney = strings.Replace(tmpMoney, "￥", "", -1)
			} else {
				sglog.Fatal("回购ROI data formation error")
				sgthread.DelayExit(2)
			}
		} else if strings.Contains(v, "短信数量") {
			contents := strings.Split(v, "\t")
			if 8 == len(contents) {
				tmpstr := contents[3]
				tmpData.PeopleNum = strings.Replace(tmpstr, "个", "", -1)
			} else {
				sglog.Fatal("短信覆盖人数 data formation error")
				sgthread.DelayExit(2)
			}
		} else if strings.Contains(v, "总消费人群") {
			contents := strings.Split(v, "\t")
			if 8 == len(contents) {
				tmpstr := contents[3]
				tmpData.TotalUse = strings.Replace(tmpstr, "￥", "", -1)
				tmp := contents[5]
				tmpData.SmsUse = strings.Replace(tmp, "￥", "", -1)
			} else {
				sglog.Fatal("总消费人群 data formation error")
				sgthread.DelayExit(2)
			}
		}

	}

	DealData(tmpData)
}

func DealData(tmpData *dearDataCollectionDef.CollectionData) {
	if !tmpData.IsEmpty() {
		tmpData.FliterChar()
		tmpData.ShowMsg()
		dearDataCollectionXlsx.AddData(tmpData)
		if tmpData.IsDataError() {
			sglog.Fatal("data error,please check")
			sgthread.DelayExit(2)
		}
	}
}
