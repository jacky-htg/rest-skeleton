package tests

import (
	"context"
	"database/sql"
	"log"
	"os"
	"rest/libraries/api"
	"rest/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
)

func TestUser(t *testing.T) {
	db, teardown := NewUnit(t)
	defer teardown()

	log := log.New(os.Stdout, "rest-skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	u := User{Db: db, Log: log}
	//t.Run("CRUD", u.Crud)
	t.Run("List", u.List)
}

// User struct for test users
type User struct {
	Db  *sql.DB
	Log *log.Logger
}

//Crud : unit test  for create get and delete user function
func (u *User) Crud(t *testing.T) {
	ctx := context.Background()
	u0 := models.User{
		Username: "Aladin",
		Email:    "aladin@gmail.com",
		Password: "1234",
		IsActive: false,
	}

	u0.Db = u.Db
	u0.Log = u.Log
	err := u0.Create(ctx)
	if err != nil {
		t.Fatalf("creating user u0: %s", err)
	}

	u1 := models.User{
		ID: u0.ID,
	}

	u1.Db = u.Db
	u1.Log = u.Log
	err = u1.Get(ctx)
	if err != nil {
		t.Fatalf("getting user u1: %s", err)
	}

	if diff := cmp.Diff(u1, u0); diff != "" {
		t.Fatalf("fetched != created:\n%s", diff)
	}

	u1.IsActive = false
	err = u1.Update(ctx)
	if err != nil {
		t.Fatalf("update user u1: %s", err)
	}

	u2 := models.User{
		ID: u1.ID,
	}

	u2.Db = u.Db
	u2.Log = u.Log
	err = u2.Get(ctx)
	if err != nil {
		t.Fatalf("getting user u2: %s", err)
	}

	if diff := cmp.Diff(u1, u2); diff != "" {
		t.Fatalf("fetched != updated:\n%s", diff)
	}

	err = u2.Delete(ctx)
	if err != nil {
		t.Fatalf("delete user u2: %s", err)
	}

	u3 := models.User{
		ID: u2.ID,
	}

	u3.Db = u.Db
	u3.Log = u.Log
	err = u3.Get(ctx)

	apiErr, ok := err.(*api.Error)
	if !ok || apiErr.Err != sql.ErrNoRows {
		t.Fatalf("getting user u3: %s", err)
	}
}

//List : unit test for user list function
func (u *User) List(t *testing.T) {
	ctx := context.Background()
	u0 := models.User{
		Username: "Aladin",
		Email:    "aladin@gmail.com",
		Password: "1234",
		IsActive: false,
	}

	u0.Db = u.Db
	u0.Log = u.Log

	err := u0.Create(ctx)
	if err != nil {
		t.Fatalf("creating user u0: %s", err)
	}

	var user models.User
	user.Db = u.Db
	user.Log = u.Log
	users, err := user.List(ctx)
	if err != nil {
		t.Fatalf("listing users: %s", err)
	}
	if exp, got := 1, len(users); exp != got {
		t.Fatalf("expected users list size %v, got %v", exp, got)
	}
}
