package config

import "github.com/coderguang/GameEngine_go/sgcfg"

func init() {
	GlobalCfg = new(Cfg)
}

var GlobalCfg *Cfg

type Cfg struct {
	Port int `json:"port"`
}

func InitCfg(filename string) error {

	err := sgcfg.ReadCfg(filename, &GlobalCfg)
	if err != nil {
		return err
	}
	return nil
}
