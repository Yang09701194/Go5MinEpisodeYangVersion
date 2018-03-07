package models

import "encoding/json"

//所以  其實Model  是一個interface
//我剛剛先寫mongo_db那個檔案  Model都沒有標是顏色
//就是為我沒有先寫這個檔案
//所以之後  趙昭的時候如果遇到  白色的字
//就是可以先 用F12找到他的Reference
//然後 先把它的定義 先寫好  這樣子的話
//之後寫起來就有顏色  就感覺會比較好!!!!!!!

//這個檔  超簡單的  我一下子就寫好了
//往下個檔前進
type Model interface {
	json.Marshaler
}
