package service

type ISubscription interface {
	Add() error
	Delete() error
	Update() error
	Find() error
}

type Subscription struct {
}

func New() ISubscription {
	return Subscription{}
}

func (s Subscription) Add() error {
	//TODO implement me
	panic("implement me")
}

func (s Subscription) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (s Subscription) Update() error {
	//TODO implement me
	panic("implement me")
}

func (s Subscription) Find() error {
	//TODO implement me
	panic("implement me")
}
