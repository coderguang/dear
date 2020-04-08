package httpHandle

import (
	"crypto/rand"
	"fmt"
	"httpServer/src/config"
	"net/http"
)

var GlobalTypeList []string
var GlobalFileType []string
var GlobalFileNum []int
var GlobalFileSuffix []string
var GlobalResultFileSuffix []string

func init() {
	//三个需要相同长度
	GlobalTypeList = []string{config.TIP_COLLECTION, config.TIP_COLLECTION_EX, config.WORLD_TO_VOICE}
	GlobalFileType = []string{"application/zip", "application/zip", "text/plain; charset=utf-8"}
	GlobalFileSuffix = []string{"xlsx", "xlsx", "txt"}
	GlobalResultFileSuffix = []string{"xlsx", "xlsx", "mp3"}
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
