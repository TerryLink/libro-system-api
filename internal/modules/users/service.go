package users

// FetchAllBooks 獲取所有書籍
func FetchAllUsers() []User {
	return GetAllUsers()
}

// FetchBookByID 根據 ID 獲取書籍
func FetchUserByID(id string) (*User, error) {
	return GetUserByID(id)
}

// CreateNewBook 新增書籍
func RegisterNewUser(newBook User) {
	AddUser(newBook)
}
