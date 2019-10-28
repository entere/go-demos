package service

import (
	"github.com/entere/go-demos/micro-user/user-srv/dao"
	user "github.com/entere/go-demos/micro-user/user-srv/proto/user"
)

// UserService ...
type UserService struct {
	userDao dao.Dao
}

// NewUserService ...
func NewUserService(d dao.Dao) Service {
	return &UserService{
		userDao: d,
	}
}

// QueryUserByID ...
func (e *UserService) QueryUserByID(userID string) (ret *user.UserInfo, err error) {
	d, err := e.userDao.QueryUserByID(userID)
	if err != nil {
		return nil, err
	}
	return d, nil
}
