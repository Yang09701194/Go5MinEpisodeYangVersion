package handlers

import "net/http"

//這個Index  還比較看的出來他在幹嘛  這個就是說 這個方法是給Render類別用的
func Index(ren Renderer) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//一樣  這個indez 就是去找 index.html  然後 layout 應該就是去套 layout 的 特殊雙括號html
			ren.Render(w, http.StatusOK, "index", struct{}{}, "layout")
		}
	)
}

//以檔案架構   index.go 和 candies.go 還滿接近的  據之前的經驗  就是一個頁面  專門由一個handlers底下的 go檔案來處理





