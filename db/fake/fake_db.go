package fakedb

import (
	"errors"
	"fmt"
	postmodel "go-module/modules/post/model"
	"os"
	"time"
)

type FakeDB struct {
	Error error
}

func OpenDB(connectStr string) (*FakeDB, error) {
	return &FakeDB{}, nil
}

func (db *FakeDB) Where(query interface{}, args ...interface{}) *FakeDB {
	return db
}

func (db *FakeDB) First(dest interface{}, conds ...interface{}) *FakeDB {
	if os.Getenv("PANIC_ERROR_DB") == "true" {
		panic(errors.New("can not connect to db"))
	}

	GetAndCopyValueFromDbToPost(dest)
	return db
}

func GetAndCopyValueFromDbToPost(dest interface{}) {
	post, ok := dest.(*postmodel.Post)
	fmt.Println(post)
	if ok {
		post.Id = "123"
		post.Title = "Catch exception in api"
		post.Content = "How to catch global exception in golang rest api"
		post.Published = false
		post.PublishedAt = nil
		post.CreateAt = time.Now()
		post.UpdateAt = time.Now()
		post.Deleted = false
	}
}
