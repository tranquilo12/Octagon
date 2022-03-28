package main

import "gopkg.in/ini.v1"

type RedisConnectionDetails struct {
	Host     string
	Port     string
	Password string
}

//ReadConfig Read the config.ini file from the current path
func ReadConfig() *ini.File {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	return cfg
}

//GetPolygonioKey Get the polygonio key from the config.ini file
func GetPolygonioKey() string {
	cfg := ReadConfig()
	return cfg.Section("POLYGON").Key("loving_aryabhata_key").String()
}

//GetRedisConnectionDetails Get the redis connection details from the config.ini file
func GetRedisConnectionDetails() (RedisConnectionDetails, error) {
	cfg := ReadConfig()
	host := cfg.Section("REDIS").Key("host").String()
	port := cfg.Section("REDIS").Key("port").String()
	password := cfg.Section("REDIS").Key("password").String()
	return RedisConnectionDetails{
		Host:     host,
		Port:     port,
		Password: password,
	}, nil
}
