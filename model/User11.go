package model

type Roles []string

type User11 struct {
	Name        []byte                 `gorm:"serializer:json"`
	Roles       Roles                  `gorm:"serializer:json"`
	Contracts   map[string]interface{} `gorm:"serializer:json"`
	JobInfo     Job                    `gorm:"type:bytes;serializer:gob"`
	CreatedTime int64                  `gorm:"serializer:unixtime;type:time"` // store int as datetime into database
}
