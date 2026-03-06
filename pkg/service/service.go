package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/exalor-solution/rest-basic/model"
	"github.com/exalor-solution/rest-basic/pkg/dao"
	"github.com/exalor-solution/rest-basic/pkg/xLogger"
	"go.uber.org/zap"
)

type ISubscription interface {
	Add(context.Context, []byte) error
	Delete(context.Context, string) error
	Update(context.Context, string, []byte) error
	Find(context.Context, string) (string, error)
}

var db *dao.Dao

type Repo struct {
	S model.Subscription
	l xLogger.ILogger
}

func New(logger xLogger.ILogger) ISubscription {
	db = dao.NewDao()
	return Repo{S: *model.New(), l: logger}
}

func (s Repo) Add(ctx context.Context, b []byte) error {
	s.l.Info(ctx, "adding subscription")

	if len(b) == 0 {
		s.l.Info(ctx, "Empty subscription")
		return errors.New("invalid argument")
	}
	if err := json.Unmarshal(b, &s.S); err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return err
	}
	if err := s.S.IsValid(); err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return err
	}
	s.l.Info(ctx, "subscription added")
	return db.Create(&s.S)

}

func (s Repo) Delete(ctx context.Context, name string) error {
	s.l.Info(ctx, "deleting subscription", zap.String("name", name))
	if name == "" {
		s.l.Info(ctx, "Empty subscription")
		return errors.New("invalid argument")
	}
	s.l.Info(ctx, "subscription deleted")
	return db.Delete(name)
}

func (s Repo) Update(ctx context.Context, name string, b []byte) error {
	s.l.Info(ctx, "updating subscription", zap.String("name", name), zap.ByteString("input", b))
	if name == "" || len(b) == 0 {
		return errors.New("invalid argument")
	}
	if err := json.Unmarshal(b, &s.S); err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return err
	}
	if err := s.S.IsValid(); err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return err
	}
	s.l.Info(ctx, "subscription updated")
	return db.Update(name, &s.S)

}

func (s Repo) Find(ctx context.Context, name string) (string, error) {
	s.l.Info(ctx, "finding subscription", zap.String("name", name))
	if name == "" {
		s.l.Info(ctx, "Empty subscription")
		return "", errors.New("invalid argument")
	}

	obj, err := db.Find(name)
	if err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return "", err
	}
	s.l.Info(ctx, "subscription returned")
	return obj.ToJson()

}
