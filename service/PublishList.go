package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
)

type Videos struct {
}

func (v Videos) Getlist(userid int64, token string) ([]model.Video, error) {
	//传入token进行验证
	_, ok := middleware.ParseToken(token)
	if ok != true {
		return []model.Video{}, errors.New("token验证不通过")
	}

	return []model.Video{}, nil
}
