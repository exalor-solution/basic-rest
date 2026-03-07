package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/exalor-solution/rest-basic/model"
	"github.com/exalor-solution/rest-basic/pkg/dao"
	"github.com/exalor-solution/rest-basic/pkg/xLogger"
	"go.uber.org/zap"
)

type ISubscription interface {
	Add(context.Context, []byte) *model.XError
	Delete(context.Context, string) *model.XError
	Update(context.Context, string, []byte) *model.XError
	Find(context.Context, string) (string, *model.XError)
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

func (s Repo) Add(ctx context.Context, b []byte) *model.XError {
	s.l.Info(ctx, "adding subscription")

	if len(b) == 0 {
		s.l.Info(ctx, "Empty subscription")
		return model.NewInvalidArg("Empty subscription")
	}
	if err := json.Unmarshal(b, &s.S); err != nil {
		s.l.Error(ctx, "json is invalid", zap.Error(err))
		return model.NewInvalidArg("json is invalid")
	}
	if err := s.S.IsValid(); err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return model.NewInvalidArg(fmt.Sprintf("invalid argument: %v", err))
	}
	s.l.Info(ctx, "subscription added")
	err := db.Create(&s.S)
	if err != nil {
		s.l.Error(ctx, "create subscription", zap.Error(err))
		return model.NewInvalidArg(err.Error())
	}
	return model.NewSuccess()

}

func (s Repo) Delete(ctx context.Context, name string) *model.XError {
	s.l.Info(ctx, "deleting subscription", zap.String("name", name))
	if name == "" {
		s.l.Info(ctx, "Empty subscription")
		return model.NewInvalidArg("Empty subscription")
	}
	s.l.Info(ctx, "subscription deleted")
	err := db.Delete(name)
	if err != nil {
		s.l.Error(ctx, "delete subscription", zap.Error(err))
		return model.NewInvalidArg(err.Error())
	}
	return model.NewSuccess()
}

func (s Repo) Update(ctx context.Context, name string, b []byte) *model.XError {
	s.l.Info(ctx, "updating subscription", zap.String("name", name), zap.ByteString("input", b))
	if name == "" || len(b) == 0 {
		return model.NewInvalidArg("Empty name/json")
	}
	if err := json.Unmarshal(b, &s.S); err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return model.NewInvalidArg("json is invalid")
	}
	if err := s.S.IsValid(); err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return model.NewInvalidArg(fmt.Sprintf("invalid argument: %v", err))
	}
	s.l.Info(ctx, "subscription updated")
	err := db.Update(name, &s.S)
	if err != nil {
		s.l.Error(ctx, "update subscription", zap.Error(err))
		return model.NewInvalidArg(err.Error())
	}
	return model.NewSuccess()

}

func (s Repo) Find(ctx context.Context, name string) (string, *model.XError) {
	s.l.Info(ctx, "finding subscription", zap.String("name", name))
	if name == "" {
		s.l.Info(ctx, "Empty subscription")
		return "", model.NewInvalidArg("Empty subscription")
	}

	obj, err := db.Find(name)
	if err != nil {
		s.l.Error(ctx, "invalid argument", zap.Error(err))
		return "", model.NewInvalidArg(err.Error())
	}
	s.l.Info(ctx, "subscription returned")
	str, err := json.Marshal(obj)
	if err != nil {
		s.l.Error(ctx, "json is invalid", zap.Error(err))
		return "", model.NewInvalidArg(err.Error())
	}
	return string(str), model.NewSuccess()

}
