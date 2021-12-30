package configuration

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/memorandum/cache"
	"github.com/saint-yellow/go-pl/toys/memorandum/model"
	"gopkg.in/ini.v1"
)

type serviceConfig struct {
	AppMode, Host, Port string
}

type cacheConfig struct {
	Host, Port, Password, DBName string
}

type databaseConfig struct {
	Host, Port, User, Password, Name, CharSet, ParseTime string
}

type jwtConfig struct {
	Secret string
}

func Service() serviceConfig {
	return sc
}

func Cache() cacheConfig {
	return cc
}

func JWT() jwtConfig {
	return jc
}

var (
	sc serviceConfig
	cc cacheConfig
	dc databaseConfig
	jc jwtConfig
)

func Init() {
	file, err := ini.Load("./configuration/config.ini")
	if err != nil {
		log.Fatalf("failed to load configuration: %s\n", err)
	}

	loadConfig(file)

	gin.SetMode(sc.AppMode)
	
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s", 
		dc.User, dc.Password, dc.Host, dc.Port, dc.Name, dc.CharSet, dc.ParseTime)
	model.InitializeDatabase(connString)

	cache.InitializeRedisClient(cc.Host, cc.Port, cc.Password, cc.DBName)
}


func loadConfig(file *ini.File) {
	sc = serviceConfig{
		AppMode: file.Section("service").Key("AppMode").String(),
		Host: file.Section("service").Key("Host").String(),
		Port: file.Section("service").Key("Port").String(),
	}

	cc = cacheConfig{
		Host: file.Section("cache").Key("Host").String(),
		Port: file.Section("cache").Key("Port").String(),
		Password: file.Section("cache").Key("Password").String(),
		DBName: file.Section("cache").Key("DBName").String(),
	}

	dc = databaseConfig{
		Host: file.Section("database").Key("Host").String(),
		Port: file.Section("database").Key("Port").String(),
		User: file.Section("database").Key("User").String(),
		Password: file.Section("database").Key("Password").String(),
		Name: file.Section("database").Key("Name").String(),
		CharSet: file.Section("database").Key("CharSet").String(),
		ParseTime: file.Section("database").Key("ParseTime").String(),
	}

	jc = jwtConfig{
		Secret: file.Section("jwt").Key("secret").String(),
	}
}
