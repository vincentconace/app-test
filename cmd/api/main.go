package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vincentconace/app-test/cmd/api/handler"
	"github.com/vincentconace/app-test/internal/report"
)

func main() {
	db, err := sql.Open("sqlite3", "./../../app-test.db")
	if err != nil {
		panic(err)
	}
	repository := report.NewRepository(db)
	service := report.NewService(repository)
	handlerReport := handler.NewHandlerReport(service)

	r := gin.Default()
	rg := r.Group("/api/v1")

	report := rg.Group("/report")
	{
		report.POST("/posts", handlerReport.CreateReportPosts())
		report.POST("/comments", handlerReport.CreateReportComments())
	}

	r.Run()
}
