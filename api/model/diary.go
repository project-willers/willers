package model

import "time"

type Diary struct {
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	SelectAt  time.Time `json:"time"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Diaries []*Diary

func GetDiary() (Diaries, error) {
	return nil
}

func UpdateDiary() error {
	return nil
}

func AddDiary() error {
	// TODO
	// すでに存在しているならUpdateDiaryにRedirectさせる処理
	return nil
}

func DeleteDiary() error {
	return nil
}
