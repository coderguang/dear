package tipCollection

import "net/http"

func DoLogic(w http.ResponseWriter, filename string, resultfile string) error {
	flag := make(chan bool)
	go StartParse(filename, resultfile, flag)
	<-flag
	str := "\nSUCCESS to deal xlsx file for tip collections,filename is " + resultfile
	w.Write([]byte(str))
	return nil
}
