package config

type DBConfig struct {
	Dialect		string
	Host		string
	Port		int
	Username	string
	Password	string
	Name		string
	Charset		string
}

func GetConfig() *DBConfig {
	return &DBConfig{
		Dialect:  "mysql",
		Host: 	  "127.0.0.1",
		Port: 	  3306,
		Username: "root",
		Password: "urban1004",
		Name: 	  "todolist",
		Charset:  "utf8",
	}
}