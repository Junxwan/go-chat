package main

// 使用者結構
type user struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
}

// 新增會員
func (a member) addUser(name, username, password string) {
	account.Account = append(a.Account, user{
		Username: username,
		Password: password,
		Name:     name,
	})
}

// 檢查帳號是否正確
func (a member) isUser(username, password string) bool {
	for _, u := range a.Account {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
