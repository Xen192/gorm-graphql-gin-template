package repository

import (
	"context"
	"errors"
	"mygpt/models"
	"mygpt/pkg/utils"

	"github.com/gabriel-vasile/mimetype"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type fileUploadRepository struct {
	db *gorm.DB
}

func NewFileUploadRepository(db *gorm.DB) *fileUploadRepository {
	return &fileUploadRepository{
		db: db,
	}
}

func (r *fileUploadRepository) Put(ctx context.Context, file_name string, file_parent string, data []byte) (string, error) {
	entry := models.File{
		ID:       ulid.Make().String(),
		ParentId: &file_parent,
		FileName: &file_name,
		MimeType: mimetype.Detect(data).String(),
		Data:     data,
	}
	res := r.db.Create(&entry)
	if res.Error != nil {
		logrus.Error(res.Error)
		return "", errors.New("Failed Inserting The File Data")
	}
	return entry.ID, nil
}

func (r *fileUploadRepository) Get(ctx context.Context, file_id string) (*string, []byte, string, string, error) {
	entry := models.File{ID: file_id}
	res := r.db.First(&entry)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil, []byte{}, "", "", errors.New("Failed Fetching Data")
	}
	return entry.FileName, entry.Data, entry.MimeType, utils.SafeDereference(entry.ParentId), nil
}

func (r *fileUploadRepository) Delete(ctx context.Context, file_id string, file_parent string) error {
	res := r.db.Where("id = ? and parent_id = ?", file_id, file_parent).First(&models.File{})
	if res.Error != nil {
		logrus.Error(res.Error)
		return errors.New("Failed Fetching Data")
	}
	res = r.db.Delete(&models.File{}, []string{file_id})
	if res.Error != nil {
		logrus.Error(res.Error)
		return errors.New("Failed Deleting Data")
	}
	return nil
}
