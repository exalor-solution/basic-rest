package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string
type currency string

const (
	valid   = "Valid"
	invalid = "Invalid"
	USD     = "usd"
	EUR     = "eur"
	GBP     = "gbp"
	JPY     = "jpy"
	Cad     = "cad"
)

type Subscription struct {
	id        uuid.UUID
	Name      string   `json:"name"`
	Price     float64  `json:"price"`
	Currency  currency `json:"currency"`
	CreatedAt string
	updatedAt string
	stat      Status
}

func New() *Subscription {
	return &Subscription{
		id:        uuid.New(),
		Name:      "",
		Price:     0,
		Currency:  "",
		CreatedAt: time.Now().String(),
		updatedAt: "",
		stat:      valid,
	}
}

func (s *Subscription) GetId() uuid.UUID {
	return s.id
}
func (s *Subscription) GetStatus() Status {
	return s.stat
}
func (s *Subscription) SetStatus(status Status) {
	s.stat = status
}
func (s *Subscription) IsValid() error {
	if s.Name == "" {
		return errors.New("missing name")
	}
	if s.Price <= 0 || s.Price > 100 {
		return errors.New("invalid price(0< p <=100)")
	}
	if exist[currency](s.Currency, USD, EUR, JPY, GBP, Cad) {
		return errors.New("invalid currency")
	}
	return nil
}
func (s *Subscription) ToJson() (string, error) {
	byt, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
func (s *Subscription) FromJson(str string) error {
	err := json.Unmarshal([]byte(str), s)
	if err != nil {
		return err
	}
	return s.IsValid()
}
