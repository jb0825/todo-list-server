package config

import "fmt"

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetMysqlConfig() *DBConfig {
	return &DBConfig{
		Dialect:  "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "urban1004",
		Name:     "todolist",
		Charset:  "utf8",
	}
}
func GetMysqlDSN(config *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)
}

func GetPostgresConfig() *DBConfig {
	return &DBConfig{
		Dialect:  "postgres",
		Host:     "127.0.0.1",
		Port:     5432,
		Username: "postgres",
		Password: "urban1004",
		Name:     "todolist",
		Charset:  "utf8",
	}
}
func GetPostgresDSN(config *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul",
		config.Host,
		config.Username,
		config.Password,
		config.Name,
		config.Port,
	)
}
