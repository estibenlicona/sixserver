package types

type Config struct {
	LoginPort   int
	LobbyPort   int
	NetworkPort int
	MainPort    int
	ServerIP    string
	Redis       RedisConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}
