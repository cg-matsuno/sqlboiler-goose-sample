package main

import (
	"context"
	"database/sql"
	"time"

	dbmodels "github.com/cg-matsuno/sqlboiler-goose-sample/db/models"
	"github.com/k0kubun/pp"
	"github.com/volatiletech/sqlboiler/boil"
)

func init() {
	postgresURL := "host=localhost port=5432 user=matsuno dbname=testdb sslmode=disable password=matsuno"
	_db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		panic(err)
	}
	boil.SetDB(_db)
	boil.SetLocation(time.Local)
	return nil
}

func main() {
	var err error
	ctx := context.Background()

	if err = insertUsers(ctx); err != nil {
		panic(err)
	}

	if err = selectUsers(ctx); err != nil {
		panic(err)
	}

	if err = selectOneUser(ctx); err != nil {
		panic(err)
	}
}

func insertUser(ctx) (err error) {
	user := dbmodels.User{}
	if users, err = dbmodels.Users().One(ctx, tx); err != nil {
		return err
	}
	pp.Println(users)

	var users dbmodels.UserSlice
	if users, err = dbmodels.Users(
		dbmodels.UserWhere.ID.EQ(1),
	).AllG(ctx, tx); err != nil {
		return err
	}
	pp.Println(users)
}

func selectUsers(ctx) (err error) {
	var users dbmodels.UserSlice
	if users, err = dbmodels.Users().AllG(ctx, tx); err != nil {
		return err
	}
	pp.Println(users)

	var users dbmodels.UserSlice
	if users, err = dbmodels.Users(
		dbmodels.UserWhere.ID.EQ(1),
	).AllG(ctx, tx); err != nil {
		return err
	}
	pp.Println(users)
}

func selectOneUser(ctx) (err error) {
	var users *dbmodels.User
	if users, err = dbmodels.Users().One(ctx, tx); err != nil {
		return err
	}
	pp.Println(users)

	var users dbmodels.UserSlice
	if users, err = dbmodels.Users(
		dbmodels.UserWhere.ID.EQ(1),
	).AllG(ctx, tx); err != nil {
		return err
	}
	pp.Println(users)
}
