// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package db

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnectToMongoDb(t *testing.T) {
	testCases := []struct {
		name          string
		mockConnect   func(ctx context.Context, client *mongo.Client) error
		mockPing      func(ctx context.Context, client *mongo.Client) error
		expectedError error
	}{
		{
			name: "happy path",
			mockConnect: func(ctx context.Context, client *mongo.Client) error {
				return nil
			},
			mockPing: func(ctx context.Context, client *mongo.Client) error {
				return nil
			},
		},
		{
			name: "error calling connect",
			mockConnect: func(ctx context.Context, client *mongo.Client) error {
				return errors.New("random error")
			},
			expectedError: errors.New("failed to connect to MongoDB server: random error"),
		},
		{
			name: "error calling ping",
			mockConnect: func(ctx context.Context, client *mongo.Client) error {
				return nil
			},
			mockPing: func(ctx context.Context, client *mongo.Client) error {
				return errors.New("random error")
			},
			expectedError: errors.New("failed to ping MongoDB server: random error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			connect = tc.mockConnect
			ping = tc.mockPing
			client, err := ConnectToMongoDb(context.TODO(), "user", "pass", "host", "db", 11111)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no error, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
				require.Nil(t, client)
			} else if tc.expectedError != nil {
				t.Fatalf(`expected error "%v", got nil`, tc.expectedError)
				require.NotNil(t, client)
			}
		})
	}
}
