package models

import (
	"encoding/json"
	"sync"
)

//簡言之   一個鎖 +  一個 key val 的 map = InMemDB
type InMemDB struct {
	rwm *sync.RWMutex
	m   map[string][]byte
}

func NewInMemoryDB() DB {
	return &InMemDB{
		rwm: &sync.RWMutex{},
		m:   make(map[string][]byte),
	}
}

//剛寫完 db 的 interface  就很有fel  一看就知道ˇ  現在在實作  DB 這個interface
func (db *InMemDB) GetAllKey([]string, error) {
	db.rwm.RLock()
	defer db.rwm.RUnlock() //有沒有  起手式
	ret := make([]string, len(db.m))
	// 看到這種寫法就會覺得很奇怪  C# 裡面單純讀迴圈  幾乎都不需要這樣特別弄一個int變數
	// 不過像是報表是有用到
	i := 0
	for key, _ := range db.m {
		ret[i] = key
		i++
	}
	return ret, nil
}

func (db *InMemDB) Get(key string, val Model) error {
	db.rwm.RLock()
	defer db.rwm.RUnlock() //有沒有  起手式
	b, ok := db.m[key]
	if !ok {
		return ErrNotFound
	}
	return json.Unmarshal(b, val)
}

func (db *InMemDB) Set(key string, val Model) error {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	if !ok {
		return ErrNotFound
	}
	b, err := val.MarshalJSON()
	if err != nil {
		return err
	}
	db.m[key] = b
	return nil
}

func (db *InMemDB) Upsert(key string, val Model) (bool, error) {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	b, err := val.MarshalJSON()
	if err != nil {
		return false, err
	}
	db.m[key] = b
	//根據 DB介面的註解  false是update  true是insert  可以對應以下驗證
	return !ok, err
}
