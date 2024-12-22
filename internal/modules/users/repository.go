package users

import "errors"

// 模擬數據庫
var test_users = []User{
	{ID: "1", Username: "Golang", Email: "Google", Password: "abc"},
	{ID: "2", Username: "Java", Email: "Oracle", Password: "abcd"},
	{ID: "3", Username: "Python", Email: "Python Software Foundation", Password: "abcdefg"},
}

// GetBooks 獲取所有書籍
func GetAllUsers() []User {
	return test_users
}

// GetBookByID 根據 ID 獲取書籍
func GetUserByID(id string) (*User, error) {
	for i, b := range test_users {
		if b.ID == id {
			return &test_users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

// AddBook 新增書籍
func AddUser(newUser User) (*User, error) {
	test_users = append(test_users, newUser)
	return &newUser, nil
}

func DeleteUser(id string) error {
	for i, b := range test_users {
		if b.ID == id {
			test_users = append(test_users[:i], test_users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
