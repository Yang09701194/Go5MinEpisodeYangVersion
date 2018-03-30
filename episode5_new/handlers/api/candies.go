package api

import (
	"encoding/json"
	"net/http"

	//說來這還滿有趣的  他竟然  自己飲用自己的package
	//不過說來也正確  就算是在自己的package裡
	//不同package  go 也沒有共用機制
	//就像是 就算在 C#裡面 namesapce不同的話
	//還是要多寫一個using以示引用
	"github.com/arschles/go-in-5-minutes/episode5/models"
)

//原來如此，我總算懂了，這裡的models 不是之前寫的那個大Model
//而是上面import的 package 名稱的結尾 ----  models!!!!
func Candies(db models.DB) http.Handler {
	type ret struct {
		Candies []string `json:"candles"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, err := db.GetAllKeys()
		if err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}

		// Note: this is a potentially large scale operation.
		// several improvements could be made:
		// -(標頁數) paginate the results, to provide an upper bound on amount of work in a single request
		// - send only the keys down to the browser, and have the browser do a GET on only the keys it needs
		candies := []string{}
		for _, key := range keys {
			candy := new(models.Candy) //一個只含有string 的一個struct
			//這邊我是有點搞不懂  他candy是接一個 Model
			//但是這個上面的new 感覺應該只是一般初始 slice map的語法
			//model.Candy是一個只包含string 的struct
			//但是Model  裡面有個奇怪的Marshal
			// type Model interface { json.Marshaler
			// 所以為什麼這樣子可以接的到 就有點匪夷所思
			//這個已經加到evernote去了
			db.Get(key, candy)
			candies = append(candies, candy.Name)
		}
		if err := json.NewEncoder(w).Encode(ret{Candies: candies}); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}
