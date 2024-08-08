package children

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

func ViperUsage2() {
	viper.SetEnvPrefix("SQ") // 正常来说这行代码应该在最开始
	viper.AutomaticEnv()     // 这样获取的不会存储在 viper 里面(也就是使用 marshal 得不到)
	// 这里的前缀仅仅是用来标识，最终存储的时候会删掉前缀 SQ_

	fmt.Println(viper.Get("ENV1"))
	fmt.Println(viper.Get("ENV2"))

	viper.SetDefault("k1", "v1") // 这样会存储在 viper 里面，这里的和 readInConfig 一样的

	viper.Set("k2", "v2")

	fmt.Println(viper.Get("k1"))

	fmt.Println(viper.Get("PATH"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error readInConfig config file: %s \n", err))
	}
	_m := make(map[string]interface{})
	err = viper.Unmarshal(&_m)
	if err != nil {
		panic(fmt.Errorf("Fatal error marshal config file: %s \n", err))
	}
	_jsonString, _ := json.Marshal(_m)
	fmt.Println(string(_jsonString))
}
