package users

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Login(accountName string, password string) (*JWTToken, error) {
	return u.repo.Login(accountName, password)
}

// get all users
func (u *UserService) FetchAllUsers() ([]User, error) {
	return u.repo.GetAllUsers()
}

// get user by id
func (u *UserService) FetchUserByID(id string) (*User, error) {
	return u.repo.GetUserByID(id)
}

// create new user
func (u *UserService) RegisterNewUser(newUser User) (*User, error) {
	return u.repo.AddUser(newUser)
}

func (u *UserService) DeleteUser(id string) (bool, error) {
	return u.repo.DeleteUser(id)
}

func (u *UserService) UpdateUser(user User) (*User, error) {
	return u.repo.UpdateUser(user)
}

func (u *UserService) SearchByEmailOrAccountName(accountName string, email string) (*User, error) {
	return u.repo.FindByEmailOrAccountName(accountName, email)
}
