package mail

var connection *SMTP

// Setup method
func Setup(config *SMTP) {
	connection = config
}
