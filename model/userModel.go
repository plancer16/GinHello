package model

type UserModel struct {
	Email 			string `form:"email" binding:"email"`//binding校验格式
	Password 		string `form:"password"`
	PasswordAgain 	string `form:"password-again" binding:"eqfield=Password"`
}
