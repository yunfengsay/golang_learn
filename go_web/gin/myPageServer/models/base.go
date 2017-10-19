package models

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "github.com/go-pg/pg"
)

type User struct {
	Id        int      `json: id`
	Nick_Name string   `json: nick_name`
	User_Name string   `json: user_name`
	Age       int      `json: age`
	Pwd       string   `json: pwd`
	Emails    []string `json: emails`
	Urls      []Urls   `json: urls`
}

type Comment struct {
	Id      int   `json: id`
	Author  *User `json: author`
	From    int   `json: from`
	To      int   `json: to`
	About   int   `json: about`
	Content int   `json: content`
}

type Urls struct {
	Id      int     `json: id`
	Url     string  `json: url`
	Tags    []Tags  `json: tags`
	Score   float64 `json: score`
	Title   string  `json: title`
	Summary string  `json: summary`
}

type Tags struct {
	Id        int    `json: id`
	Tag       string `json: tag`
	Create_By *User  `json: create_by` // who create this tag
	Click_Num int64  `json: click_num` // 点击次数
}

func (u User) String() string {
	return fmt.Sprintf("Persion<%d %s %d>", u.Id, u.Nick_Name, u.Age)
}
func Run() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mypage").C("urls")
	err = c.Insert(&Tags{1, "学习", &User{}, 0})
	if err != nil {
		log.Fatal(err)
	}
	result := Tags{}
	err = c.Find(bson.M{"tag": "学习"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tag:", result)
}

// func createSchema(db *pg.DB) error {

// 	for _, model := range []interface{}{&User{}, &Comment{}, &Urls{}, &Tags{}} {
// 		err := db.CreateTable(model, nil)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	db := pg.Connect(&pg.Options{
// 		User: "zhangyunfeng",
// 	})
// 	createSchema(db)
// }
