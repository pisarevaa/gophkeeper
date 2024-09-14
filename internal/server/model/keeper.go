package model

import (
	"time"
)

type Keeper struct {
	ID        int64
	Name      string
	Data      string
	Type      DataTypeEnum
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddTextData struct {
	Name string `json:"name" validate:"required,lt=0,lte=250"`
	Data string `json:"data" validate:"required,lt=0"`
}

type AddKeeper struct {
	Name string
	Data string
	Type DataTypeEnum
}

type DataResponse struct {
	ID        int64        `json:"id"`
	Nmae      string       `json:"name"`
	Data      string       `json:"data"`
	Type      DataTypeEnum `json:"type"`
	CreatedAt DateTime     `json:"createdAt"`
	UpdatedAt DateTime     `json:"updatedAt"`
}
