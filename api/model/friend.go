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

type F struct {
	Friends Friends
}

func (friend *Friend) FindFriend() error {
	result := db.Database.QueryRowContext(context.Background(), "SELECT * FROM friends WHERE name=? AND other=?", friend.MyName, friend.OtherName)
	if err := result.Scan(&friend.MyName, &friend.OtherName); err != nil {
		return err
	}
	return nil
}

func FindFriends(name string) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friends WHERE name=?", name)
	if err != nil {
		log.Println("FindFriends QueryContext() Error: ", err)
		return nil, err
	}
	defer result.Close()

	var friends Friends
	for result.Next() {
		var f Friend
		if err := result.Scan(&f.MyName, &f.OtherName); err != nil {
			log.Println("FindFriends Scan() Error: ", err)
			return nil, err
		}
		friends = append(friends, &f)
	}
	log.Println(friends)
	return friends, nil
}

func FindFriendRequest(name string) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friendrequests WHERE name=?", name)
	if err != nil {
		log.Println("FindFriendRequest QueryContext() Error: ", err)
		return nil, err
	}
	defer result.Close()

	var friReqs Friends
	for result.Next() {
		var friReq Friend
		if err := result.Scan(&friReq.MyName, &friReq.OtherName); err != nil {
			log.Println("FindFriendRequest Scan() Error: ", err)
			return nil, err
		}
		friReqs = append(friReqs, &friReq)
	}
	return friReqs, nil
}

func GetMyFriendRequests(name string) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friendrequests WHERE name=?", name)
	if err != nil {
		log.Println("FindFriendRequests QueryContext() Error: ", err)
		return nil, err
	}
	defer result.Close()

	var friReqs Friends
	for result.Next() {
		var friReq Friend
		if err := result.Scan(&friReq.MyName, &friReq.OtherName); err != nil {
			log.Println("FindFriendRequests Scan() Error: ", err)
			return nil, err
		}
		friReqs = append(friReqs, &friReq)
	}
	return friReqs, nil
}

func GetOtherFriendRequests(other string) (Friends, error) {
	result, err := db.Database.QueryContext(context.Background(), "SELECT * FROM friendrequests WHERE other=?", other)
	if err != nil {
		log.Println("FindFriendRequests QueryContext() Error: ", err)
		return nil, err
	}
	defer result.Close()

	var friReqs Friends
	for result.Next() {
		var friReq Friend
		if err := result.Scan(&friReq.MyName, &friReq.OtherName); err != nil {
			log.Println("FindFriendRequests Scan() Error: ", err)
			return nil, err
		}
		friReqs = append(friReqs, &friReq)
	}
	return friReqs, nil
}

func FriendRequest(req *Friend) error {
	if err := req.FindFriend(); err == nil {
		log.Println("FriendRequest FindFriend() Error: ", err)
		return err
	}

	insert, err := db.Database.Prepare("INSERT INTO friendrequests(name, other) VALUE(?, ?)")
	if err != nil {
		log.Println("FriendRequest Prepare() Error: ", err)
		return err
	}
	defer insert.Close()
	result, err := insert.ExecContext(context.Background(), req.MyName, req.OtherName)
	if err != nil {
		log.Println("FriendRequest ExecContent() Error: ", err)
		return err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		log.Println("FriendRequest RowsAffected() Error: ", err)
		return err
	}
	log.Println("FriendRequest rowCnt: ", rowCnt)
	return nil
}

func AddFriend(res *FriendResponse) error {
	req := &Friend{
		MyName:    res.MyName,
		OtherName: res.OtherName,
	}
	if _, err := GetOtherFriendRequests(res.MyName); err != nil {
		log.Println("AddFriend FindFriendRequest() Error: ", err)
		return err
	}
	if res.Ok {
		insert, err := db.Database.Prepare("INSERT INTO friends(name, other) VALUE(?, ?)")
		if err != nil {
			log.Println("AddFriend Insert-Prepare() Error: ", err)
			return err
		}
		defer insert.Close()
		_, err = insert.ExecContext(context.Background(), req.MyName, req.OtherName)
		if err != nil {
			log.Println("AddFriend 1-Insert-ExecContext() Error: ", err)
			return err
		}
		_, err = insert.ExecContext(context.Background(), req.OtherName, req.MyName)
		if err != nil {
			log.Println("AddFriend 2-Insert-ExecContext() Error: ", err)
			return err
		}
	}
	del, err := db.Database.Prepare("DELETE FROM friendrequests WHERE name=? AND other=?")
	if err != nil {
		log.Println("AddFriend Delete-Prepare() Error: ", err)
		return err
	}
	defer del.Close()
	_, err = del.ExecContext(context.Background(), req.MyName, req.OtherName)
	if err != nil {
		log.Println("AddFriend Delete-ExecContext() Error: ", err)
		return err
	}
	return nil
}

func DeleteFriend(friend *Friend) error {
	if err := friend.FindFriend(); err != nil {
		return err
	}

	del, err := db.Database.Prepare("DELETE FROM friends WHERE name=? AND other=?")
	if err != nil {
		log.Println("DeleteFriend Prepare() Error: ", err)
		return err
	}
	defer del.Close()
	_, err = del.ExecContext(context.Background(), friend.MyName, friend.OtherName)
	if err != nil {
		log.Println("DeleteFriend ExecContext() Error: ", err)
		return err
	}
	return nil
}
