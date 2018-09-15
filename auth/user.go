package auth

// 使用者結構
type User struct {
	Account  string `form:"account" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
}

// 已註冊的使用者
type account []User

var Member account

// 新增會員
func (a account) add(name, account, password string) {
	Member = append(a, User{
		Account:  account,
		Password: password,
		Name:     name,
	})
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
func (a account) Get(name string) (User, bool) {
	for _, u := range a {
		if u.Name == name {
			return u, true
		}
	}
	return User{}, false
}

// 根據帳號取user
func (a account) GetByAccount(account string) (User, bool) {
	for _, u := range a {
		if u.Account == account {
			return u, true
		}
	}
	return User{}, false
}
