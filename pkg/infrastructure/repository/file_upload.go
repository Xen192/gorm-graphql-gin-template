package repository

import (
	"context"
	"errors"
	"mygpt/model"
	"mygpt/pkg/utils"
	"mygpt/query"

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
	entry := model.File{
		ID:       ulid.Make().String(),
		ParentID: &file_parent,
		FileName: &file_name,
		MimeType: mimetype.Detect(data).String(),
		Data:     data,
	}
	res := query.File.WithContext(ctx).Create(&entry)
	if res != nil {
		logrus.Error(res)
		return "", errors.New("Failed Inserting The File Data")
	}
	return entry.ID, nil
}

func (r *fileUploadRepository) Get(ctx context.Context, file_id string) (*string, []byte, string, string, error) {
	entry, res := query.File.WithContext(ctx).Where(query.File.ID.Eq(file_id)).First()
	if res != nil {
		logrus.Error(res)
		return nil, []byte{}, "", "", errors.New("Failed Fetching Data")
	}
	return entry.FileName, entry.Data, entry.MimeType, utils.SafeDereference(entry.ParentID), nil
}

func (r *fileUploadRepository) Delete(ctx context.Context, file_id string, file_parent string) error {
	q := query.File
	_, res := q.WithContext(ctx).Where(q.ID.Eq(file_id), q.ParentID.Eq(file_parent)).First()
	if res != nil {
		logrus.Error(res)
		return errors.New("Failed Fetching Data")
	}
	_, res = q.WithContext(ctx).Where(q.ID.Eq(file_id), q.ParentID.Eq(file_parent)).Delete()
	if res != nil {
		logrus.Error(res)
		return errors.New("Failed Deleting Data")
	}
	return nil
}
