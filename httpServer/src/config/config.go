package config

import "github.com/coderguang/GameEngine_go/sgcfg"

func init() {
	GlobalCfg = new(Cfg)
	GlobalTTCfg = new(SttCfg)
}

var GlobalCfg *Cfg
var GlobalTTCfg *SttCfg

type Cfg struct {
	Port         int    `json:"port"`
	MaxSize      int    `json:"maxSize"`
	UploadPath   string `json:"uploadPath"`
	DownloadPath string `json:"downloadPath"`
	Upload       string `json:"upload"`
	Download     string `json:"download"`
}

type SttCfg struct {
	Appid  string `json:"appid"`
	Apikey string `json:"apikey"`
	Apisec string `json:"apisec"`
	Host   string `json:"host"`
}

func InitCfg(filename string) error {

	err := sgcfg.ReadCfg(filename+"config.json", &GlobalCfg)
	if err != nil {
		return err
	}

	err = sgcfg.ReadCfg(filename+"tts.json", &GlobalTTCfg)
	if err != nil {
		return err
	}

	return nil
}

func GetTTSCfg() *SttCfg {
	return GlobalTTCfg
}
