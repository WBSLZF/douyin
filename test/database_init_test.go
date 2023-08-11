package test

import (
	"testing"

	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func TestDB(t *testing.T) {
	DB := model.DB
	_ = DB
}
