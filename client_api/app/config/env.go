package config

import (
	"client_api/utils"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
)

func newEnv() EnvVariables {
	return EnvVariables{
		loadEnvFile: loadEnvFileI,
		loadValues:  loadValuesI,
	}
}

type EnvVariables struct {
	GolangEnv                  string
	Port                       int
	JwtVerificationKey         string
	JwtSigningKey              string
	DatabaseUrl                string
	PostgresHostname           string
	PostgresPort               int
	PostgresDbName             string
	PostgresUsername           string
	PostgresPassword           string
	PostgresSslMode            string
	PostgresMaxIdleConnections int
	PostgresMaxConnections     int
	AuthKey                    string
	MailgunApiKey              string
	MailgunValidationKey       string
	MailgunSigningKey          string
	Domain                     string
	MailgunSendingKey          string
	SlackBotToken              string
	SlackErrorChannelId        string
	SlackInfoChannelId         string
	loadEnvFile                func(filename string) error
	loadValues                 func(*EnvVariables)
}

func (e *EnvVariables) load() (err error) {
	switch os.Getenv("GOLANG_ENV") {
	case "test":
		err = e.loadEnvFile(".env.test")
	case "development":
		err = e.loadEnvFile(".env")
	case "production":

	default:
		err = e.loadEnvFile(".env")
	}

	//err = e.loadEnvFile(".env")
	e.loadValues(e)

	return
}

var loadEnvFileI = func(filename string) (err error) {
	if filename == "" {
		return fmt.Errorf("empty filename, specify the file that holds the environment variables")
	}

	envPath := path.Join(
		utils.RootPath(),
		filename,
	)

	err = godotenv.Load(envPath)
	return
}

var loadValuesI = func(e *EnvVariables) {
	values := reflect.ValueOf(e).Elem()
	fields := reflect.TypeOf(e).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		snakeCaseName := strings.ToUpper(utils.ToSnakeCase(field.Name))

		envValue := os.Getenv(field.Name)
		if envValue == "" {
			envValue = os.Getenv(snakeCaseName)
		}

		switch value.Kind() {
		case reflect.Int:
			v, _ := strconv.ParseInt(envValue, 10, 0)
			value.SetInt(v)
		case reflect.String:
			value.SetString(envValue)
		}
	}
}
