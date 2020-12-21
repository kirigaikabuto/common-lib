package access_token_middleware

type RedisConfig struct {
	Host     string
	Password string
	Port     int
	DB       int
	SSL      bool
}
