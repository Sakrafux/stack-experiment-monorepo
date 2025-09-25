package config

type Config struct {
	Port               string
	DbConnectionString string
	JWT                struct {
		AccessSecret  string
		RefreshSecret string
	}
}
