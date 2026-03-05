package service

import (
	"encoding/json"

	"github.com/exalor-solution/rest-basic/model"
)

type ISubscription interface {
	Add([]byte) error
	Delete() error
	Update() error
	Find() error
}

type Repo struct {
	S model.Subscription
}

func New() ISubscription {
	return Repo{S: *model.New()}
}

func (s Repo) Add(b []byte) error {
	if err := json.Unmarshal(b, &s.S); err != nil {
		return err
	}
	if err := s.S.IsValid(); err != nil {
		return err
	}

	return nil

}

func (s Repo) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (s Repo) Update() error {
	//TODO implement me
	panic("implement me")
}

func (s Repo) Find() error {
	//TODO implement me
	panic("implement me")
}
