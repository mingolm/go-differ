package go_differ

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type User struct {
	Id        uint64       `differ:"id"`
	Name      string       `differ:"name"`
	Gender    Gender       `differ:"gender"`
	IsTest    bool         `differ:"is_test"`
	CreatedAt time.Time    `differ:"created_at"`
	UpdatedAt time.Time    `differ:"-"` // 该字段忽略
	DeletedAt sql.NullTime `differ:"deleted_at"`
}

type Gender uint8

const (
	GenderUnknown Gender = iota
	GenderMale
	GenderFemale
	GenderNonBinary
)

func TestDiffer_GetChanges(t *testing.T) {
	a := assert.New(t)
	userA := User{
		Id:        1,
		Name:      "mingo",
		Gender:    GenderMale,
		IsTest:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
	}
	df := NewDiffer(userA, WithIgnoreFields("id"))
	a.Equal(false, df.IsDirty(userA))

	userA.Id = 2 // 配置了忽略
	userA.IsTest = false
	userA.Name = "mingo_test"
	userA.UpdatedAt = time.Now().Add(time.Hour) // tag 标记忽略
	a.Equal(true, df.IsDirty(userA))
	a.Equal(2, len(df.GetChangeKeys(userA)))

	for key, val := range df.GetChanges(userA) {
		t.Logf("key: %s, val: %+v\n", key, val)
	}
}
