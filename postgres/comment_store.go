package postgres

import (
	"github.com/asgharalitaj/chatychat"
	"gorm.io/gorm"
)

func NewCommentStore() *CommentStore{
  return &CommentStore {
    db: GETDB(),
  }
}

type CommentStore struct {
  db *gorm.DB
}


func (p *CommentStore) GetComment(id int) (comment chatychat.Comment, err error) {
	result := p.db.First(&comment, id)
	if result.Error != nil {
		err = result.Error
		return 
	}
	return comment, nil
}

func (p *CommentStore) GetComments() (comment []chatychat.Comment, err error) {
	result := p.db.Find(&comment)
	if result.Error != nil {
		return
	}
	return comment, nil
}

func (p *CommentStore) CreateComment(ps chatychat.Comment) (err error) {
	result := p.db.Create(&ps)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (p *CommentStore) UpdateComment(ps chatychat.Comment) (err error) {
	result := p.db.Save(&ps)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (p *CommentStore) DeleteComment(id int) (err error) {
	result := p.db.Delete(&chatychat.Comment{}, id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}
