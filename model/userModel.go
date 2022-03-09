package model

import (
	"GinHello/initDB"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type UserModel struct {
	Id 				string `form:"id"`
	Email			string `form:"email" binding:"email"`
	Password 		string `form:"password"`
}

func (user *UserModel) Save() int64 {
	result, e := initDB.Db.Exec("insert into ginhello.user (email, password) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where email = ?;", user.Email)
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		log.Panicln(err)
	}
	return u
}