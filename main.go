package main

import (
	"context"
	"database/sql"
	"time"

	dbmodels "github.com/cg-matsuno/sqlboiler-goose-sample/db/models"
	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func init() {
	postgresURL := "host=localhost port=5432 user=matsuno dbname=testdb sslmode=disable password=matsuno"
	_db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		panic(err)
	}
	boil.SetDB(_db)
	boil.SetLocation(time.Local)
}

func main() {
	var err error
	ctx := context.Background()

	if err = deleteAll(ctx); err != nil {
		panic(err)
	}
	if err = insertUser(ctx); err != nil {
		panic(err)
	}

	if err = selectUsers(ctx); err != nil {
		panic(err)
	}

	if err = selectOneUser(ctx); err != nil {
		panic(err)
	}
}

func insertUser(ctx context.Context) (err error) {
	var tx boil.ContextTransactor
	if tx, err = boil.BeginTx(ctx, nil); err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	org := dbmodels.Org{
		Name: "org name",
	}
	if err = org.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	user := dbmodels.User{
		OrgID: org.ID,
		UID:   null.StringFrom("uid"),
		Name:  "user name",
		Owner: true,
	}

	if err = user.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func selectUsers(ctx context.Context) (err error) {
	var users dbmodels.UserSlice
	if users, err = dbmodels.Users().AllG(ctx); err != nil {
		return err
	}
	for _, u := range users {
		pp.Println("User Name", u.Name)
	}

	var usersWithWhere dbmodels.UserSlice
	if usersWithWhere, err = dbmodels.Users(
		dbmodels.UserWhere.Name.EQ("user name"),
	).AllG(ctx); err != nil {
		return err
	}
	for _, u := range usersWithWhere {
		pp.Println("User Name", u.Name)
	}
	return nil
}

func selectOneUser(ctx context.Context) (err error) {
	var user *dbmodels.User
	if user, err = dbmodels.Users(
		dbmodels.UserWhere.Name.EQ("user name"),
		qm.Load(dbmodels.UserRels.Org),
	).OneG(ctx); err != nil {
		if err == sql.ErrNoRows {
			pp.Println("no user found")
			return nil
		}
		return err
	}
	pp.Println("User Name", user.Name)
	pp.Println("Org Name", user.R.Org.Name)
	return nil
}

func deleteAll(ctx context.Context) (err error) {
	var tx boil.ContextTransactor
	if tx, err = boil.BeginTx(ctx, nil); err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var orgs dbmodels.OrgSlice
	if orgs, err = dbmodels.Orgs(
		qm.Load(dbmodels.OrgRels.Users),
	).AllG(ctx); err != nil {
		return err
	}

	for _, org := range orgs {
		if _, err = org.R.Users.DeleteAll(ctx, tx); err != nil {
			return err
		}

		if _, err = org.Delete(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}
