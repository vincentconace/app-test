package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentconace/app-test/internal/report"
)

type HandlerReport struct {
	service report.Service
}

func NewHandlerReport(service report.Service) *HandlerReport {
	return &HandlerReport{service}
}

type request struct {
	IDPosts    int    `json:"id_posts"`
	IDComments int    `json:"id_comments"`
	UserName   string `json:"user_name"`
	Text       string `json:"text"`
}

func (h *HandlerReport) CreateReportPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		fmt.Println(req.UserName)
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "can not bind"})
			return
		}

		report, err := h.service.CreateReportPosts(req.UserName, req.IDPosts, req.Text)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, report)
	}
}

func (h *HandlerReport) CreateReportComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request

		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "can not bind"})
			return
		}

		report, err := h.service.CreateReportComments(req.UserName, req.IDComments, req.Text)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, report)
	}
}
