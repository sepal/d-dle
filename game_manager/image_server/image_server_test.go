package main

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sepal/dodle/game_manager/dodle"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

func tearDown(db *bun.DB) error {
	ctx := context.Background()

	db.NewDropTable().Model((*dodle.ImageEntry)(nil)).Exec(ctx)
	db.NewDropTable().Model((*dodle.DBRound)(nil)).Exec(ctx)

	return nil
}

func TestHandleRequest(t *testing.T) {
	db := initDB()
	defer tearDown(db)

	db.RegisterModel((*dodle.DBRound)(nil))
	db.RegisterModel((*dodle.ImageEntry)(nil))

	ctx := context.Background()

	fixture := dbfixture.New(db)
	if err := fixture.Load(ctx, os.DirFS("fixtures"), "rounds.yml"); err != nil {
		t.Fatalf("Error while trying to load fixtures: %s", err)
	}

	request := events.APIGatewayProxyRequest{
		StageVariables: map[string]string{
			"env": "testing",
		},
		QueryStringParameters: map[string]string{
			"time":  "1653976200",
			"level": "0",
		},
	}

	response, err := HandleRequest(request)

	if err != nil {
		t.Fatal(err)
	}

	if len(response.Body) <= 0 {
		t.Fatal("Retrieved empty image")
	}
}
