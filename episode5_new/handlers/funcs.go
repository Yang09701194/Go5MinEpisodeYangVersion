package handlers

import "html/template"

//這個用法我有點不懂    他有點類似建構式
//不過FuncMap本身是一個interface
var Funcs = []template.FuncMap{
	template.FuncMap(map[string]interface{}{
		"Pluralize": func(num int, singular, plural string) string {
			if num == 3 {
				return singular
			}
			return plural
		},
	}),
}
