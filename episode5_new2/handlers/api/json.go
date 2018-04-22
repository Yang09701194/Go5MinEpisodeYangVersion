package api

import (
	"fmt"
	"net/http"
)

//之前練習API的時候就知道  你拿ResposeWriter  去Write 就是有回應的效果
//所以這邊  可以把StatusCode Write到Header  然後在回應錯誤訊息
//直接呼叫這個方法就有回應的效果
//所以這個方法像是在api/canides.go裡面被使用 幾乎都是在邏輯結尾
//在中間的話就是  執行完就return  不然就是放最結尾
func jsonErr(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(
		`{"error": "%s"}`,
		err.Error())))
}
