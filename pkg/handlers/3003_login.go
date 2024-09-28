package handlers

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/core"
	"sixserver/pkg/database"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x3003(pkt types.Packet, conn gnet.Conn, cfg *types.Config) (out []byte, action gnet.Action) {

	ctx := pes6.GetConnectionContext(conn)

	cpr := core.NewCipher(cfg.CipherKey)
	clientRosterHash := getClientRosterHash(pkt)
	cpr.Block.Decrypt(clientRosterHash, clientRosterHash)
	userHash := getUserHash(pkt)

	redisConn := getRedisClient()
	defer func(conn *redis.Client) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing redis connection: %v", err)
		}
	}(redisConn)

	ctx.User = getUser(redisConn, userHash)
	ctx.User.Online = false

	if ctx.User.IsOnline() {
		data := PackCodeResponse(0xffffff11)
		err := pes6.SendPacketWithData(conn, 0x3004, data)
		helpers.HandleError(err)
		return
	} else if !isRosterHashValid(clientRosterHash) {
		data := PackCodeResponse(0xffffff12)
		err := pes6.SendPacketWithData(conn, 0x3004, data)
		helpers.HandleError(err)
		return
	} else {
		database.UpdateUserOnlineStatus(redisConn, ctx.User, true)
		err := pes6.SendPacketWithZeros(conn, 0x3004, 4)
		helpers.HandleError(err)
		return
	}
}

func PackCodeResponse(codeResponse uint32) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, codeResponse)
	if err != nil {
		log.Fatalf("binary.Write failed: %v", err)
	}

	return buf.Bytes()
}

func getRedisClient() *redis.Client {
	client := database.GetRedisConnection()
	return client
}

func getClientRosterHash(pkt types.Packet) []byte {
	return pkt.Data[48:64]
}

func getUserHash(pkt types.Packet) string {
	userHash := pkt.Data[32:48]
	return hex.EncodeToString(userHash)
}

func getUser(conn *redis.Client, hash string) *types.User {
	ctx := context.Background()
	userKey := "user:" + hash
	userData, err := conn.Get(ctx, userKey).Result()
	helpers.HandleError(err)

	var user types.User
	err = json.Unmarshal([]byte(userData), &user)
	helpers.HandleError(err)

	return &user
}

func isRosterHashValid(hash []byte) bool {
	if bytes.Contains(hash, []byte{0, 0, 0, 0}) {
		return false
	}
	return true
}
