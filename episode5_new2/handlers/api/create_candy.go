package api

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func CreateCandy(db models.DB) http.Handler {
	type jsonInput struct {
		Name string `json:"name"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//把Request的Body  解析為 in 的json物件  先說為Candy物件好了
		in := jsonInput{}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			jsonErr(w, http.StatusBadRequest, err)
			return
		}
		//解析後的in 應該已經昇華成為了一個Candy物件   去取它的名稱  然後做 insert or update
		// = upsert
		candy := models.Candy{Name: in.Name}
		if _, err := db.Upsert(bson.NewObjectId().String(), candy); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
		}
		//回傳 create success
		w.WriteHeader(http.StatusCreated)
	})
}
