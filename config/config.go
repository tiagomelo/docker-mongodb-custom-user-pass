// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config holds all configuration needed by this app.
type Config struct {
	MongoDbUser     string `envconfig:"MONGODB_USER"`
	MongoDbPassword string `envconfig:"MONGODB_PASSWORD"`
	MongoDbDatabase string `envconfig:"MONGODB_DATABASE"`
	MongoDbHostName string `envconfig:"MONGODB_HOST_NAME"`
	MongoDbPort     int    `envconfig:"MONGODB_PORT"`
}

var (
	godotenvLoad     = godotenv.Load
	envconfigProcess = envconfig.Process
)

func ReadConfig() (*Config, error) {
	if err := godotenvLoad(); err != nil {
		return nil, errors.Wrap(err, "loading env vars")
	}
	config := new(Config)
	if err := envconfigProcess("", config); err != nil {
		return nil, errors.Wrap(err, "processing env vars")
	}
	return config, nil
}
