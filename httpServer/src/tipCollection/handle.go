package tipCollection

import "net/http"

func DoLogic(w http.ResponseWriter, filename string, resultfile string) error {
	flag := make(chan bool)
	go StartParse(filename, resultfile, flag)
	<-flag
	w.Write([]byte("SUCCESS to deal xlsx file for tip collections"))
	return nil
}
