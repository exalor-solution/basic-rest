package dao

import (
	"errors"

	"github.com/exalor-solution/rest-basic/model"
)

type Dao struct {
	Subs []model.Subscription
}

func NewDao() *Dao {
	return &Dao{Subs: []model.Subscription{}}
}

func (d *Dao) Find(name string) (*model.Subscription, error) {
	if d.Subs == nil || len(d.Subs) == 0 {
		return nil, errors.New("no subs")
	}
	for _, sub := range d.Subs {
		if sub.Name == name {
			return &sub, nil
		}
	}
	return nil, errors.New("not found")
}
func (d *Dao) Create(s *model.Subscription) error {
	if s == nil {
		return errors.New("no subs")
	}
	for _, sub := range d.Subs {
		if sub.Name == sub.Name {
			return errors.New("already exists")
		}
	}
	d.Subs = append(d.Subs, *s)
	return nil
}

func (d *Dao) Update(name string, s *model.Subscription) error {
	if d.Subs == nil || len(d.Subs) == 0 {
		return errors.New("no subs")
	}
	for k, sub := range d.Subs {
		if sub.Name == name {
			d.Subs[k] = *s
			return nil
		}
	}
	return errors.New("not found")
}

func (d *Dao) Delete(name string) error {
	if d.Subs == nil || len(d.Subs) == 0 {
		return errors.New("no subs")
	}
	for k, sub := range d.Subs {
		if sub.Name == name {
			d.Subs = append(d.Subs[:k], d.Subs[k+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
