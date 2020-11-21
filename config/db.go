package config

type Config struct{
	AppName string `default:App`
	Port string `default:8080`
}