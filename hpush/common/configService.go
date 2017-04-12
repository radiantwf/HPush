package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ConfigService 定义
type ConfigService struct {
	configStruct *ConfigStruct
}

// ConfigJSONStruct 定义
type ConfigJSONStruct struct {
}

// ConfigStruct 定义
type ConfigStruct struct {
}

// ProvinceInfoStruct 定义
type ProvinceInfoStruct struct {
	province string
	area     AreaStruct
}

// AreaStruct 定义
type AreaStruct struct {
	longitude [2]float64
	latitude  [2]float64
}

// NewConfig 定义
func NewConfig() (config ConfigService, err error) {
	var jsonStruct ConfigJSONStruct
	config.configStruct = new(ConfigStruct)
	err = config.loadJSONFile(&jsonStruct)
	if err != nil {
		return
	}
	StructDeepCopy(&jsonStruct, config.configStruct)

	fmt.Println("当前配置文件信息为：")
	fmt.Println(*config.configStruct)
	return
}

// loadJSONFile 定义
func (config *ConfigService) loadJSONFile(jsonStruct *ConfigJSONStruct) (err error) {
	var jsonStr []byte
	jsonStr, err = ioutil.ReadFile("./resources/config/config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonStr, jsonStruct)
	return
}
