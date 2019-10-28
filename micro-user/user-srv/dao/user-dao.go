package dao

import (
	"log"

	"github.com/entere/go-demos/micro-user/base/pkg/db"
	user "github.com/entere/go-demos/micro-user/user-srv/proto/user"
)

// UserDao ...
type UserDao struct {
}

// NewUserDao ...
func NewUserDao() Dao {
	return &UserDao{}
}

// QueryUserByID ...
func (e *UserDao) QueryUserByID(userID string) (ret *user.UserInfo, err error) {

	sql := "select user_id,mobile,avatar_url,gender,nickname,created_at,updated_at from users where user_id = ?"
	row := db.MyDB.Self.QueryRow(sql, userID)
	scanErr := row.Scan(&ret.UserID, &ret.Mobile, &ret.AvatarUrl, &ret.Gender, &ret.Nickname, &ret.CreatedAt, &ret.UpdatedAt)
	if scanErr != nil {
		log.Printf("% + v", scanErr)
		return
	}

	return
}
