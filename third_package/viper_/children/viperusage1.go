package children

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

func ViperUsage1() {
	var v = viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}
	_m := make(map[string]interface{})
	err = v.Unmarshal(&_m)
	if err != nil {
		panic(fmt.Errorf("Fatal error marshal config file: %s \n", err))
	}
	_jsonString, _ := json.Marshal(_m)
	fmt.Println(string(_jsonString))
}
