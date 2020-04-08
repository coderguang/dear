package worldToVoice

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"httpServer/src/config"

	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/gorilla/websocket"
)

func TestConn() {
	authStr := assembleAuthUrl(config.GetTTSCfg().Host, config.GetTTSCfg().Apikey, config.GetTTSCfg().Apisec)
	sglog.Debug("str:", authStr)

	c, resp, err := websocket.DefaultDialer.Dial(authStr, nil)
	if err != nil {
		sglog.Error("create websocket error:", err, ",status code:", resp.StatusCode)
		return
	}
	defer c.Close()

	if resp.StatusCode != 101 {
		sglog.Error("read resp status code error,status", resp.StatusCode)
		return
	}

	sglog.Debug("receive resp status code,", resp.StatusCode)

	worlds := []byte("测试一下语音效果")

	msg := NewParam()
	msg.Common.AppID = config.GlobalTTCfg.Appid
	msg.Data.Status = 2
	//msg.Data.Text = base64.RawURLEncoding.EncodeToString(worlds)

	msg.Data.Text = base64.StdEncoding.EncodeToString(worlds)

	js, err := json.Marshal(msg)
	if err != nil {
		sglog.Error("param to json error,", err)
		return
	}

	sglog.Debug("js is ", string(js))

	done := make(chan struct{})

	go func() {
		defer close(done)
		voiceBytesList := [][]byte{}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				sglog.Error("read msg from server err,", err)
				return
			}
			//sglog.Debug("recv from server:", string(message))
			result := new(SResponResult)
			err = json.Unmarshal(message, &result)

			if err != nil {
				sglog.Error("parse data error")
				return
			}
			if result.Code != 0 {
				sglog.Error("transform world error,code=", result.Code)
				return
			}
			tmpBytes, err := base64.StdEncoding.DecodeString(result.Data.Audio)
			if err != nil {
				sglog.Error("transform result to base64 error,code=", err)
				return
			}
			voiceBytesList = append(voiceBytesList, tmpBytes)
			sglog.Info("recv from server,status:", result.Data.Status)
			if result.Data.Status == 2 {
				//sglog.Debug("final bytes:", str)

				voiceBytes := bytes.Join(voiceBytesList, []byte{})
				num, filename, err := sgfile.WriteFile("./data", "t.mp3", voiceBytes)

				if err != nil {
					sglog.Error("write file error", err)
					return
				}
				sglog.Debug("write file ok,num:", num, filename)

			}

		}

	}()

	err = c.WriteMessage(websocket.TextMessage, js)

	if err != nil {
		sglog.Error("push param to server,", err)
		return
	}

	sgcmd.StartCmdWaitInputLoop()

}
