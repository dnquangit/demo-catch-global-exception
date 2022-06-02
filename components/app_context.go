package components

import fakedb "go-module/db/fake"

type AppContext struct {
	FakeDB *fakedb.FakeDB
	//GormDB      *gorm.DB
}

func NewAppContext(db *fakedb.FakeDB) *AppContext {
	return &AppContext{FakeDB: db}
}
