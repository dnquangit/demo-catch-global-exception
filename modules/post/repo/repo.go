package postrepo

import (
	fakedb "go-module/db/fake"
	postmodel "go-module/modules/post/model"
)

type PostRepo interface {
	FindPostById(string) (*postmodel.Post, error)
}

type postRepo struct {
	db *fakedb.FakeDB
}

func NewPostRepo(db *fakedb.FakeDB) *postRepo {
	return &postRepo{db: db}
}

func (repo *postRepo) FindPostById(id string) (*postmodel.Post, error) {
	var post postmodel.Post
	if err := repo.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
