package main

import (
	handlers "./handlers"
	api "./handlers/api"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arschles/go-in-5-minutes/episode5/models"
	//這個util  和config有關  之後去研究一下怎麼用
	//我猜好像可以用  文字檔  或是下command的方式去設定
	"github.com/arschles/go-in-5-minutes/episode5/util"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	cfg := util.Config{Environment: "dev", Port: 33333}

	if err := envconfig.Process("trickortreat", &cfg); err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	env, err := cfg.Env()
	if err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	dev := env == util.EnvDev
	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)
	var db models.DB
	if dev {
		db = models.NewInMemoryDB()
	} else {
		d, err := models.NewMongoDB(cfg.MongoURL)
		if err != nil {
			log.Fatal("error connecting to Mongo [%s]", err)
			os.Exit(1)
		}
		db = d
	}

	r := mux.NewRouter()
	//這邊有點有趣  面不同的Routec對應後面的方法   後面的方法分別寫在handler資料夾底下
	//所以類似各個Route都開一個裡檔案到handlers資料夾底下
	r.Handle("/", handlers.Index(renderer)).Methods("Get")
	r.Handle("/candies", handlers.Candies(renderer)).Methods("Get")

	//話說這裡就是直接在使用  之前覺得很難的  api套件
	//這麼清楚的指引   讓我有如處於夢境之中!!
	//gorilla/mux
	r.Handle("/api/candies", api.Candies(db)).Methods("GET")
	r.Handle("/api/candy", api.CreateCandy(db)).Methods("PUT")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, http.StatusNotFound, "not_found", map[string]string{
			"url": r.URL.String(),
		}, "layout")
	})

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", cfg.Port)
	n.Run(hostStr)
}
