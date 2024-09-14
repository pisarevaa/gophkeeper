package model

import (
	"time"
)

type Keeper struct {
	ID        int64
	Data      string
	Type      DataTypeEnum
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddDataText struct {
	Data string `json:"data" validate:"required,data"`
}

type AddKeeper struct {
	Data string
	Type DataTypeEnum
}

type DataResponse struct {
	ID        int64          `json:"id"`
	Data      string         `json:"data"`
	Type      DataTypeEnum   `json:"type"`
	CreatedAt DateTime `json:"createdAt"`
	UpdatedAt DateTime `json:"updatedAt"`
}
