package tipCollection

import "net/http"

func DoLogic(w http.ResponseWriter, filename string, resultfile string) error {
	flag := make(chan bool)
	go func() {
		err := StartParse(filename, resultfile, flag)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			str := "\nSUCCESS to deal xlsx file for tip collections,filename is " + resultfile
			w.Write([]byte(str))
		}
	}()
	<-flag
	return nil
}
