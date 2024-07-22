package config

import (
	"reflect"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kanztu/OracleFreeTierInstanceCreator/validators"
	"github.com/spf13/viper"
)

type Config struct {
	CompartmentOCID    string `mapstructure:"COMPARTMENT_OCID" validate:"required"`
	AvailabilityDomain string `mapstructure:"AVAILABILITY_DOMAIN" validate:"required"`
	DisplayName        string `mapstructure:"DISPLAY_NAME" validate:"required"`
	ImageOCID          string `mapstructure:"IMAGE_OCID" validate:"required"`
	SubnetOCID         string `mapstructure:"SUBNET_OCID" validate:"required"`
	SshKey             string `mapstructure:"SSH_KEY" validate:"required"`
}

func New() (*Config, error) {
	envPath := "./env/.env"
	if testing.Testing() {
		envPath = "." + envPath
	}
	if err := godotenv.Load(envPath); err != nil {
		return nil, err
	}
	config := Config{}
	bindEnvs(config)
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, validate(&config)
}

func validate(c *Config) error {
	return validators.Validate(c)
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
