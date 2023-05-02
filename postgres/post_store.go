package postgres

import (
	"github.com/asgharalitaj/chatychat"
	"gorm.io/gorm"
)

func NewPostStore() *PostStore{
  return &PostStore {
    db: GETDB(),
  }
}

type PostStore struct {
  db *gorm.DB
}


func (p *PostStore) GetItem(id int) (post chatychat.Post, err error) {
	result := p.db.First(&post, id)
	if result.Error != nil {
		err = result.Error
		return 
	}
	return post, nil
}

func (p *PostStore) GetItems() (Posts []chatychat.Post, err error) {
	result := p.db.Find(&Posts)
	if result.Error != nil {
		return
	}
	return Posts, nil
}

func (p *PostStore) CreateItem(ps chatychat.Post) (err error) {
	result := p.db.Create(&ps)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (p *PostStore) UpdateItem(ps chatychat.Post) (err error) {
	result := p.db.Save(&ps)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (p *PostStore) DeleteItem(id int) (err error) {
	result := p.db.Delete(&chatychat.Post{}, id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}
