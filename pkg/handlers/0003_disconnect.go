package handlers

import (
	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/database"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x0003(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	ctx := pes6.GetConnectionContext(conn)

	redisConn := getRedisClient()
	defer func(conn *redis.Client) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing redis connection: %v", err)
		}
	}(redisConn)

	database.UpdateUserOnlineStatus(redisConn, ctx.User, false)
	return
}
