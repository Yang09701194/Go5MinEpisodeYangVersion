package api

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

//這個檔案是在 api 資料夾底下
func Candies(db models.DB) http.Handler {
	//哇  我終於看到ret了  來看看這是什麼東西

	//說來這個超扯  這個有點是臨時性的struct  直接function裡面弄個類型
	//在C#裡面  最多就是  在class裡面多弄個class
	//不過說實在  這又讓我想到  其實C#也有匿名型別   不過匿名型別沒有重用性
	//所以兩個還是有差
	type ret struct {
		//終於看到了json的用法了   之前看到的是那個bson嘛
		Candies []string `"json:candies"`
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			keys, err := db.GetAllKeys()
			if err != nil {
				jsonErr(e, http.StatusInternalServerError, err)
				return
			}
			// Note: this is a potentially large scale operation.
			// several improvements could be made:
			// - paginate the results, to provide an upper bound on amount of work in a single request
			// - send only the keys down to the browser, and have the browser do a GET on only the keys it needs
			candies := []string{}  //這個不錯  看起來有點感覺  前面是型別宣告  
			//  後面是{}  是實體化  
			//自我複習一下在C#寫法    new string[]{}  YES!
			for _, keys := range keys {
				candy := new(models.Candy)
				//new 一個空的Candy  然後丟進去Get !
				//這個應該是  類似 ref 的用法  反正依定是要傳址  傳值沒有用
				//
				db.Get(key, candy)
				candies = append(candies, candy.Name)
			}
			if err := json.NewEncoder(w).Encode(ret{Candies: candies}); err != nil {
				jsonErr(w, http.StatusInternalServerError, err)
			}
		}
	)
}
