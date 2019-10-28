package dao

import user "github.com/entere/go-demos/micro-user/user-srv/proto/user"

// Dao ...
type Dao interface {
	QueryUserByID(userID string) (*user.UserInfo, error)
}
