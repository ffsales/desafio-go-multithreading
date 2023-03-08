package configs

import "github.com/spf13/viper"

type Conf struct {
	UrlApiCep string `mapstructure:"URL_APICEP"`
	UrlViaCep string `mapstructure:"URL_VIACEP"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	//essa função faz as variáveis de ambiente sobrescresverem o as configs do arquivo .env
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
