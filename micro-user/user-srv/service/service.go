package service

import user "github.com/entere/go-demos/micro-user/user-srv/proto/user"

// Service ...
type Service interface {
	QueryUserByID(userID string) (*user.UserInfo, error)
}
