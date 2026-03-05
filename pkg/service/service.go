package service

import (
	"encoding/json"
	"errors"

	"github.com/exalor-solution/rest-basic/model"
	"github.com/exalor-solution/rest-basic/pkg/dao"
)

type ISubscription interface {
	Add([]byte) error
	Delete(string) error
	Update(string, []byte) error
	Find(string) (string, error)
}

var db *dao.Dao

type Repo struct {
	S model.Subscription
}

func New() ISubscription {
	db = dao.NewDao()
	return Repo{S: *model.New()}
}

func (s Repo) Add(b []byte) error {
	if len(b) == 0 {
		return errors.New("invalid argument")
	}
	if err := json.Unmarshal(b, &s.S); err != nil {
		return err
	}
	if err := s.S.IsValid(); err != nil {
		return err
	}

	db.Create(&s.S)
	return nil

}

func (s Repo) Delete(name string) error {
	if name == "" {
		return errors.New("invalid argument")
	}
	return db.Delete(name)
}

func (s Repo) Update(name string, b []byte) error {
	if name == "" || len(b) == 0 {
		return errors.New("invalid argument")
	}
	if err := json.Unmarshal(b, &s.S); err != nil {
		return err
	}
	if err := s.S.IsValid(); err != nil {
		return err
	}

	return db.Update(name, &s.S)

}

func (s Repo) Find(name string) (string, error) {
	if name == "" {
		return "", errors.New("invalid argument")
	}
	obj, err := db.Find(name)
	if err != nil {
		return "", err
	}
	return obj.ToJson()

}
