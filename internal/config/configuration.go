package config

type DBConfiguration struct {
	DbHost              string
	DbPort              int
	DbName              string
	DbUser              string
	DbPass              string
	DbLogEnable         bool
	DbMaxConnection     int
	DbMaxIdleConnection int
}

type Configuration struct {
	DBConfiguration
}

func Config() Configuration {
	return Configuration{
		DBConfiguration: DBConfiguration{
			DbHost: "localhost",
			DbPort: 25433,
			DbName: "app_db",
			DbUser: "app_user",
			DbPass: "app_pass",
			DbLogEnable: true,
			DbMaxConnection: 100,
		},
	}
}

