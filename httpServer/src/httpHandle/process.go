package httpHandle

import (
	"errors"
	"httpServer/src/config"
	"httpServer/src/tipCollection"
	"net/http"
	"strconv"
)

func doLogic(w http.ResponseWriter, index int, filename string, resultfile string) error {

	logicType := GlobalTypeList[index]

	switch logicType {
	case config.TIP_COLLECTION:
		return tipCollection.DoLogic(w, filename, resultfile)
	}

	return errors.New("unknow logic type,index:" + strconv.Itoa(index) + ",file:" + filename + ",result:" + resultfile)
}
