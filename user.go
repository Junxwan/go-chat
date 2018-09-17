package main

// 使用者結構
type user struct {
	ID       int
	Account  string `form:"account" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
}

// 已註冊的使用者
type account []user

var member account

// 新增會員
func (a account) add(name, account, password string) int {
	id := len(a) + 1

	member = append(a, user{
		ID: id,
		Account:  account,
		Password: password,
		Name:     name,
	})

	return id
}

// 檢查帳號是否正確
func (a account) exist(account, password string) bool {
	for _, u := range a {
		if u.Account == account && u.Password == password {
			return true
		}
	}
	return false
}

// 根據名稱取user
func (a account) get(id int) (user, bool) {
	for _, u := range a {
		if u.ID == id {
			return u, true
		}
	}
	return user{}, false
}

func (a account) getByName(name string) (user, bool) {
	for _, u := range a {
		if u.Name == name {
			return u, true
		}
	}
	return user{}, false
}

// 根據帳號取user
func (a account) getByAccount(account string) (user, bool) {
	for _, u := range a {
		if u.Account == account {
			return u, true
		}
	}
	return user{}, false
}
