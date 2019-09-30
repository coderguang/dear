package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sgthread"
	"github.com/mohae/deepcopy"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

type Data struct {
	NewTips        string
	Head           string
	NickName       string
	WxBeiZhu       string
	TaobaoNickName string
	YouZanNickName string
	JDNickName     string
	WxOpenId       string
	IsGZHFans      string
	Tips           string
	Sex            string
	Lock           string
	BirthDay       string
	LastMsgSendDt  string
	LastMsgRecvDt  string
	FriendDt       string
	JiaoYiDt       string
}

type TipData struct {
	Tips string
	Num  int
}

type TipDataList []TipData

func (data TipDataList) Len() int {
	return len(data)
}
func (data TipDataList) Less(i, j int) bool {
	return data[i].Num > data[j].Num
}
func (data TipDataList) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func StartParse() {
	filename := "./data/data.xlsx"
	sglog.Info("start parse ", filename)

	xls, err := excelize.OpenFile(filename)
	if err != nil {
		sglog.Fatal("read file:", filename, "error,err:s", err)
		return
	}

	sheetName := "root"
	rows, err := xls.Rows(sheetName)
	if err != nil {
		sglog.Error("读取 ", sheetName, " 工作表 错误,err=", err)
		sgthread.DelayExit(2)
	}

	totalline := 0
	for rows.Next() {
		totalline++
	}

	alldata := make(map[string][]*Data)

	for i := 2; i <= totalline; i++ {
		tmp := new(Data)
		posStr := strconv.Itoa(i)
		tmp.Head = xls.GetCellValue(sheetName, "A"+posStr)
		tmp.NickName = xls.GetCellValue(sheetName, "B"+posStr)
		tmp.WxBeiZhu = xls.GetCellValue(sheetName, "C"+posStr)
		tmp.TaobaoNickName = xls.GetCellValue(sheetName, "D"+posStr)
		tmp.YouZanNickName = xls.GetCellValue(sheetName, "E"+posStr)
		tmp.JDNickName = xls.GetCellValue(sheetName, "F"+posStr)
		tmp.WxOpenId = xls.GetCellValue(sheetName, "G"+posStr)
		tmp.IsGZHFans = xls.GetCellValue(sheetName, "H"+posStr)
		tmp.Sex = xls.GetCellValue(sheetName, "J"+posStr)
		tmp.Lock = xls.GetCellValue(sheetName, "K"+posStr)
		tmp.BirthDay = xls.GetCellValue(sheetName, "L"+posStr)
		tmp.LastMsgSendDt = xls.GetCellValue(sheetName, "M"+posStr)
		tmp.LastMsgRecvDt = xls.GetCellValue(sheetName, "N"+posStr)
		tmp.FriendDt = xls.GetCellValue(sheetName, "O"+posStr)
		tmp.JiaoYiDt = xls.GetCellValue(sheetName, "P"+posStr)

		oldTips := xls.GetCellValue(sheetName, "I"+posStr)

		if oldTips == "" {
			tmp.Tips = oldTips

			if _, ok := alldata[tmp.Tips]; ok {
				alldata[tmp.Tips] = append(alldata[tmp.Tips], tmp)
			} else {
				tmplist := []*Data{tmp}
				alldata[tmp.Tips] = tmplist
			}
		} else {
			tipslist := strings.Split(oldTips, ",")
			for _, v := range tipslist {
				ttdata := deepcopy.Copy(tmp)
				ttdataV, ok := ttdata.(*Data)
				if !ok {
					sglog.Error("parse data error")
					sgthread.DelayExit(2)
					continue
				}
				ttdataV.Tips = v
				if _, ok := alldata[ttdataV.Tips]; ok {
					alldata[ttdataV.Tips] = append(alldata[ttdataV.Tips], ttdataV)
				} else {
					tmplist := []*Data{ttdataV}
					alldata[ttdataV.Tips] = tmplist
				}
			}
		}

	}

	sglog.Info("total line:", totalline, "total tips:", len(alldata))
	WriteXlsx(alldata)
}

