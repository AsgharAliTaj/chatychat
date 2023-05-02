package postgres

import (
	"github.com/asgharalitaj/chatychat"
	"gorm.io/gorm"
)

func NewThreadStore() *ThreadStore {
  return &ThreadStore {
    db: GETDB(),
  }
}

type ThreadStore struct {
  db *gorm.DB
}

func (t *ThreadStore) GetThread(id int) (thread chatychat.Thread, err error) {
	result := t.db.First(&thread, id)
	if result.Error != nil {
		err = result.Error
		return 
	}
	return thread, nil
}

func (t *ThreadStore) GetThreads() (threads []chatychat.Thread, err error) {
	result := t.db.Find(&threads)
	if result.Error != nil {
		return
	}
	return threads, nil
}

func (t *ThreadStore) CreateThread(th chatychat.Thread) (err error) {
	result := t.db.Create(&th)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (t *ThreadStore) UpdateThread(th chatychat.Thread) (err error) {
	result := t.db.Save(&th)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (t *ThreadStore) DeleteThread(id int) (err error) {
	result := t.db.Delete(&chatychat.Thread{}, id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}
