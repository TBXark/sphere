package storage

import (
	"context"
	"io"
)

type URLHandler interface {
	GenerateURL(key string) string
	GenerateURLs(keys []string) []string
	ExtractKeyFromURL(uri string) string
	ExtractKeyFromURLWithMode(uri string, strict bool) (string, error)
}

type ImageURLHandler interface {
	URLHandler
	GenerateImageURL(key string, width int) string
}

type TokenGenerator interface {
	// GenerateUploadToken [token, key, url]
	GenerateUploadToken(ctx context.Context, fileName string, dir string, nameBuilder func(filename string, dir ...string) string) ([3]string, error)
}

type FileUploader interface {
	UploadFile(ctx context.Context, file io.Reader, key string) (string, error)
	UploadLocalFile(ctx context.Context, file string, key string) (string, error)
}

type FileDownloader interface {
	IsFileExists(ctx context.Context, key string) (bool, error)
	DownloadFile(ctx context.Context, key string) (io.ReadCloser, string, int64, error) // reader, mime, size
}

type FileDeleter interface {
	DeleteFile(ctx context.Context, key string) error
}

type FileMoverCopier interface {
	MoveFile(ctx context.Context, sourceKey string, destinationKey string, overwrite bool) error
	CopyFile(ctx context.Context, sourceKey string, destinationKey string, overwrite bool) error
}

type Storage interface {
	URLHandler
	FileDeleter
	FileUploader
	FileDownloader
	FileMoverCopier
}

type CDNStorage interface {
	Storage
	TokenGenerator
}

type ImageStorage interface {
	ImageURLHandler
	CDNStorage
}
