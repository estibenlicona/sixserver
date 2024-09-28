package database

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"sixserver/pkg/helpers"
	"sixserver/pkg/types"
)

func UpdateUserOnlineStatus(conn *redis.Client, user *types.User, isOnline bool) {
	if user != nil {
		user.Online = isOnline
		ctx := context.Background()
		userKey := "user:" + user.Hash

		userData, err := json.Marshal(user)
		helpers.HandleError(err)

		err = conn.Set(ctx, userKey, userData, 0).Err()
		helpers.HandleError(err)
	}
}
