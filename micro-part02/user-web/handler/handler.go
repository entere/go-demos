package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	user "github.com/entere/go-demos/micro-part02/user-srv/proto/user"
	"github.com/micro/go-micro/client"
	// user "path/to/service/proto/user"
)

// UserCall ...
func UserCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	userClient := user.NewUserService("mu.micro.ci.srv.user", client.DefaultClient)
	rsp, err := userClient.Call(context.TODO(), &user.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
