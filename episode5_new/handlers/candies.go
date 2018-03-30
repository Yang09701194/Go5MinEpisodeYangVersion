package handlers

import "net/http"

func Candies(ren Renderer) http.Handler {
	//這有意思  我看到這個很有感覺  這個好像跟api有關係
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.Render(
			w, //現在參數的含義一一明瞭  這是ResponseWriter
			//code就是httpStatusCode
			http.StatusOK,
			//這是template名稱
			"candies",
			//這我是不確定為什麼需要兩個大括號
			//還有這邊是要求傳入一個interface
			//但是他竟然是傳入一個  感覺無意義的interface
			//不禁讓人好奇這個參數是不是根本沒用
			struct{}{},

			//實體方法在renderer.go  裡面只判斷!= ""
			//所以我猜 應該有特別的功效
			"layout")
	})
}
