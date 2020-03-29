package tipCollectionEx

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/coderguang/GameEngine_go/sglog"
)

func StartParse(filename string, resultfile string) error {

	sglog.Info("start parse tipCollection ex", filename)

	xls, err := excelize.OpenFile(filename)
	if err != nil {
		sglog.Fatal("read file:", filename, "error,err:s", err)
		return errors.New("read file:" + filename + "error,err:s" + err.Error())
	}

	sheetName := "root"
	rows, err := xls.Rows(sheetName)
	if err != nil {
		sglog.Error("读取 ", sheetName, " 工作表 错误,err=", err)
		return errors.New("读取 " + sheetName + " 工作表 错误,err=" + err.Error())
	}

	totalline := 1
	breakLine := 0
	breakUserName := ""
	for rows.Next() {
		posStr := strconv.Itoa(totalline)
		userName := xls.GetCellValue(sheetName, "A"+posStr)
		if userName == "" {
			//recheck next two column
			userName = xls.GetCellValue(sheetName, "A"+strconv.Itoa(totalline+1))
			if userName == "" {
				userName = xls.GetCellValue(sheetName, "A"+strconv.Itoa(totalline+2))
				if userName == "" {
					break
				}
			}

		}
		totalline++
		breakLine = totalline
		breakUserName = userName
	}

	allTags := [13][]string{}
	phoneTipMaps := make(map[string][]string)
	realLine := 0
	for i := 2; i <= totalline; i++ {
		posStr := strconv.Itoa(i)
		userName := xls.GetCellValue(sheetName, "A"+posStr)
		if userName == "" {
			continue
		}
		phone := xls.GetCellValue(sheetName, "K"+posStr)
		Tag0 := xls.GetCellValue(sheetName, "L"+posStr)
		Tag1 := xls.GetCellValue(sheetName, "M"+posStr)
		Tag2 := xls.GetCellValue(sheetName, "N"+posStr)
		Tag3 := xls.GetCellValue(sheetName, "O"+posStr)
		Tag4 := xls.GetCellValue(sheetName, "P"+posStr)
		Tag5 := xls.GetCellValue(sheetName, "Q"+posStr)
		Tag6 := xls.GetCellValue(sheetName, "R"+posStr)
		Tag7 := xls.GetCellValue(sheetName, "S"+posStr)
		Tag8 := xls.GetCellValue(sheetName, "T"+posStr)
		Tag9 := xls.GetCellValue(sheetName, "U"+posStr)
		Tag10 := xls.GetCellValue(sheetName, "V"+posStr)
		Tag11 := xls.GetCellValue(sheetName, "W"+posStr)
		Tag12 := xls.GetCellValue(sheetName, "X"+posStr)

		allTags[0] = append(allTags[0], Tag0)
		allTags[1] = append(allTags[1], Tag1)
		allTags[2] = append(allTags[2], Tag2)
		allTags[3] = append(allTags[3], Tag3)
		allTags[4] = append(allTags[4], Tag4)
		allTags[5] = append(allTags[5], Tag5)
		allTags[6] = append(allTags[6], Tag6)
		allTags[7] = append(allTags[7], Tag7)
		allTags[8] = append(allTags[8], Tag8)
		allTags[9] = append(allTags[9], Tag9)
		allTags[10] = append(allTags[10], Tag10)
		allTags[11] = append(allTags[11], Tag11)
		allTags[12] = append(allTags[12], Tag12)

		if phone != "" {
			tipList, ok := phoneTipMaps[phone]
			if ok {
				tipList = append(tipList, Tag0, Tag1, Tag2, Tag3, Tag4, Tag5, Tag6, Tag7, Tag8, Tag9, Tag10, Tag11, Tag12)
				phoneTipMaps[phone] = tipList
			} else {
				phoneTipMaps[phone] = []string{Tag0, Tag1, Tag2, Tag3, Tag4, Tag5, Tag6, Tag7, Tag8, Tag9, Tag10, Tag11, Tag12}
			}
			//sglog.Debug("tiplist", tipList)
		}
		realLine++
	}

	sglog.Info("total line:", totalline, "total tips:", len(allTags), ",realLine:", realLine)

	mapAllTags := make(map[int]map[string]int)

	for i := 0; i < 13; i++ {
		for _, v := range allTags[i] {
			if v == "" {
				continue
			}
			tiplist := strings.Split(v, ",")
			for _, tip := range tiplist {

				indexMap, ok := mapAllTags[i]
				if ok {
					tipMap, ok := indexMap[tip]
					if ok {
						indexMap[tip] = tipMap + 1
					} else {
						indexMap[tip] = 1
					}
				} else {
					mapAllTags[i] = make(map[string]int)
					mapAllTags[i][tip] = 1
				}

			}
		}

	}

	return WriteXlsx(resultfile, breakUserName, breakLine, mapAllTags, phoneTipMaps)
}

