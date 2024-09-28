package types

type Config struct {
	ServerIP    string
	LoginPort   int
	LobbyPort   int
	NetworkPort int
	MainPort    int
	CipherKey   string
	Redis       RedisConfig
	Roster      RosterConfig
}

type RosterConfig struct {
	EnforceHash bool
	CompareHash bool
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}
