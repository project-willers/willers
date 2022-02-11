package model

import (
	"context"
	"time"
	"willers-api/db"
)

type Comment struct {
	DiaryUser   string    `json:"diaryUser" validate:"required"`
	DiaryTime   time.Time `json:"diaryTime" validate:"required"`
	CommentUser string    `json:"commentUser" validate:"required"`
	Comment     string    `json:"comment" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}

type Comments []*Comment

func GetComments(diary *Diary) (Comments, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM comments WHERE diary_user=? AND diary_time=?", diary.UserName, diary.SelectAt)
	if err != nil {
		return nil, err
	}

	var comments Comments
	for result.Next() {
		c := &Comment{}
		if err := result.Scan(c.DiaryUser, c.DiaryTime, c.CommentUser, c.Comment, c.CreatedAt, c.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func getComment(comment *Comment) (*Comment, error) {
	result := db.Database.QueryRowContext(context.Background(), "SELECT * FROM comments WHERE diary_user=? AND diary_time=? AND cmt_user=? AND cmt=?", comment.DiaryUser, comment.DiaryTime, comment.CommentUser, comment.Comment)
	if err := result.Scan(comment.DiaryUser, comment.DiaryTime, comment.CommentUser, comment.Comment, comment.CreatedAt, comment.UpdatedAt); err != nil {
		return nil, err
	}
	return comment, nil
}

func UpdateComment(comment *Comment) error {
	if _, err := getComment(comment); err != nil {
		return err
	}

	result := db.Database.QueryRowContext(context.Background(), "UPDATE comments SET cmt=? WHERE diary_user=? AND diary_time=? AND cmt_user=?", comment.DiaryUser, comment.DiaryTime, comment.CommentUser)
	if err := result.Scan(comment.DiaryUser, comment.DiaryTime, comment.CommentUser, comment.Comment, comment.CreatedAt, comment.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func AddComment(comment *Comment) error {
	if _, err := getComment(comment); err != nil {
		return err
	}

	result := db.Database.QueryRowContext(context.Background(), "INSERT INTO comments(diary_user, diary_time, cmt_user, cmt) VALUE(?, ?, ?, ?)", comment.DiaryUser, comment.DiaryTime, comment.CommentUser, comment.Comment)
	if err := result.Scan(comment.DiaryUser, comment.DiaryTime, comment.CommentUser, comment.Comment, comment.CreatedAt, comment.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func DeleteComment(comment *Comment) error {
	result := db.Database.QueryRowContext(context.Background(), "DELETE FROM comments WHERE diary_user=? AND diary_time=? AND cmt_user=? AND cmt=?", comment.DiaryUser, comment.DiaryTime, comment.CommentUser, comment.Comment)
	if err := result.Scan(comment.DiaryUser, comment.DiaryTime, comment.CommentUser, comment.Comment, comment.CreatedAt, comment.UpdatedAt); err != nil {
		return err
	}
	return nil
}