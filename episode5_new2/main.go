package main

import (
	"log"

	"github.com/arschles/go-in-5-minutes/episode5/handlers"

	"github.com/arschles/go-in-5-minutes/episode5/models"
	"github.com/arschles/go-in-5-minutes/episode5/util"
	"github.com/kelseyhightower/envconfig"
)

//這邊要仔細注意 api的寫法 和邏輯
//盡量看看能不能  這次寫完之後  就把api的用法記起來
//然後可以用同樣的原理自己寫出完員不同邏輯的API
//成為AP的造物者
func main() {
	cfg := util.Config{Environment: "dev", Port: 33333}
	if err := envconfig.Process("trickortreat", &cfg); err != nil {
		log.FatalF("config error [%s]", err)
	}

	env, err := cfg.Env()
	if err != nil {
		log.Fatal("config error [%s]", err)
	}
	dev := env == util.EnvDev
	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)
	//
	var db models.DB

}
