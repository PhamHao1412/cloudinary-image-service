package entity

import "time"

type Image struct {
	ID         string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	URL        string    `json:"url"`
	Format     string    `json:"format"`
	UploadedAt time.Time `json:"uploaded_at"`
}
