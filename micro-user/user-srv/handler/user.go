package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	user "github.com/entere/go-demos/micro-user/user-srv/proto/user"
)

// User ...
type User struct{}

// QueryUserByID ...
func (e *User) QueryUserByID(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResponse) error {
	userID := req.GetUserID()
	log.Log("Received User.QueryUserByID request" + userID)
	// rsp.Msg = "Hello " + req.Name
	return nil
}
