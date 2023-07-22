package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

var (
	// Global config
	confs = Config{}
	// lock  = sync.Mutex{}
)

// Config is grpc of configs we need for project
type Config struct {
	CORS     string   `yaml:"cors" `
	Service  Service  `yaml:"service" `
	Postgres Database `yaml:"database"`
	Jaeger   tracer   `yaml:"jaeger"`
}

func validate(c any) error {
	errmsg := ""
	numFields := reflect.TypeOf(c).NumField()
	for i := 0; i < numFields; i++ {
		fieldType := reflect.TypeOf(c).Field(i)
		tagval, ok := fieldType.Tag.Lookup("required")
		isRequired := ok && tagval == "true"
		if !isRequired {
			continue
		}
		fieldVal := reflect.ValueOf(c).Field(i)
		if fieldVal.Kind() == reflect.Struct {
			if err := validate(fieldVal.Interface()); err != nil {
				errmsg += fmt.Sprintf("%s > [%v], ", fieldType.Name, err)
			}
		} else {
			if fieldVal.IsZero() {
				errmsg += fmt.Sprintf("%s, ", fieldType.Name)
			}
		}
	}
	if errmsg == "" {
		return nil
	}
	return errors.New(errmsg)
}

func C() *Config {
	return &confs
}

// init configs
func InitConfig() {
	dir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.AddConfigPath(dir)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("config problem : " + err.Error())
	}
	loadConfigs()

}

func loadConfigs() {
	must(viper.Unmarshal(&confs),
		"could not unmarshal config file")
	must(validate(confs),
		"some required configurations are missing")
	log.Printf("configs loaded from file successfully \n")
}

func must(err error, logMsg string) {
	if err != nil {
		log.Println(logMsg)
		panic(err.Error())
	}
}
