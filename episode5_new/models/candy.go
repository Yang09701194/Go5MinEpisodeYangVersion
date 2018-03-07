package models

import "encoding/json"

type Candy struct {
	Name string `bson:"name"`
}

//我 Q  這個方法竟然寫在這裡  竟然是自己寫的
func (c Candy) MarshalJSON() ([]byte, error) {
	//這應該是Serialize的意思 也就是物件轉字串
	//所以可以看出  回傳值其實是 byte[]
	return json.Marshal(map[string]string{"name": c.Name})
}
