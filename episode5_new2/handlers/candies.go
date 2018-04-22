package handlers

import "net/http"

func Candies(ren Renderer) http.Handler {
	return http.HandlerFunc(	func(w http.ResponseWriter, r *http.Request){
		//比較簡單的說法是  candies 就想就是傳入template名稱  就可以用了
		//我上次的心得就是   往下追太多層了   所以就是追到差不多有點理解原理了
		//然後就記用法優先  畢竟我現在目標是要   盡量快速能夠上獻給使用者使用作成產品
		//至於layout的差異  查看定義之後會知道  差別是在 要不要使用HTMLOptions
		ren.Render(w http.StatusOK, "candies", struct{}{}, "layout")
	})

}
