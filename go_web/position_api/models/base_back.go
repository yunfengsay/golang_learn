package models_back

// "github.com/go-pg/pg"
import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	Id          int         `json: id`
	Nick_Name   string      `json: nick_name`
	User_Name   string      `json: user_name`
	Age         int         `json: age`
	Pwd         string      `json: pwd`
	Emails      []string    `pg:",array"`
	Create_Date *time.Time  `json: create_date`
	Update_Date *time.Time  `json: update_date`
	Gender      int         `json: gender`
	Summary     string      `json: summary`
	Phone       string      `json: phone`
	Is_Delete   string      `json: is_delete`
	Open_Id     string      `json: open_id`
	Locations   []*Location `json: locations`
	Avatar_Url  string      `json: avatar_url`
}
type Location struct {
	L_Type []int    `json: l_type`
	imgs   []string `json: imgs`
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
