package models

import (	
	"encoding/json"
	"sync"
)

//被我猜對了!!!!! 
//跟之前開會我猜的一樣   DB藏在記憶體裡!!!!
//再來猜猜看  是不是一個list
type inMemDB struct{
	rwm *sync.RWMutex
	//這個滿有趣的  我猜他是一個map 
	//也就是一個字典  key是string  value是byte array 吧~
	m map[string][]byte
}

//說起來這個和mongo_db裡面的方法  命名方式有點類似
func NewInMemoryDB() DB{
	return &inMemDB{
		//其實這個滿有意思的  
		//你看它上面  宣告類別的時候  是沒有大括號的
		//現在到底下就有大括號    應該類似  要取實體的話  就要加大括號
		//這點之後也值得注意
		rwm : &sync.RWMutex{},  
		m :   make(map[string][]byte)
	}
}


//說實在  寫道這裡大概就知道   應該會跟mongo_db完全一樣了
func (db *inMemDB) GetAllKeys() ([]string, error){
	//讀鎖
	//說來實在很有趣  我之前在SO上發了一篇文章
	//問MSSQL  在極微小的時間內    如果有毒血衝突的時候怎麼處理
	//那時候想得很複雜   就算是EF  也會要用很複雜的方式實現
	//最後是有人  建議用primary key的方式去 做
	//保證不會重複

	//而從C#本身  依我平常的經驗   是幾乎沒碰觸到  讀寫鎖的Code
	//不過在mongo db 就輕而易舉的讓我看到了
	db.rwm.RLock()
	defer db.rwm.RUnlock()	
	//db的m是一個字典  key是string  value是byte array 吧~
	ret := make([]string, len(db.m))
		i := 0
	for key, _ := range db.m{
		ret[i] = key   //也就是 把key都記錄進來的意思
		i++
	}	
	return ret, nil
}


func (db *inMemDB) Get(key string, val Model)  error{
	//這可以說是 讀取時的 DB方法起手式
	//你等一下可以看到  如果是Set的時候  就是Lock
	//而不只是 RLock了
	db.rwm.RLock()
	//第二句起手式
	defer db.rwm.RUnlock()
	//這個我就看得懂了  這就是書裡面說的那個  檢查key在不在的  最安全方法
	b, ok := db.m[key]
	if !ok {
		//這個東西是存在於db.go裡面  所以我先去寫db.go
		return ErrNotFound 
	}
	//感覺有點像是 把 b轉成val
	return json.Unmarshal(b, val)
}
 
func (db *inMemDB) Set(key string, val Model) error{
	//簡而言之   Lock搭配 Unlock
	db.rwm.Lock()
	defer db.rwm.Unlock()
	//這句就像是 if (db.Keys.Contains(key))一樣  檢查key是否存在
  _, ok := db.m[key]	

	if !ok {
		return ErrNotFound
	}
	b, err := val.MarshalJSON()
	if err != nil{
		return err
	}
	//這個應該是set的動作吧
	db.m[key] = b
	return nil
}

func (db *inMemDb) Upsert(key string, val Model) (bool, error){
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	b, err := val.MarshalJSON()
	if err != nil{
		return false, err
	}
	db.m[key] = b
	return !ok, nil
}




//我可以先告訴你   這個檔案很神奇  他沒有Close  

