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

		sglog.Info("require data:", r.Form)
		sglog.Info("muti:", r.MultipartForm)

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

		filePath, filename, resultfile := getPathAndFileName(index)
		sgfile.AutoMkDir(filePath)
		sglog.Debug("receive file:", filename)

		// write file
		err = writefile(filename, fileBytes)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		sglog.Info("SUCCESS to commit file,please wait......\n")
		doLogic(w, index, filename, resultfile)
	})
}
