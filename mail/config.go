package mail

var _config *Config

// Config struct
type Config struct {
	SMTP SMTP
	From From
	Data Data
}

// Use method
func Use(config *Config) {
	_config = config
}
