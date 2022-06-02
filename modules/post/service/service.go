package postService

import (
	"errors"
	postmodel "go-module/modules/post/model"
	postrepo "go-module/modules/post/repo"
	"os"
)

type PostService interface {
	GetPostById(string) (*postmodel.PostResponse, error)
}

type postService struct {
	postrepo postrepo.PostRepo
}

func NewPostService(repo postrepo.PostRepo) *postService {
	return &postService{postrepo: repo}
}

func (s *postService) GetPostById(id string) (*postmodel.PostResponse, error) {

	posts, err := s.postrepo.FindPostById(id)
	if err != nil {
		return nil, err
	}

	response := MappingPostResponse(posts)

	return response, nil
}

func MappingPostResponse(from *postmodel.Post) *postmodel.PostResponse {
	if from == nil {
		return nil
	}

	to := postmodel.PostResponse{
		Id:          from.Id,
		Title:       from.Title,
		Content:     from.Content,
		Published:   from.Published,
		PublishedAt: from.PublishedAt,
	}

	if os.Getenv("PANIC_ERROR_SERVICE") == "true" {
		panic(errors.New("cannot mapping"))
	}

	return &to
}
