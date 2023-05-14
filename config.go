package urleap

type Config struct {
	Token string
}

func NewConfig(token string) *Config {
	return &Config{Token: token}
}
