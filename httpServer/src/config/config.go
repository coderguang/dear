package config

import "github.com/coderguang/GameEngine_go/sgcfg"

func init() {
	GlobalCfg = new(Cfg)
}

var GlobalCfg *Cfg

type Cfg struct {
	Port         int    `json:"port"`
	MaxSize      int    `json:"maxSize"`
	UploadPath   string `json:"uploadPath"`
	DownloadPath string `json:"downloadPath"`
	Upload       string `json:"upload"`
	Download     string `json:"download"`
}

func InitCfg(filename string) error {

	err := sgcfg.ReadCfg(filename, &GlobalCfg)
	if err != nil {
		return err
	}
	return nil
}
