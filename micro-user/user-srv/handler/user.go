package handler

import (
	"context"
	"fmt"

	"github.com/entere/go-demos/micro-user/user-srv/dao"
	"github.com/entere/go-demos/micro-user/user-srv/service"

	user "github.com/entere/go-demos/micro-user/user-srv/proto/user"
)

// User ...
type User struct{}

// QueryUserByID ...
func (e *User) QueryUserByID(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResponse) error {
	userID := req.GetUserID()
	userService := service.NewUserService(dao.NewUserDao())

	userData, err := userService.QueryUserByID(userID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(userData)
	// log.Log("Received User.QueryUserByID request" + userID)
	// rsp.Msg = "Hello " + req.Name
	return nil
}
