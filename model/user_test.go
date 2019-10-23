package model

import (
	"testing"
)
var user User
func TestGetAllUser(t *testing.T)  {
	var  us[]User
	user.GetAllUser(&us)
	t.Log(us)
}
func TestGetUser(t *testing.T)  {
	user.Id =2
	u,_:= user.GetUser()
	t.Log(u)
}

func TestAddUser(t *testing.T) {
	var u User
	u.UserName ="ass"
	u.PassWord = "ss221s"
	user.AddUser()
	t.Log(u)
}

func TestUpUser(t *testing.T)  {
	user.Id =2
	u,_:= user.GetUser()
	t.Log(u)
	u.PassWord="22222222"
	u.UserName="22"
	u.UpdateUser()
	t.Log(u)
}

func TestDelUser(t *testing.T)  {
	user.Id =2
	u,_:= user.GetUser()
	t.Log(u)
	user.DeleteUser()
}
