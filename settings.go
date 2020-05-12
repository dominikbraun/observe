package main

import "github.com/spf13/viper"

const settingsFilename string = "settings"

type Settings struct {
	Mail struct {
		From string
		To   string
	}
	Sendgrid struct {
		Key string
	}
}

func readSettings(path string) (*Settings, error) {
	var settings Settings
	viper.SetConfigName(settingsFilename)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&settings); err != nil {
		return nil, err
	}

	return &settings, nil
}
