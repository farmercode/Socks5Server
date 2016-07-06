package config

import(
	"io/ioutil"
	"encoding/json"
)

type ServerConfig struct  {
	Ip string
	Listen string
}

//加载Json配置文件
//path为json格式配置文件路径,config为配置文件对象
func LoadConfig(path string,config interface{}) (error){
	content,error := ioutil.ReadFile(path)
	if error != nil {
		return error
	}
	error = json.Unmarshal(content,&config)
	if error != nil {
		return error
	}
	return nil
}
