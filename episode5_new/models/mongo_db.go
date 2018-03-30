package models

import (
	//應該是import mongo db的package
	//也就是整個mongo db的go檔案  只需要import這兩個package就可以運作
	//這點值得注意
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const(
	dbName			=	"trickortreat"
	candlesColl = "candles"
)

//wrapper object to ensure all data in mongo has a key
type objWrapper struct{
	Key string `bson:"key"`
	Data Model `bson:"data"`
}

//MongoDB is a DB implementation that talks to an external MongoDB server.
//Note: this is UNTESTED code, only to be used as a getting started reference.
//make sure you test it throughly against a real MongoDB server before using it
//in production
type MongoDB struct{
	sess *mgo.Session
}

// New MongoDB connects to an external MongoDB sever, and returns a DB implementation
// that fronts that connection. call Close() on the returned value when done.
func NewMongoDB(urlStr string) (*MongoDB, error){
	s, err := mgo.Dial(urlStr)
	if err != nil {
		return nil, err
	}
	return &MongoDB(sess: s), nil
}

func (m *MongoDB) GetAllKeys() ([]string, error) {
	sess := m.sess.Copy()
	defer sess.Close()
	//這邊的 C 跟 M 目前看不懂   感覺好像有什麼特殊含義
	coll := sess.DB(dbName).C(candiesColl)
	wrapper := []objWrapper{}
	if err := coll.Find(bson.M{}.All(&wrapper); err != nil{
		return nil, err
	}
	ret := make([]string, len(wrapper))
	for i, obj := range wrapper {
		ret[i] := obj.Key
	}
	return ret, nil
}

func (m *MongoDB) Get(key string, val Model) error{
	//應該類似  copy session的意思  session 應該類似 資料庫的代理人?
	//或是通訊的代言人
	sess := m.sess.Copy()
	//現在就看得懂  這相當於  C# 裡面的finally 裡面加上 Close
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	wrapper := objWrapper()
	if err := coll.Find(bson.M{"key": key}).One(&wrapper); err != nil{
		return err
	}
	//將傳入的val  賦值
	val = wrapper.Data
	//回傳無錯誤
	return nil
}

//Set is the interface implementation
func (m *MongoDB) Set(key string, val Model){
	//DB的方法的起手式就是這個  這也算是學到了一課
	sess := m.sess.Copy()
	//第二具固定就是defer 關閉  這也算是學到了一課
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	return coll.Update(bson.M{"key": key}, objWrapper(Key: key, Data: val))
}

//Upsert is the interface implementation
//其實寫到這裡  就很像是在  教怎麼寫C#的檔案DB的感覺了
//有點越來越順手的感覺
func (m *MongoDB) Upsert(key string, val Model) (bool, error) {
	//到這裡有點誤出了新體會
	//因為每次執行DB方法都會Close
	//不知道mongoDB Close之後能不能再開啟
	//如果可以的話  Copy有可能是類似Open的效果
	//不然就是這個m.sess永遠都開著   
	//每次Copy都會拿到一個open狀態的  open之後  
	//之後再close  原本的還是open  所以之後再copy  還是會拿到open狀態的
	//還有一種可能  因為這個m是傳進來的
	//所以有可能m  船進來前都要先另外做好開啟的動作
	sess := m.sess.Copy()
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	cInfo, err := coll.Upsert(bson.M{"key": key},
		 objWrapper{Key: key, Data: val})
	if err != nil {
		return false, err
	}
	//the Updated field is set when already existed, otherwise the UpsertedID field is set
	//這個有可能upsert成功  UpsertedId就不會是nil  
	//但是他與一有點不完全是這樣   留待之後驗證
	return cInfo.UpsertedID != nil, nil;
}

// Close releases the underlying connections. always call this when
// completely done with operations, not before
//簡而言之，就是最後再呼叫的意思
func (m *MongoDB) Close(){
	m.sess.Close()
}




