package models

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct{
	gorm.Model
	NickName   string      `gorm:"not null"`
	UserName   string      `gorm:"not null"`
	Age         int         
	Pwd         string      
	Emails      []string    `pg:",array"`
	Gender      int        
	Summary     string     
	Phone       string      `gorm:"unique"`
	IsDelete   string      
	OpenId     string      
	Locations   []*Location 
	AvatarUrl  string      
}

type Location struct{
	gorm.Model
}

type Like struct{
	gorm.Model	
}

type Comment struct{
	gorm.Model	
}

func Run() {

}