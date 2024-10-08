package model

import (
	"io"
	"time"
)

type Keeper struct {
	ID        int64
	Name      string
	Data      string
	ObjectID  string
	FileName  string
	Type      DataTypeEnum
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddTextData struct {
	Name string `json:"name" validate:"required,gt=0,lte=250"`
	Data string `json:"data" validate:"required,gt=0"`
}

type AddBinarytData struct {
	Name string
	File []byte
}

type AddKeeper struct {
	Name     string
	Data     string
	ObjectID string
	FileName string
	Type     DataTypeEnum
}

type DataResponseShort struct {
	ID   int64        `json:"id"`
	Name string       `json:"name"`
	Type DataTypeEnum `json:"type"`
}

type DataResponse struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Data      string       `json:"data,omitempty"`
	FileName  string       `json:"filename,omitempty"`
	ObjectID  string       `json:"object_id,omitempty"`
	Type      DataTypeEnum `json:"type"`
	CreatedAt DateTime     `json:"createdAt"`
	UpdatedAt DateTime     `json:"updatedAt"`
}

type UploadedFile struct {
	Size        int64
	ContentType string
	File        io.Reader
	FileName    string
	FileContent string
	Data        []byte
}

type MinioOperationError struct {
	ObjectID string
	Error    error
}
