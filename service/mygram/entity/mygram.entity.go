package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Username  string    `gorm:"unique;column:username" json:"username"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Age       int       `gorm:"column:age" json:"age"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "mygram.user"
}

type SocialMedia struct {
	ID             uint      `gorm:"primaryKey;column:id" json:"id"`
	Name           string    `gorm:"column:name" json:"name"`
	SocialMediaURL string    `gorm:"column:social_media_url" json:"social_media_url"`
	UserID         int64     `gorm:"column:user_id; not null" json:"user_id"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (SocialMedia) TableName() string {
	return "mygram.social_media"
}

type Photo struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	Caption   string    `gorm:"column:caption" json:"caption"`
	PhotoURL  string    `gorm:"column:photo_url" json:"photo_url"`
	UserID    int64     `gorm:"column:user_id; not null" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Photo) TableName() string {
	return "mygram.photo"
}

type Comment struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID    int64     `gorm:"column:user_id; not null" json:"user_id"`
	PhotoID   int       `gorm:"column:photo_id; not null" json:"photo_id"`
	Message   string    `gorm:"column:message" json:"message"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Comment) TableName() string {
	return "mygram.comment"
}
