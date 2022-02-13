package model

import (
	"context"
	"database/sql"
	"willers-api/db"
)

type Friend struct {
	MyName    string `json:"myname" validate:"required"`
	OtherName string `json:"othername" validate:"required"`
}

type FriendResponse struct {
	MyName    string `json:"myname" validate:"required"`
	OtherName string `json:"othername" validate:"required"`
	Ok        bool   `json:"response" validate:"required"`
}

type Friends []*Friend

func FindFriend(name string) (*Friend, error) {
	result := db.Database.QueryRowContext(context.Background(), "SELECT * FROM friends WHERE name=?", name)
	friend := &Friend{}
	if err := result.Scan(friend.MyName, friend.OtherName); err != nil {
		return nil, err
	}
	return friend, nil
}

func FindFriends(friend *Friend) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friends WHERE name=?", friend.MyName)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var friends Friends
	for result.Next() {
		fri := &Friend{}
		if err := result.Scan(fri.MyName, fri.OtherName); err != nil {
			return nil, err
		}
		friends = append(friends, fri)
	}
	return friends, nil
}

func FindFriendRequest(req *Friend) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friendrequests WHERE name=?", req.MyName)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var friReqs Friends
	for result.Next() {
		friReq := &Friend{}
		if err := result.Scan(friReq.MyName, friReq.OtherName); err != nil {
			return nil, err
		}
		friReqs = append(friReqs, friReq)
	}
	return friReqs, nil
}

func FriendRequest(req *Friend) error {
	if _, err := FindFriendRequest(req); err != nil {
		return err
	}
	result := db.Database.QueryRowContext(context.Background(), "INSERT INTO friendrequests(name, other) VALUE(?, ?)", req.MyName, req.OtherName)
	if err := result.Scan(req.MyName, req.OtherName); err != nil {
		return err
	}
	return nil
}

func AddFriend(res *FriendResponse) error {
	req := &Friend{
		MyName:    res.MyName,
		OtherName: res.OtherName,
	}
	if _, err := FindFriendRequest(req); err != nil {
		return err
	}
	var result *sql.Row
	if res.Ok {
		result = db.Database.QueryRowContext(context.Background(), "INSERT INTO friends(name, other) VALUE(?, ?)", req.MyName, req.OtherName)
	}
	db.Database.QueryRowContext(context.Background(), "DELETE FROM friendrequests WHERE name=? AND other=?", req.MyName, req.OtherName)
	if err := result.Scan(req.MyName, req.OtherName); err != nil {
		return err
	}
	return nil
}

func DeleteFriend(friend *Friend) error {
	result := db.Database.QueryRowContext(context.Background(), "DELETE FROM friendrequests WHERE name=? AND other=?", friend.MyName, friend.OtherName)
	if err := result.Scan(friend.MyName, friend.OtherName); err != nil {
		return err
	}
	return nil
}
