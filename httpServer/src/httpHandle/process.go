package httpHandle

import (
	"httpServer/src/config"
	"httpServer/src/tipCollection"
	"httpServer/src/tipCollectionEx"
	"net/http"
	"strconv"
)

func doLogic(w http.ResponseWriter, index int, filename string, resultfile string) {

	logicType := GlobalTypeList[index]

	switch logicType {
	case config.TIP_COLLECTION:
		tipCollection.DoLogic(w, filename, resultfile)
		return
	case config.TIP_COLLECTION_EX:
		tipCollectionEx.DoLogic(w, filename, resultfile)
		return
	}

	w.Write([]byte("unknow logic type,index:" + strconv.Itoa(index) + ",file:" + filename + ",result:" + resultfile))
}