func WriteXlsx(alldatas map[string][]*Data) {
	totalline := 0
	tipsArray := TipDataList{}
	for k, v := range alldatas {
		totalline += len(v)
		tmp := TipData{
			Tips: k, Num: len(v)}
		tipsArray = append(tipsArray, tmp)
	}
	sort.Sort(tipsArray)

	for _, v := range tipsArray {
		sglog.Info(v.Tips, ":", v.Num)
	}

	sglog.Info("start write to file ,total write is ", totalline)

	file := excelize.NewFile()
	sheetName := "result"
	index := file.NewSheet(sheetName)
	file.SetActiveSheet(index)

	orderStr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
	descStr := []string{"好友标签", "头像", "微信昵称", "微信备注", "淘宝昵称", "有赞昵称", "京东昵称", "微信号",
		"是否公众号粉丝", "性别", "地区", "生日", "最近消息发送时间", "最近消息接收时间", "添加好友时间", "最近交易时间"}

	for index, v := range orderStr {
		file.SetCellStr(sheetName, v+"1", descStr[index])
	}

	writeIndex := 2
	for _, tips := range tipsArray {
		if v, ok := alldatas[tips.Tips]; ok {
			for _, vv := range v {
				writeIndexStr := strconv.Itoa(writeIndex)
				file.SetCellStr(sheetName, orderStr[0]+writeIndexStr, vv.Tips)
				file.SetCellStr(sheetName, orderStr[1]+writeIndexStr, vv.Head)
				file.SetCellStr(sheetName, orderStr[2]+writeIndexStr, vv.NickName)
				file.SetCellStr(sheetName, orderStr[3]+writeIndexStr, vv.WxBeiZhu)
				file.SetCellStr(sheetName, orderStr[4]+writeIndexStr, vv.TaobaoNickName)
				file.SetCellStr(sheetName, orderStr[5]+writeIndexStr, vv.YouZanNickName)
				file.SetCellStr(sheetName, orderStr[6]+writeIndexStr, vv.JDNickName)
				file.SetCellStr(sheetName, orderStr[7]+writeIndexStr, vv.WxOpenId)
				file.SetCellStr(sheetName, orderStr[8]+writeIndexStr, vv.IsGZHFans)
				file.SetCellStr(sheetName, orderStr[9]+writeIndexStr, vv.Sex)
				file.SetCellStr(sheetName, orderStr[10]+writeIndexStr, vv.Lock)
				file.SetCellStr(sheetName, orderStr[11]+writeIndexStr, vv.BirthDay)
				file.SetCellStr(sheetName, orderStr[12]+writeIndexStr, vv.LastMsgSendDt)
				file.SetCellStr(sheetName, orderStr[13]+writeIndexStr, vv.LastMsgRecvDt)
				file.SetCellStr(sheetName, orderStr[14]+writeIndexStr, vv.FriendDt)
				file.SetCellStr(sheetName, orderStr[15]+writeIndexStr, vv.JiaoYiDt)
				writeIndex++
			}
		} else {
			sglog.Error("can't find tips,", tips.Tips)
			sgthread.DelayExit(2)
		}
	}

	tipSheetName := "tips"
	tipIndex := file.NewSheet(tipSheetName)
	file.SetActiveSheet(tipIndex)

	writeIndex = 1
	for _, v := range tipsArray {
		writeIndexStr := strconv.Itoa(writeIndex)
		file.SetCellStr(tipSheetName, orderStr[0]+writeIndexStr, v.Tips)
		file.SetCellStr(tipSheetName, orderStr[1]+writeIndexStr, strconv.Itoa(v.Num))
		writeIndex++
	}

	err := file.SaveAs("./data/result.xlsx")

	if err != nil {
		sglog.Error("save file error,err:", err)
		sgthread.DelayExit(2)
	}
	sglog.Info("write all data complete")
}

func main() {

	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)

	sglog.Info("welcome to tips collection !any question can ask royalchen@royalchen.com")

	StartParse()

	sgcmd.StartCmdWaitInputLoop()

}
