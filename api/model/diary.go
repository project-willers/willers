package model

import (
	"context"
	"log"
	"time"
	"willers-api/db"
)

type Diary struct {
	UserName  string `json:"name" validate:"required"`
	Content   string `json:"content" validate:"required"`
	SelectAt  string `json:"selectAt" validate:"required"`
	UpdatedAt string `json:"updatedAt" validate:"required"`
}

type Diaries []*Diary

type D struct {
	Diaries Diaries
}

func (diary *Diary) GetDiary() error {
	t, err := time.Parse("2006-01-02 15:04:05", diary.SelectAt)
	if err != nil {
		log.Println("GetDiary Parse() Error: ", err)
		return err
	}
	result := db.Database.QueryRowContext(context.Background(), "SELECT * FROM diaries WHERE name=? AND content=? AND select_at=?", diary.UserName, diary.Content, t)

	if err := result.Scan(&diary.UserName, &diary.Content, &diary.SelectAt, &diary.UpdatedAt); err != nil {
		log.Println("GetDiary Scan() Error: ", err)
		return err
	}
	return nil
}

func GetDiaries(name string) (Diaries, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM diaries WHERE name=?", name)
	if err != nil {
		log.Println("GetDiaries QueryContext() Error: ", err)
		return nil, err
	}
	defer result.Close()

	var diaries Diaries
	for result.Next() {
		var d Diary
		if err := result.Scan(&d.UserName, &d.Content, &d.SelectAt, &d.UpdatedAt); err != nil {
			log.Println("GetDiaries Scan() Error: ", err)
			return nil, err
		}
		diaries = append(diaries, &d)
	}
	log.Println(diaries)
	return diaries, nil
}

func UpdateDiary(diary *Diary) error {
	if _, err := GetDiaries(diary.UserName); err != nil {
		log.Println("UpdateDiary GetDiaries() Error: ", err)
		return err
	}

	update, err := db.Database.Prepare("UPDATE diaries SET content=? WHERE name=? AND select_at=?")
	if err != nil {
		log.Println("UpdateDiary GetDiaries() Error: ", err)
		return err
	}
	defer update.Close()
	t, err := time.Parse("2006-01-02 15:04:05", diary.SelectAt)
	if err != nil {
		log.Println("UpdateDiary time.Parse() Error: ", err)
		return err
	}
	result, err := update.ExecContext(context.Background(), diary.Content, diary.UserName, t)
	if err != nil {
		log.Println("UpdateDiary ExecContext() Error: ", err)
		return err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		log.Println("UpdateDiary RowsAffected() Error: ", err)
		return err
	}
	log.Println(rowCnt)
	return nil
}

func AddDiary(diary *Diary) error {
	if _, err := GetDiaries(diary.UserName); err != nil {
		log.Println("UpdateDiary RowsAffected() Error: ", err)
		return err
	}
	insert, err := db.Database.Prepare("INSERT INTO diaries(name, content, select_at) VALUE(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer insert.Close()
	t, err := time.Parse("2006-01-02 15:04:05", diary.SelectAt)
	if err != nil {
		return err
	}
	result, err := insert.ExecContext(context.Background(), diary.UserName, diary.Content, t)
	if err != nil {
		log.Println(err)
		return err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(rowCnt)
	return nil
}

func DeleteDiary(diary *Diary) error {
	if err := diary.GetDiary(); err == nil {
		log.Println(err)
		return err
	}
	log.Println(diary)
	del, err := db.Database.Prepare("DELETE FROM diaries WHERE name=? AND content=? AND select_at=?")
	if err != nil {
		log.Println(err)
		return err
	}
	defer del.Close()
	result, err := del.ExecContext(context.Background(), diary.UserName, diary.Content, diary.SelectAt)
	if err != nil {
		log.Println(err)
		return err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(rowCnt)
	return nil
}
