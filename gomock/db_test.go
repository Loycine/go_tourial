package gomock

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	m.EXPECT().Get(gomock.Eq("Rory")).Return(100, nil)

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}

	if v := GetFromDB(m, "Rory"); v != 100 {
		t.Fatal("expected 100, but got", v)
	}
}
