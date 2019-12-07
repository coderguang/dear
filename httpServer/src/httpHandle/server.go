package httpHandle

import (
	"net/http"

	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sglog"
)

func UploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		err := checkFileMaxSize(w, r)
		if err != nil {
			sglog.Error("check size error", err)
			return
		}
		// parse and validate file and post parameters
		file, index, err := checkAllowFiles(w, r)
		if err != nil {
			sglog.Error("check all file error", err)
			return
		}
		defer file.Close()

		fileBytes, err := checkFileTypeMatch(w, r, index, file)
		if err != nil {
			return
		}

		filePath, filename := getPathAndFileName(index)
		sgfile.AutoMkDir(filePath)
		sglog.Debug("receive file:", filename)

		// write file
		newFile, err := sgfile.Create(filename)
		if err != nil {
			sglog.Error("crete file error,", err)
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			sglog.Error("write to file error")
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS to commit file,please wait"))
	})
}
