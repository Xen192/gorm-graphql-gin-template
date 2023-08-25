package service

import (
	"context"
	"errors"
	"fmt"
	"mygpt/pkg/lib/encrypt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type FileUploadRepoInterface interface {
	Put(ctx context.Context, file_name string, file_parent string, data []byte) (string, error)
	Get(ctx context.Context, file_id string) (*string, []byte, string, string, error)
	Delete(ctx context.Context, file_id string, file_parent string) error
}

type fileUploadService struct {
	repository     FileUploadRepoInterface
	contextTimeout time.Duration
}

func NewFileUploadService(uploadRepo FileUploadRepoInterface, timeout time.Duration) *fileUploadService {
	return &fileUploadService{
		repository:     uploadRepo,
		contextTimeout: timeout,
	}
}

func (s *fileUploadService) Put(c context.Context, file_name string, file_parent string, data []byte) (string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if len(data) > 1048576 {
		return "", errors.New("File Size Exceeds 1MB")
	}

	file_id, err := s.repository.Put(ctx, file_name, file_parent, data)
	if err != nil {
		return "", err
	}

	return file_id, nil
}

func (s *fileUploadService) GenerateURL(c context.Context, file_id string, file_parent string) (string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	_, _, _, parent, err := s.repository.Get(ctx, file_id)
	if err != nil {
		return "", err
	}
	if parent != file_parent {
		return "", errors.New("Invalid File ParentId")
	}

	exp := time.Now().Add(time.Minute * 16).Unix()
	encTxt, err := encrypt.Encrypt(fmt.Sprintf("%s@%d", file_id, exp))
	if err != nil {
		logrus.Error(err)
		return "", errors.New("Encryption Failed")
	}
	return encTxt, nil
}

func (s *fileUploadService) Get(c context.Context, file_id string) (*string, []byte, string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	decTxt, err := encrypt.Decrypt(file_id)
	if err != nil {
		logrus.Error(err)
		return nil, []byte{}, "", errors.New("Decryption Failed")
	}
	indx := strings.Index(decTxt, "@")
	if indx == -1 {
		return nil, []byte{}, "", errors.New("Invalid Param Signature")
	}
	exp, err := strconv.ParseInt(decTxt[indx+1:], 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil, []byte{}, "", errors.New("Invalid Param Constraints")
	}
	if time.Now().Unix() > exp {
		return nil, []byte{}, "", errors.New("Access Expired")
	}

	name, data, mime, _, err := s.repository.Get(ctx, decTxt[:indx])
	return name, data, mime, err
}

func (s *fileUploadService) Delete(c context.Context, file_id string, file_parent string) error {
	return s.repository.Delete(c, file_id, file_parent)
}
