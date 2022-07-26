package report

import (
	"errors"
	"fmt"
	"time"

	"github.com/vincentconace/app-test/internal/domain"
)

type Service interface {
	CreateReportPosts(userName string, idPost int, text string) (domain.Report, error)
	CreateReportComments(userName string, idComment int, text string) (domain.Report, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) CreateReportPosts(userName string, idPost int, text string) (domain.Report, error) {
	exist := s.repository.Exists(userName)
	if !exist {
		return domain.Report{}, errors.New(fmt.Sprintf("user_name %s not found", userName))
	}
	existPosts := s.repository.ExistsPosts(idPost)
	if !existPosts {
		return domain.Report{}, errors.New(fmt.Sprintf("post with id %d not found", idPost))
	}
	date := time.Now()
	id, err := s.repository.SaveReportPosts(userName, idPost, text, date)
	if err != nil {
		return domain.Report{}, errors.New("not create report")
	}

	report := domain.Report{
		ID:         id,
		IDPosts:    idPost,
		UserName:   userName,
		Text:       text,
		CreateDate: date,
	}

	return report, nil
}

func (s *service) CreateReportComments(userName string, idComment int, text string) (domain.Report, error) {
	exist := s.repository.Exists(userName)
	if !exist {
		return domain.Report{}, errors.New(fmt.Sprintf("user_name with id %s not found", userName))
	}
	existComment := s.repository.ExistsComment(idComment)
	if !existComment {
		return domain.Report{}, errors.New(fmt.Sprintf("post with id %d not found", idComment))
	}
	date := time.Now()
	id, err := s.repository.SaveReportPosts(userName, idComment, text, date)
	if err != nil {
		return domain.Report{}, errors.New("not create report")
	}

	report := domain.Report{
		ID:         id,
		IDComments: idComment,
		UserName:   userName,
		Text:       text,
		CreateDate: date,
	}

	return report, nil
}
