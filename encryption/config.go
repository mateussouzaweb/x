package encryption

var _config *Config

// Config struct
type Config struct {
	PrivateKey string
	PublicKey  string
}

// Use method
func Use(config *Config) {
	_config = config
}
