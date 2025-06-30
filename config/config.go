package config

var APIConfig = struct {
	AdminUser 			string
	AdminPassword 		string
	ServerPrivateKey 	[]byte
}{
	AdminUser:		"user",
	AdminPassword:	"pass",
}