package dao

import (
	proto "github.com/entere/go-demos/micor-user/user-srv/proto/user"
)

type UserDao struct {
}

func (d *UserDao) QueryUserByID(name string) (ret *proto.User, err error) {
	return
}
