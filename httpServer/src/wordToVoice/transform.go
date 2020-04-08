package wordToVoice

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"httpServer/src/config"
	"io/ioutil"
	"log"
	"strings"

	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/gorilla/websocket"
)

func StartParse(filename string, resultfile string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("world to voice read file:", filename, " error:", err)
		return err
	}
	s := string(b)
	//sglog.Debug("txt file str:", s)
	replaceStr := strings.Replace(s, "，", ",", -1)
	rawlist := strings.Split(replaceStr, ",")

	strlist := []string{}
	tmpstr := ""
	for _, v := range rawlist {
		if len(tmpstr)+len(v) > WORD_MAX_LEN {
			strlist = append(strlist, tmpstr)
			tmpstr = v
		} else {
			tmpstr += "," + v
		}
	}
	strlist = append(strlist, tmpstr)

	sglog.Info("world to voice txt split to ", len(strlist), " part")

	err = transWordToVoice(strlist, resultfile)

	return err
}

func transWordToVoice(strlist []string, resultfile string) error {

	authStr := assembleAuthUrl(config.GetTTSCfg().Host, config.GetTTSCfg().Apikey, config.GetTTSCfg().Apisec)

	voiceBytesList := [][]byte{}
	receiveNum := 0
	needReceiveNum := len(strlist)

	for _, v := range strlist {

		worlds := []byte(v)

		msg := NewParam()
		msg.Common.AppID = config.GlobalTTCfg.Appid
		msg.Data.Status = 2
		msg.Data.Text = base64.StdEncoding.EncodeToString(worlds)
		js, err := json.Marshal(msg)
		if err != nil {
			sglog.Error("param to json error,", err)
			return err
		}
		//sglog.Debug("js is ", string(js))

		c, resp, err := websocket.DefaultDialer.Dial(authStr, nil)
		if err != nil {
			sglog.Error("create websocket error:", err, ",status code:", resp.StatusCode)
			return err
		}
		defer c.Close()

		if resp.StatusCode != 101 {
			sglog.Error("read resp status code error,status", resp.StatusCode)
			return err
		}

		receiveAllFlag := make(chan bool)

		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					sglog.Error("read msg from server err,", err)
					break
				}
				//sglog.Debug("recv from server:", string(message))
				result := new(SResponResult)
				err = json.Unmarshal(message, &result)
				if err != nil {
					sglog.Error("parse data error")
					break
				}
				if result.Code != 0 {
					sglog.Error("transform world error,code=", result.Code)
					break
				}
				tmpBytes, err := base64.StdEncoding.DecodeString(result.Data.Audio)
				if err != nil {
					sglog.Error("transform result to base64 error,code=", err)
					break
				}
				voiceBytesList = append(voiceBytesList, tmpBytes)
				//sglog.Info("recv from server,status:", result.Data.Status)
				if result.Data.Status == 2 {
					receiveNum++
					sglog.Debug("receive part success,current recv:", receiveNum, ",need receive:", needReceiveNum)
					if receiveNum == needReceiveNum {
						voiceBytes := bytes.Join(voiceBytesList, []byte{})
						filePath, _ := sgfile.GetPath(resultfile)
						tmpfilename, _ := sgfile.GetFileName(resultfile)
						num, filename, err := sgfile.WriteFile(filePath, tmpfilename, voiceBytes)
						if err != nil {
							sglog.Error("write file error", err)
							break
						}
						sglog.Debug("write file ok,num:", num, filename)
					}
					break
				}
			}
			receiveAllFlag <- true
		}()

		err = c.WriteMessage(websocket.TextMessage, js)
		if err != nil {
			sglog.Error("push param to server,", err)
			return err
		}

		<-receiveAllFlag
	}

	return nil

}
