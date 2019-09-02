package dearDataCollectionXlsx

import (
	dearDataCollectionDef "planCollection/src/def"
	"strconv"
	"strings"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var datalist []*dearDataCollectionDef.CollectionData

func init() {
	datalist = []*dearDataCollectionDef.CollectionData{}
}

func AddData(d *dearDataCollectionDef.CollectionData) {
	datalist = append(datalist, d)
}

func ResetData() {
	datalist = nil
	datalist = []*dearDataCollectionDef.CollectionData{}
}

func WriteDataToXlsx(filename string) {
	xlsxFile := strings.Replace(filename, "txt", "xlsx", -1)

	sheetName := "data"
	file := excelize.NewFile()
	index := file.NewSheet(sheetName)

	file.SetActiveSheet(index)

	columns := []string{"A", "D", "G", "J", "M", "P", "S", "V", "X", "Z"}
	rows := []string{}

	totalSize := len(datalist)

	for i := 3; i < totalSize*10; i = i + 2 {
		numstr := strconv.Itoa(i)
		rows = append(rows, numstr)
	}

	titlelist := []string{"计划主题", "短信覆盖人数", "回购人数", "回购总额", "总消费总额", "短信花费"}

	for k := range titlelist {
		file.SetColWidth(sheetName, columns[k], columns[k], 30)
		file.SetCellStr(sheetName, columns[k]+"1", titlelist[k])
	}

	// xlsxStyle, err := file.NewStyle(`{
	// 	"alignment":{
	// 		"horizontal":"center",
	// 		"vertical":"center",
	// 		"wrap_text":true
	// 	}
	// }`)
	// if err != nil {
	// 	sglog.Error("create xlsx style error,%s", err)
	// 	sgthread.DelayExit(2)
	// }

	rowIndex := 0
	for _, v := range datalist {
		rowPos := rows[rowIndex]
		rowIndex++

		// file.SetCellStyle(sheetName, columns[0]+rowPos, columns[0]+rowPos, xlsxStyle)
		// file.SetCellStyle(sheetName, columns[1]+rowPos, columns[1]+rowPos, xlsxStyle)
		// file.SetCellStyle(sheetName, columns[2]+rowPos, columns[2]+rowPos, xlsxStyle)
		// file.SetCellStyle(sheetName, columns[3]+rowPos, columns[3]+rowPos, xlsxStyle)
		// file.SetCellStyle(sheetName, columns[4]+rowPos, columns[4]+rowPos, xlsxStyle)
		// file.SetCellStyle(sheetName, columns[5]+rowPos, columns[5]+rowPos, xlsxStyle)

		file.SetCellStr(sheetName, columns[0]+rowPos, v.Title)
		file.SetCellStr(sheetName, columns[1]+rowPos, v.PeopleNum)
		file.SetCellStr(sheetName, columns[2]+rowPos, v.BuyBack)
		file.SetCellStr(sheetName, columns[3]+rowPos, v.BuyMoney)
		file.SetCellStr(sheetName, columns[4]+rowPos, v.TotalUse)
		file.SetCellStr(sheetName, columns[5]+rowPos, v.SmsUse)

	}

	err := file.SaveAs(xlsxFile)

	if err != nil {
		sglog.Error("save file error,file:%s,err:%s", xlsxFile, err)
		sgthread.DelayExit(2)
	}

}
