package main

import (
	"database/sql"
	"fmt"
	go_differ "github.com/mingolm/go-differ"
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
	GenderMale           = iota
	GenderFemale
	GenderNonBinary
)

func main() {
	userA := User{
		Id:        1,
		Name:      "mingo",
		Gender:    GenderMale,
		IsTest:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
	}

	df := go_differ.NewDiffer(userA, go_differ.WithIgnoreFields("id"))
	fmt.Printf("id_dirty: %T\n", df.IsDirty(userA))

	userA.Id = 2 // options 配置了忽略
	userA.IsTest = false
	userA.Name = "mingo_test"
	userA.UpdatedAt = time.Now().Add(time.Hour) // tag 配置了忽略
	for key, val := range df.GetChanges(userA) {
		fmt.Printf("key: %s, val: %+v\n", key, val)
	}
}
