// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/tiagomelo/docker-mongodb-custom-user-pass/config"
	"github.com/tiagomelo/docker-mongodb-custom-user-pass/db"
)

func run() error {
	ctx := context.Background()
	config, err := config.ReadConfig()
	if err != nil {
		return errors.Wrap(err, "reading config")
	}
	_, err = db.ConnectToMongoDb(ctx,
		config.MongoDbUser,
		config.MongoDbPassword,
		config.MongoDbHostName,
		config.MongoDbDatabase,
		config.MongoDbPort,
	)
	if err != nil {
		return errors.Wrap(err, "connecting to MongoDB")
	}
	fmt.Println("successfully connected to MongoDB.")
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
