package controller

import (
	"context"
	"fmt"
	"io"
	"mygpt/pkg/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FileUploadServiceInterface interface {
	Put(c context.Context, file_name string, file_parent string, data []byte) (string, error)
	GenerateURL(c context.Context, file_id string, file_parent string) (string, error)
	Get(c context.Context, file_id string) (*string, []byte, string, error)
	Delete(c context.Context, file_id string, file_parent string) error
}

type FileUploadController struct {
	Service FileUploadServiceInterface
}

func (c *FileUploadController) Put(g *gin.Context) {
	parent := g.Param("parent_id")
	file, err := g.FormFile("file")
	if err != nil {
		g.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	fl, err := file.Open()
	if err != nil {
		g.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	fileContents, err := io.ReadAll(fl)
	if err != nil {
		g.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	fileId, err := c.Service.Put(g.Request.Context(), file.Filename, parent, fileContents)
	if err != nil {
		g.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	g.String(http.StatusOK, fileId)
}

func (c *FileUploadController) GenerateURL(g *gin.Context) {
	fileId := g.Param("file_id")
	parentId := g.Query("parent_id")

	fileId, err := c.Service.GenerateURL(g.Request.Context(), fileId, parentId)
	if err != nil {
		g.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	scheme := "http"
	if g.Request.TLS != nil {
		scheme = "https"
	}

	g.String(http.StatusOK, scheme+"://"+g.Request.Host+g.Request.URL.Path[:strings.Index(g.Request.URL.Path, "file")+4]+"/"+fileId)
}

func (c *FileUploadController) Get(g *gin.Context) {
	fileId := g.Param("file_id")

	name, fileContent, mime, err := c.Service.Get(g.Request.Context(), fileId)
	if err != nil {
		g.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	if name != nil && !utils.Contains(utils.DisplayableMime, mime, true) {
		g.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", *name))
	}
	g.Writer.Header().Set("Content-Type", mime)
	g.Writer.Header().Set("Content-Length", strconv.Itoa(len(fileContent)))
	g.Writer.Write(fileContent)
}

func (c *FileUploadController) Delete(g *gin.Context) {
	fileId := g.Param("file_id")
	parentId := g.Query("parent_id")

	err := c.Service.Delete(g.Request.Context(), fileId, parentId)
	if err != nil {
		g.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	g.String(http.StatusOK, "Deleted")
}
