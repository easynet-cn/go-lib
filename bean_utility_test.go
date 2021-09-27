package golib

import (
	"testing"
	"time"
)

const (
	dateTimePattern = "2006-01-02 15:04:05"
)

type A struct {
	Id         int64
	Name       *string
	CreateTime *time.Time
	UpdateTime time.Time
}

type B struct {
	Id         int64
	Name       string
	CreateTime time.Time
	UpdateTime string
}

func TestCopy(t *testing.T) {
	name := "testName"
	createTime := time.Now()

	a := A{Id: 1, Name: &name, CreateTime: &createTime, UpdateTime: time.Now()}
	b := &B{}

	beanUtility := NewBeanUtiltiy(dateTimePattern, LowerCamelCase)

	if _, err := beanUtility.Copy(a, b); err != nil {
		t.Error(a, b, err)
	}

	if b.Id != a.Id ||
		b.Name != *a.Name ||
		b.CreateTime.Format(dateTimePattern) != a.CreateTime.Format(dateTimePattern) ||
		b.UpdateTime != a.UpdateTime.Format(dateTimePattern) {
		t.Error(b)
	}
}
