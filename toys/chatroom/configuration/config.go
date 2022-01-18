package configuration

type serviceConfig struct {
	AppMode, Host, Port string
}

type redisConfig struct {
	Host, Port, Password, DBName string
}

type mariadbConfig struct {
	Host, Port, User, Password, Name, CharSet, ParseTime string
}

type mongodbConfig struct {
	Host, Port, Password, DBName string
}

var (
	
)
