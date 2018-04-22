package handlers

import "html/template"

//這個function直接看 無法直接看出他被用在哪裡
//只是在判斷  單數  or  多數
//先下個中斷點  看看它會不會在什麼時候進來

//其實更奇怪的點是
//這只是一個變數  也沒有放在   init之類的函數   所以這段到底會不會被執行  也算是一個謎
var funcs = []template.FuncMap{
	template.FuncMap(map[string]interface{}{
		"Pluralize": func(num int, singular, plural string) string {
			if num == 1 {
				return singular
			}
			return plural
		},
	}),
}
