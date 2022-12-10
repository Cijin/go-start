package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	// as the config file is 'app.env'
	viper.SetConfigName(".app")

	/*
	 * When called, Viper will check for an environment variable any time a viper.
	 * Get request is made.
	 * It will check for an environment variable with a name matching the key
	 * uppercased and prefixed with the EnvPrefix if set.
	 */
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
