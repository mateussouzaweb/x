package mail

var _config *Config

// Config struct
type Config struct {
	SMTP      SMTP
	Addresses []Address
	Data      Data
}

// Use method
func Use(config *Config) {
	_config = config
}

// Get address by reference
func Get(reference string) Address {

	var address Address
	for _, item := range _config.Addresses {
		if item.Reference == reference {
			address = item
			break
		}
	}

	return address
}
