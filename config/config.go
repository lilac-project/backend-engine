package config

type PostgreConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

type MongoConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

type Configs struct {
	Mongo   MongoConfig   `json:"mongo"`
	Postgre PostgreConfig `json:"postgre"`
}
