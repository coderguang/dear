package dearDataCollectionDef

import (
	"strings"

	"github.com/coderguang/GameEngine_go/sglog"
)

type CollectionData struct {
	Title     string
	PeopleNum string
	BuyBack   string
	BuyMoney  string
	TotalUse  string
	SmsUse    string
}

func (data *CollectionData) IsEmpty() bool {
	if "" == data.Title && "" == data.PeopleNum && "" == data.BuyBack && "" == data.BuyMoney && "" == data.TotalUse && "" == data.SmsUse {
		return true
	}
	return false
}

func (data *CollectionData) IsDataError() bool {
	if "" == data.Title || "" == data.PeopleNum || "" == data.BuyBack || "" == data.BuyMoney || "" == data.TotalUse || "" == data.SmsUse {
		return true
	}
	return false
}

func (data *CollectionData) FliterChar() {
	data.Title = strings.Replace(data.Title, "\t", "", -1)
	data.PeopleNum = strings.Replace(data.PeopleNum, "\t", "", -1)
	data.BuyBack = strings.Replace(data.BuyBack, "\t", "", -1)
	data.BuyMoney = strings.Replace(data.BuyMoney, "\t", "", -1)
	data.TotalUse = strings.Replace(data.TotalUse, "\t", "", -1)
	data.SmsUse = strings.Replace(data.SmsUse, "\t", "", -1)

	data.BuyMoney = strings.Replace(data.BuyMoney, ",", "", -1)
	data.TotalUse = strings.Replace(data.TotalUse, ",", "", -1)
	data.SmsUse = strings.Replace(data.SmsUse, ",", "", -1)
}

func (data *CollectionData) ShowMsg() {
	sglog.Info("start show msg=============")
	sglog.Info("计划主题:%s", data.Title)
	sglog.Info("短信覆盖人数:%s", data.PeopleNum)
	sglog.Info("回购人数:%s", data.BuyBack)
	sglog.Info("回购总额:%s", data.BuyMoney)
	sglog.Info("总消费总额:%s", data.TotalUse)
	sglog.Info("短信花费:%s", data.SmsUse)
	sglog.Info("end show msg=============")
}

func (data *CollectionData) Reset() {
	data.Title = ""
	data.PeopleNum = ""
	data.BuyBack = ""
	data.BuyMoney = ""
	data.TotalUse = ""
	data.SmsUse = ""
}