func WriteXlsx(resultfile string, breakName string, breakLine int, mapData map[int]map[string]int, phoneTipMaps map[string][]string) error {

	sglog.Info("start write to file,breakline:", breakLine, ",breakName:", breakName)

	file := excelize.NewFile()
	sheetName := "result"
	index := file.NewSheet(sheetName)
	file.SetActiveSheet(index)

	orderStr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}
	descStr := []string{"标签组1(客户登记)", "标签组2(兴趣爱好)", "标签组3(职业)", "标签组4(客户类型)",
		"标签组5(站外来源渠道)", "标签组6(站内来源渠道)", "标签组7(人生阶段)", "标签组8(年龄段)",
		"标签组9(性别)", "标签组10(肌肤类型)", "标签组11(肌肤问题)", "标签组12(购买商品类目)", "标签组13(购物侧重点)"}

	for index, v := range orderStr {
		file.SetCellStr(sheetName, v+"1", descStr[index])
	}

	//write detail tags
	totalTagType := 0
	totalTagNum := 0
	for k, v := range mapData {
		curLine := 2
		for kk, vv := range v {
			totalTagType++
			for i := 0; i < vv; i++ {
				curLineStr := strconv.Itoa(curLine)
				file.SetCellStr(sheetName, orderStr[k]+curLineStr, kk)
				curLine++
				totalTagNum++
			}
		}
	}

	//write phone list

	sheetNamePhone := "phone"
	index = file.NewSheet(sheetNamePhone)
	file.SetActiveSheet(index)

	file.SetCellStr(sheetNamePhone, "A1", "电话号码")
	file.SetCellStr(sheetNamePhone, "B1", "标签")
	file.SetCellStr(sheetNamePhone, "C1", "所属标签组")

	phoneList := []string{}
	for k, _ := range phoneTipMaps {
		phoneList = append(phoneList, k)
	}

	sort.Strings(phoneList)

	writePhoneIndex := 2
	for _, v := range phoneList {
		tipList, ok := phoneTipMaps[v]
		if ok {
			for k, tipV := range tipList {
				if tipV == "" {
					continue
				}
				singleTipList := strings.Split(tipV, ",")
				for _, singleV := range singleTipList {
					curLineStr := strconv.Itoa(writePhoneIndex)
					file.SetCellStr(sheetNamePhone, "A"+curLineStr, v)
					file.SetCellStr(sheetNamePhone, "B"+curLineStr, singleV)
					file.SetCellStr(sheetNamePhone, "C"+curLineStr, descStr[k%13])
					writePhoneIndex++
				}
			}
		}
	}

	sheetNameNum := "num"
	indexNum := file.NewSheet(sheetNameNum)
	file.SetActiveSheet(indexNum)

	//write tagNum tags
	file.SetCellStr(sheetNameNum, "A1", "标签总类别数量")
	file.SetCellStr(sheetNameNum, "A2", strconv.Itoa(totalTagType))

	file.SetCellStr(sheetNameNum, "B1", "标签总数量")
	file.SetCellStr(sheetNameNum, "B2", strconv.Itoa(totalTagNum))

	file.SetCellStr(sheetNameNum, "C1", "终止检测行数")
	file.SetCellStr(sheetNameNum, "C2", strconv.Itoa(breakLine))

	file.SetCellStr(sheetNameNum, "D1", "最后一个检测的客户端名称")
	file.SetCellStr(sheetNameNum, "D2", breakName)

	err := file.SaveAs(resultfile)

	if err != nil {
		sglog.Error("save file error,err:", err)
		return errors.New("save file error,err:," + err.Error())
	}
	sglog.Info("write all data complete")
	return nil
}
