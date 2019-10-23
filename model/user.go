package model

type User struct {
	Id    int     `gorm:"column:id" json:"id"`
	UserName string  `gorm:"column:user_name" json:"username"`
	PassWord string  `gorm:"column:pass_word" json:"password"`
	Age int  `gorm:"column:age;default:22" json:"age"`
}
func (User) TableName() string {
	return "hk_user"
}

func (user *User) GetAllUser(users  *[]User) {
	DB.Find(users)
}

func (user *User) GetUser()(User,error)  {
	result := DB.First(&user, user.Id)
	return *user, result.Error
}

func (user *User) AddUser(){
	DB.Create(user)
}

func (user *User) UpdateUser(){
	DB.Save(user)
}

func (user *User) DeleteUser()  {
	DB.Delete(user)
}

