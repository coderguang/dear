package wordToVoiceEx

import (
	"httpServer/src/config"
	"io/ioutil"

	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgtts/wordToVoice"
)

func StartParse(filename string, resultfile string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		sglog.Error("word to voice read file:", filename, " error:", err)
		return err
	}
	s := string(b)

	param := wordToVoice.NewParam()
	param.Business.Vcn = "x2_xiaoyuan"
	voiceBytes, err := wordToVoice.WorldToVoice(s, config.GetTTSCfg().Host, config.GetTTSCfg().Appid, config.GetTTSCfg().Apikey, config.GetTTSCfg().Apisec, param)
	if err != nil {
		sglog.Error("word to voice error:", err)
		return err
	}
	filePath, _ := sgfile.GetPath(resultfile)
	tmpfilename, _ := sgfile.GetFileName(resultfile)
	num, filename, err := sgfile.WriteFile(filePath, tmpfilename, voiceBytes)
	if err != nil {
		sglog.Error("word to voice write file error", err)
		return err
	}
	sglog.Debug("word to voice write file ok,num:", num, filename)

	return nil
}
