package config

type Config struct {
	App struct {
		Port string `json:"port"`
		Env  string `json:"env"`
	}
	Mysql struct {
		DataSource string `json:"data_source"`
	}
}
