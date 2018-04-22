package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


const (
	dbName      = "trickortreat"
	candiesColl = "candies"
)

// wrapper object to ensure all data in mongo has a key
type objWrapper struct {
	Key  string `bson:"key"`
	Data Model  `bson:"data"`
}

// MongoDB is a DB implementation that talks to an external MongoDB server.
// Note: this is UNTESTED code, only to be used as a getting started reference.
// make sure you test it thouroughly against a real MongoDB server before using it
// in production
type MongoDB struct {
	//第二次寫就會多記一些東西
	//這裡很清楚的差異是  InMemDB  就有一個lock  而且每次有對DB動作時  就需要先鎖
	//接著是  defer 解鎖
	//但是不知道為什麼mongo db就不用   我在想說不定這個套件裡面  就有實作好  lock
	//所以  在寫的時候就不用考慮
	sess *mgo.Session
}

// NewMongoDB connects to an external MongoDB server and returns a DB implementation
// that fronts that connection. call Close() on the returned value when done.
func NewMongoDB(urlStr string) (*MongoDB, error) {
	s, err := mgo.Dial(urlStr)
	if err != nil {
		return nil, err
	}
	//這句的重點應該會是  s  s是Dial新拿到的  我覺得對應到C# 的話  可能有點類似new SqlConnection拿到的連線
	//這邊的 &MongoDB  應該有點類似  new 一個新物件的感覺
	//但是惡裡有點有趣的是  他回傳值明明就是	 *MongoDB   但是 這裡return 竟然是用&
	//這個也加到問題區好了
	return &MongoDB{sess: s}, nil
}

func (m *MongoDB) GetAllKey() ([]string, error) {
	sess := m.sess.Copy()
	defer sess.Close() //說起來這說不定是mongodb的起手式  應該待會會常常看到

	/*說來也是怪 這個dbName應該是放在const裡  但是vscode竟然按下F12會無法偵測
	這也有可能是個bug*/

	//這個C  按下F12  可以查到他的註解是
// C returns a value representing the named collection.
//
// Creating this value is a very lightweight operation, and
// involves no network communication.
	coll := sess.DB(dbName).C(candiesColl)
	wrapper := []objWrapper{}
	if err := coll.Find(bson.M{}).All(&wrapper); err != nil {
		return nil, err
	}
	//說起來這個ret是什麼意思
	ret := make([]string, len(wrapper))
	for i, obj := range wrapper {
		ret[i] = obj.Key
	}
	return ret, nil
}

// Get is the interface implementation
func (m *MongoDB) Get(key string, val Model) error {
	sess := m.sess.Copy()
	defer sess.Close() //起手式
	coll := sess.DB(dbName).C(candlesColl)
	wrapper := objWrapper{}

	//這個wrapper不知道是不是有什麼特殊的約定  直接看這個程式碼的感覺
	//有點像是  丟進去   他會把找到的值  丟給wrapper  他是傳址嘛
	//所以丟進去如果有被賦值  回到這個方法的時候  改變的結果也會回來
	//不過詳細的約定方法是什麼 就還不確定
	if err := coll.Find(bson.M{"key": key}).One(&wrapper); err != nil {
		return err
	}
	val = wrapper.Data
	return nil
}

func (m *MongoDB) Set(key string, val Model) error {
	sess := m.sess.Copy()
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	return coll.Update(
		//這邊這個bson看起來就有點意思
		bson.M{"key": key}, objWrapper{Key: key, Data: nil})
}

func (m *MongoDB) Upsert(key string, val Model) (bool, error) {
	sess := m.sess.Copy()
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	cInfo, err := coll.Upsert(bson.M{"key": key}, objWrapper{Key: key, Data: nil})
	if err != nil {
		return false, err
	}
	// the Updated field is set when already existed, otherwise the UpsertedID field is set.
	// see the func at https://bazaar.launchpad.net/+branch/mgo/v2/view/head:/session.go#L18	

	//這個跟NH就有點像  這會讓我回想到對term insert的時候
	return cInfo.UpsertedId != nil, nil
}

// Close releases the underlying connections. always call this when
// completely done with operations, not before.
func (m *MongoDB) Close(){
	m.sess.Close()	
}



