package model

import (
	"context"
	"time"
	"willers-api/db"
)

type Diary struct {
	UserName  string    `json:"name"`
	Content   string    `json:"content"`
	SelectAt  time.Time `json:"selectAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Diaries []*Diary

func GetDiary(name string) (Diaries, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM diaries WHERE name=?", name)
	if err != nil {
		return nil, err
	}

	var diaries Diaries
	for result.Next() {
		d := &Diary{}
		if err := result.Scan(d.UserName, d.Content, d.SelectAt, d.CreatedAt, d.UpdatedAt); err != nil {
			return nil, err
		}
		diaries = append(diaries, d)
	}
	return diaries, nil
}

func UpdateDiary(diary *Diary) error {
	if _, err := GetDiary(diary.UserName); err != nil {
		return err
	}

	result := db.Database.QueryRowContext(context.Background(), "UPDATE diaries SET content=? WHERE name=? AND select_at=?", diary.Content, diary.UserName, diary.SelectAt)
	if err := result.Scan(diary.UserName, diary.Content, diary.SelectAt, diary.CreatedAt, diary.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func AddDiary(diary *Diary) error {
	// TODO
	// すでに存在しているならUpdateDiaryにRedirectさせる処理
	// フロント処理にした方がいい？
	if _, err := GetDiary(diary.UserName); err != nil {
		return err
	}

	result := db.Database.QueryRowContext(context.Background(), "INSERT INTO diaries(name, content, select_at) VALUE(?, ?, ?)", diary.UserName, diary.Content, diary.SelectAt)
	if err := result.Scan(diary.UserName, diary.Content, diary.SelectAt, diary.CreatedAt, diary.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func DeleteDiary(diary *Diary) error {
	result := db.Database.QueryRowContext(context.Background(), "DELETE FROM diaries WHERE name=? AND content=?, select_at", diary.UserName, diary.Content, diary.SelectAt)
	if err := result.Scan(diary.UserName, diary.Content, diary.SelectAt); err != nil {
		return err
	}
	return nil
}
