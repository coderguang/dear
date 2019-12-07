package httpHandle

import (
	"crypto/rand"
	"httpServer/src/config"
	"fmt"
	"net/http"
)

var GlobalTypeList []string
var GlobalFileType []string
var GlobalFileNum []int
var GlobalFileSuffix []string

func init() {
	GlobalTypeList = []string{config.TIP_COLLECTION}
	GlobalFileType = []string{"application/zip"}
	GlobalFileSuffix = []string{".xlsx"}
	GlobalFileNum = []int{}
	for i := 0; i < len(GlobalTypeList); i++ {
		GlobalFileNum = append(GlobalFileNum, 0)
	}
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
