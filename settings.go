package main

import "github.com/spf13/viper"

type Settings struct {
	Mail struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"mail"`
	Sendgrid struct {
		Key string `json:"key"`
	} `json:"sendgrid"`
}

func readSettings(path, filename string) (*Settings, error) {
	var settings Settings
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&settings); err != nil {
		return nil, err
	}

	return &settings, nil
}
