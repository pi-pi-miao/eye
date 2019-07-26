package initialize

import "testing"

func TestNeweye(t *testing.T) {
	Neweye("./configs/application.toml")
	Testeye_InitConfig(t)
	Testeye_InitLog(t)
	Testeye_InitAll(t)
}


func Testeye_InitAll(t *testing.T) {
	Neweye("/Users/wanglei/Desktop/eye/configs/application.toml").InitAll()
}

func Testeye_InitConfig(t *testing.T) {
	Neweye("/Users/wanglei/Desktop/eye/configs/application.toml").InitConfig()
}

func Testeye_InitLog(t *testing.T) {
	Neweye("/Users/wanglei/Desktop/eye/configs/application.toml").InitConfig().InitLog()
}