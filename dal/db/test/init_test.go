package test

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"testing"
)

func TestInitDb(t *testing.T) {
	database := db.GetRdbms()
	if database == nil {
		fmt.Println("db is nil")
		t.Fail()
	}
	err := db.CreateUser(context.TODO(), []*db.User{
		&db.User{
			UserName: "Jeff",
			Password: "123456",
		},
	})
	if err != nil {
		t.Fail()
	}
}
