package httpHandle

import (
	"errors"
	"httpServer/src/config"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/coderguang/GameEngine_go/sglog"
)

func checkFileMaxSize(w http.ResponseWriter, r *http.Request) error {

	maxUploadSize := int64(config.GlobalCfg.MaxSize * 1024 * 1024)

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return err
	}
	return nil
}

func checkAllowFiles(w http.ResponseWriter, r *http.Request) (multipart.File, int, error) {
	for k, v := range GlobalTypeList {
		// parse and validate file and post parameters
		file, _, err := r.FormFile(v)
		if err == nil {
			return file, k, nil
		}
	}
	renderError(w, "INVALID_FILE_UPLOAD", http.StatusBadRequest)
	return nil, 0, errors.New("INVALID_FILE_UPLOAD no support upload file")
}

func checkFileTypeMatch(w http.ResponseWriter, r *http.Request, index int, file multipart.File) ([]byte, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		sglog.Error("read file from require error")
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		sglog.Error("read file error", err)
		return fileBytes, err
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	detectedFileType := http.DetectContentType(fileBytes)
	if detectedFileType != GlobalFileType[index] {
		sglog.Error("file not match")
		renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
	}
	return fileBytes, nil
}
