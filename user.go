package main

// 使用者結構
type user struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
}

// 已註冊的使用者
type account []user

var member account

// 新增會員
func (a account) add(name, username, password string) {
	member = append(a, user{
		Username: username,
		Password: password,
		Name:     name,
	})
}

// 檢查帳號是否正確
func (a account) exist(username, password string) bool {
	for _, u := range a {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
