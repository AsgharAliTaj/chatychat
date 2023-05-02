package chatychat

import (
	"time"
	"gorm.io/gorm"
)

type Thread struct {
  gorm.Model
  ID int `gorm:"primaryKey" json:"id"`
  Title string `gorm:"unique" json:"title"`
  Description string `json:"description"`
  Post []Post  `json:"posts"`
  CreatedAt time.Time `json:"createdAt"`
}

type Post struct {
	gorm.Model
  ID int `gorm:"primaryKey" json:"id"`
  ThreadID int `json:"threadId"`
  Title string `json:"title"`
  Content string `json:"content"`
  Comment []Comment `json:"comments"`
  Votes int `gorm:"default:0" json:"votes"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}

type Comment struct {
	gorm.Model
  ID int `json:"id"`
  PostID int `json:"postId"`
  Content string `json:"content"`
  Votes int `gorm:"default:0" json:"votes"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}

type ThreadStore interface {
	GetThread(id int)
	GetThreads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(i int) error
}

type PostStore interface {
	GetPost(id int) 
	GetPosts() ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id int) error
}

type CommentStore interface {
	GetComment(id int)
	GetComments() ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(c int) error
}

type Store struct {
	ThreadStore
	PostStore
	CommentStore
}
