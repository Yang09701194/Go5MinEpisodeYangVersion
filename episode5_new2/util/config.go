package util

import "fmt"

//Env 環境變數
type Env string

func (e Env) String() string {
	return string(e)
}

const (
	EnvDev   Env = "dev"
	EnvStage Env = "stage"
	EnvProd  Env = "prod"
)

type Config struct {
	Port        int    `envconfig:"port"`
	Environment string `envconfig:environment`
	MongoURL    string `envconfig:mongo_url`
}

//這邊還滿有趣的  我覺得其實說不定回傳string 就可以了  不知道為什麼特別故意要弄個同型別異名
//也有可能是 把string改叫做Env  會比較好辨認吧

//不過下面這個我隨便想都知道可以優化為
// if( 傳進來的值  in  這三個值     )
//    直接回傳傳進來的值 和 nil
// 其他則   回傳 default值
//這樣其實三行就可以搞定  一半的程式碼就夠了
func (c Config) Env() (Env, error) {
	switch c.Environment {
	case EnvDev.String():
		return EnvDev, nil
	case EnvStage.String():
		return EnvStage, nil
	case EnvProd.String():
		return EnvProd, nil
	default:
		return Env(""), fmt.Errorf("invalid environment %s", c.Environment)
	}

}
