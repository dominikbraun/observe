// Package main provides the observe application.
package main

import "github.com/spf13/viper"

// settingsFilename is the name of the file that will be used to
// parse the observe settings (without extension).
const settingsFilename string = "settings"

// Settings represents a collection of settings required for an
// observation. At the moment, these primarily are mail settings.
type Settings struct {
	Mail struct {
		From string
		To   string
	}
	Sendgrid struct {
		Key string
	}
}

// readSettings reads the settings file under a given path and
// unmarshalls the file content into a Settings instance.
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
