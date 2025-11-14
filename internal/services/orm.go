package services

type Task struct {
	Task string `json:"task"`
	ID   string `gorm:"primaryKey" json:"id"`
}
