package httpHandle

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

var GlobalTypeList []string
var GlobalFileType []string
var GlobalFileNum []int
var GLobalFileSuffix []string

func init() {
	GlobalTypeList = []string{"tipCollection"}
	GlobalFileType = []string{"application/zip"}
	GLobalFileSuffix = []string{".xlsx"}
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
