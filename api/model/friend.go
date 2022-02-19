package model

import (
	"context"
	"log"
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

func FindFriend(friend *Friend) (*Friend, error) {
	result := db.Database.QueryRowContext(context.Background(), "SELECT * FROM friends WHERE name=? AND other=?", friend.MyName, friend.OtherName)
	if err := result.Scan(friend.MyName, friend.OtherName); err != nil {
		return nil, err
	}
	return friend, nil
}

func FindFriends(name string) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friends WHERE name=?", name)
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

func FindFriendRequest(friend *Friend) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friendrequests WHERE name=? AND other=?", friend.MyName, friend.OtherName)
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

func FindFriendRequests(name string) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friendrequests WHERE name=?", name)
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
	//	if _, err := FindFriendRequest(req); err != nil {
	//		return err
	//	}

	insert, err := db.Database.Prepare("INSERT INTO friendrequests(name, other) VALUE(?, ?)")
	if err != nil {
		return err
	}
	defer insert.Close()
	result, err := insert.ExecContext(context.Background(), req.MyName, req.OtherName)
	if err != nil {
		return err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Println(rowCnt)
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
	if res.Ok {
		insert, err := db.Database.Prepare("INSERT INTO friends(name, other) VALUE(?, ?)")
		if err != nil {
			return err
		}
		defer insert.Close()
		_, err = insert.ExecContext(context.Background(), req.MyName, req.OtherName)
		if err != nil {
			return err
		}
	}
	del, err := db.Database.Prepare("DELETE FROM friendrequests WHERE name=? AND other=?")
	if err != nil {
		return err
	}
	defer del.Close()
	_, err = del.ExecContext(context.Background(), req.MyName, req.OtherName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFriend(friend *Friend) error {
	if _, err := FindFriend(friend); err != nil {
		return err
	}

	del, err := db.Database.Prepare("DELETE FROM friends WHERE name=? AND other=?")
	if err != nil {
		return err
	}
	defer del.Close()
	_, err = del.ExecContext(context.Background(), friend.MyName, friend.OtherName)
	if err != nil {
		return err
	}
	return nil
}
