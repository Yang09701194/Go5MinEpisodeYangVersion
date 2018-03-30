package handlers

import (
	//是 html/template耶!! 那有沒有 text/template???
	"html/template"
	"net/http"

	//這個要先get一下
	"github.com/unrolled/render"
)

// Renderer is the interface to rendering templates. This example repository
// has an unrolled/render (https://godoc.org/github.com/unrolled/render) implementation,
// but other implementations are certainly possible
type Renderer interface {
	// Render renders a template with data and passes the rendered result to w with the given code.
	// if layout is non-empty, this func renders the template with that name, and the templateName
	// is rendered under the {{yield}} block. otherwise this func renders the template called templateName
	Render(w http.ResponseWriter, code int, templateName string, data interface{}, layout string)
}

// RenderRenderer is a Renderer implementation that uses unrolled/render (https://godoc.org/github.com/unrolled/render#Render.HTML)
// to do template rendering
type RenderRenderer struct {
	r *render.Render
}

// NewRenderRenderer returns a new RenderRenderer, where the underlying render.Render
// serves templates out of dir with the given func map. it's configured in dev mode
// according to the dev boolean
func NewRenderRenderer(dir string, extensions []string, funcs []template.FuncMap, dev bool) *RenderRenderer {
	opts := render.Options{
		//我猜這是建構式
		Directory:     dir,
		Extensions:    extensions,
		Funcs:         funcs,
		IsDevelopment: dev,
	}
	return &RenderRenderer{
		//render是package名稱
		r: render.New(opts)}
}

// Render is the interface implementation
//這邊的interfac imletaion
//在C#裡  所謂的impltion 有兩個要件  虛跟實
//虛是代表   只宣告方法簽名在interface
//實是代表  class是實體  在擴充介面之後 趣將簽名實做

//但是Go 稍微特別了點  我本來以為差很多  結果沒有
//實際上也有兩個部分   虛實
//實 的部分 就是方法的參數  參數本身是一個實體的struct
//虛的話就有點跳痛  這就是差異之處
//  需他會直接用  方法名稱  和方法簽章  去自己找對應的interface  然後視為擴充
//也就是  C#裡面的  : IXX 這件事情  被自動做掉了
//這應該算是唯一差異的地方

//另一個問題會是 如果有兩個介面 裡面有名稱簽章完全相同的宣告
//這樣子如果幫某個struct實做這個簽章方法  那算不算同時實做了兩個介面
//這個我打算之後碰到再查  先放到evernote裡的question

func (r *RenderRenderer) Render(w http.ResponseWriter, code int, templateName string, data interface{}, layout string) {
	if layout != "" {
		//說來看到templateName  就會讓我想到
		//這應該是類似   先做好很多模板
		//然後就可以藉由templateName  直接指定套用
		//YES!!
		r.r.HTML(w, code, templateName, data, render.HTMLOptions{Layout: layout})
		return
	}

	//看以來  這個layout有重要含義
	r.r.HTML(w, code, templateName, data)
}

//YES!!搞定了 renderer.go 看起來我可以render了!!!
