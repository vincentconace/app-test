package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vincentconace/app-test/internal/domain"
	"github.com/vincentconace/app-test/internal/report"
)

var Reports []domain.Report

type RepositoryMock struct {
	db []domain.Report
}

func (r *RepositoryMock) SaveReportPosts(userName string, idPost int, text string, date time.Time) (int, error) {
	report := domain.Report{
		ID:         Reports[len(Reports)-1].ID + 1,
		IDPosts:    idPost,
		UserName:   userName,
		Text:       text,
		CreateDate: date,
	}

	r.db = append(r.db, report)

	return report.ID, nil
}

func (r *RepositoryMock) SaveReportComments(userName string, idComment int, text string, date time.Time) (int, error) {
	report := domain.Report{
		ID:         Reports[len(Reports)-1].ID + 1,
		IDComments: idComment,
		UserName:   userName,
		Text:       text,
		CreateDate: date,
	}

	r.db = append(r.db, report)

	return report.ID, nil
}

func (r *RepositoryMock) Exists(userName string) bool {
	return true
}

func (r *RepositoryMock) ExistsPosts(idPost int) bool {
	return true
}

func (r *RepositoryMock) ExistsComment(idPost int) bool {
	return true
}

func createServer() *gin.Engine {
	repository := RepositoryMock{Reports}
	service := report.NewService(&repository)
	handlerReport := NewHandlerReport(service)

	r := gin.Default()
	rg := r.Group("/api/v1")
	report := rg.Group("/reports")
	{
		report.POST("/posts", handlerReport.CreateReportPosts())
		report.POST("/comments", handlerReport.CreateReportComments())
	}

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestCreateReportsPosts(t *testing.T) {
	r := createServer()

	body := `{
		"id_comments": 1,
		"user_name": "conacevincent@gmail.com",
		"text": "This is a report test"
	}`

	req, res := createRequestTest(http.MethodPost, "/api/v1/reports/posts/", body)

	r.ServeHTTP(res, req)

	var response domain.Report

	json.Unmarshal(res.Body.Bytes(), &response)
	assert.Equal(t, http.StatusCreated, res.Code)
}
