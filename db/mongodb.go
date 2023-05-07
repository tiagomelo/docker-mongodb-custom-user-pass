// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package db

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	database string
	client   *mongo.Client
}

// For ease of unit testing.
var (
	connect = func(ctx context.Context, client *mongo.Client) error {
		return client.Connect(ctx)
	}
	ping = func(ctx context.Context, client *mongo.Client) error {
		return client.Ping(ctx, nil)
	}
)

// ConnectToMongoDb connects to a running MongoDB instance.
func ConnectToMongoDb(ctx context.Context, user, pass, host, database string, port int) (*MongoDb, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		uri(user, pass, host, database, port),
	))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create MongoDB client")
	}
	err = connect(ctx, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to MongoDB server")
	}
	err = ping(ctx, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping MongoDB server")
	}
	return &MongoDb{
		database: database,
		client:   client,
	}, nil
}

// uri generates uri string for connecting to MongoDB.
func uri(user, pass, host, database string, port int) string {
	const format = "mongodb://%s:%s@%s:%d/%s"
	bla := fmt.Sprintf(format, user, pass, host, port, database)
	return bla
}
