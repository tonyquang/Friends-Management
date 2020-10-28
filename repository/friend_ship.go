package repository

import (
	"database/sql"
	"errors"
	model_common "friends_management/models/respone"
)

// InsertFriendship insert new friendship to DB
func InsertFriendship(dbconn *sql.DB, UserOne string, UserTwo string) (*model_common.CommonRespone, error) {

	IsFriend, err := CheckIsFriendShip(dbconn, UserOne, UserTwo)

	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}

	if IsFriend == true {
		return nil, errors.New(UserOne + " and " + UserTwo + " were friend")
	}

	sqlStatement := `INSERT INTO public.friends(user_one_email, user_two_email, update_status) VALUES ($1, $2, $3)`
	_, err = dbconn.Exec(sqlStatement, UserOne, UserTwo, true)
	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}
	return &model_common.CommonRespone{Success: true}, nil
}

// Delete Friendship to DB
func DeleteFriendship(dbconn *sql.DB, id int) (*model_common.CommonRespone, error) {
	IsFriendship, err := CheckFriendShipExist(dbconn, id)

	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}

	if IsFriendship == false {
		return nil, errors.New("ID Friendship not found!")
	}

	sqlStatement := `delete from friends as f where f.id = $1`
	_, err = dbconn.Exec(sqlStatement, id)
	if err != nil {
		return nil, errors.New("Unfriend error: " + err.Error())
	}
	return &model_common.CommonRespone{Success: true}, nil
}

// Get List Friends By Email
func GetListFriendByEmail(dbconn *sql.DB, mailAdress string) (*model_common.ListFriendsRespone, error) {
	ExistEmail, err := CheckEmailInTableFriendship(dbconn, mailAdress)

	if err != nil {
		return nil, errors.New("Check exist email error: " + err.Error())
	}

	if ExistEmail == false {
		return nil, errors.New("User not found!")
	}

	var row *sql.Rows
	count := 0

	sqlStatement := `SELECT f1.user_two_email friend FROM friends as f1 WHERE f1.user_one_email = $1 UNION ALL SELECT f2.user_one_email friend FROM friends as f2 WHERE f2.user_two_email = $1`
	row, err = dbconn.Query(sqlStatement, mailAdress)
	if err != nil {
		return nil, errors.New("List friends query error: " + err.Error())
	}

	var list_friend model_common.ListFriendsRespone

	for row.Next() {
		var friend string
		err := row.Scan(&friend)
		if err != nil {
			return nil, errors.New("List friends scan error: " + err.Error())
		}
		list_friend.ListFriends = append(list_friend.ListFriends, friend)
		count++
	}
	list_friend.Success = true
	list_friend.Count = count
	return &list_friend, nil
}

// Get Common Friends List Between Two Users By Email
func GetCommonFriendsListByEmail(dbconn *sql.DB, UserOne string, UserTwo string) (*model_common.ListFriendsRespone, error) {
	IsFriend, err := CheckIsFriendShip(dbconn, UserOne, UserTwo)
	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}

	if IsFriend == false {
		return nil, errors.New(UserOne + " and " + UserTwo + " are not friend")
	}

	var row *sql.Rows
	count := 0

	sqlStatement := `SELECT f1.user_two_email friend FROM friends as f1 WHERE f1.user_one_email = $1 UNION ALL SELECT f2.user_one_email friend FROM friends as f2 WHERE f2.user_two_email = $2`
	row, err = dbconn.Query(sqlStatement, UserOne, UserTwo)
	if err != nil {
		return nil, errors.New("List friends query error: " + err.Error())
	}

	var list_friend model_common.ListFriendsRespone

	for row.Next() {
		var friend string
		err := row.Scan(&friend)
		if err != nil {
			return nil, errors.New("List friends scan error: " + err.Error())
		}
		list_friend.ListFriends = append(list_friend.ListFriends, friend)
		count++
	}
	list_friend.Success = true
	list_friend.Count = count
	return &list_friend, nil
}

//Subscribe Update between two user
func UpdateSubscribeStatusFriend(dbconn *sql.DB, requestor string, tagert string) (*model_common.CommonRespone, error) {
	IsFriend, err := CheckIsFriendShip(dbconn, requestor, tagert)
	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}

	if IsFriend == false {
		return nil, errors.New(requestor + " and " + tagert + " are not friend")
	}
	err = UpdateStatusFriend(dbconn, requestor, tagert, true) // subscribe status is true
	if err != nil {
		return nil, err
	} else {
		return &model_common.CommonRespone{Success: true}, nil
	}
}

//Subscribe Update between two user
func BlockStatusFriend(dbconn *sql.DB, requestor string, tagert string) (*model_common.CommonRespone, error) {
	IsFriend, err := CheckIsFriendShip(dbconn, requestor, tagert)
	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}

	if IsFriend == false {
		return nil, errors.New(requestor + " and " + tagert + " are not friend")
	}
	err = UpdateStatusFriend(dbconn, requestor, tagert, false) // block status is false
	if err != nil {
		return nil, err
	} else {
		return &model_common.CommonRespone{Success: true}, nil
	}
}

//Func Exce Update status friend
func UpdateStatusFriend(dbconn *sql.DB, requestor string, tagert string, status bool) error {
	stm := `Update friends SET update_status = $3 where (user_one_email = $1 OR user_two_email = $1) AND (user_one_email = $2 OR user_two_email = $2)`
	_, err := dbconn.Exec(stm, requestor, tagert, status)
	if err != nil {
		return errors.New("Error Exec Update Status Friend: " + err.Error())
	} else {
		return nil
	}
}

//List all email addresses that can receive updates from an email address.
func GetAllEmailReceiveUpdate(dbconn *sql.DB, mailAdress string) (*model_common.ListFriendsRecviceUpdateRespone, error) {
	IsExist, err := CheckEmailInTableFriendship(dbconn, mailAdress)
	if err != nil {
		return nil, errors.New("Check exist email error: " + err.Error())
	}

	if IsExist == false {
		return nil, errors.New("User not found!")
	}

	var row *sql.Rows
	count := 0

	sqlStatement := `select f1.user_two_email email from friends as f1 where f1.user_one_email = $1 AND f1.update_status = true UNION ALL select f2.user_one_email email from friends as f2 where f2.user_two_email = $1 AND f2.update_status = true`
	row, err = dbconn.Query(sqlStatement, mailAdress)
	if err != nil {
		return nil, errors.New("List friends query error: " + err.Error())
	}

	var list_friend model_common.ListFriendsRecviceUpdateRespone

	for row.Next() {
		var friend string
		err := row.Scan(&friend)
		if err != nil {
			return nil, errors.New("List friends scan error: " + err.Error())
		}
		list_friend.Recipients = append(list_friend.Recipients, friend)
		count++
	}
	list_friend.Success = true
	return &list_friend, nil
}

//Check Two email was friend?
func CheckIsFriendShip(dbconn *sql.DB, UserOne string, UserTwo string) (bool, error) {
	var rs *sql.Rows
	var err error
	rs, err = dbconn.Query("select COUNT(*) from public.friends as f where (f.user_one_email = $1 OR f.user_one_email = $2) AND (f.user_two_email = $1 OR f.user_two_email = $2)", UserOne, UserTwo)
	if err != nil {
		return false, err
	}

	defer rs.Close()

	var count int

	for rs.Next() {
		if err := rs.Scan(&count); err != nil {
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}

}

//Check ID Friendship is exist
func CheckFriendShipExist(dbconn *sql.DB, id int) (bool, error) {
	var rs *sql.Rows
	var err error
	rs, err = dbconn.Query("select COUNT(*) from public.friends as f where f.id = $1", id)
	if err != nil {
		return false, err
	}

	defer rs.Close()

	var count int

	for rs.Next() {
		if err := rs.Scan(&count); err != nil {
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}

}

//Check An Email In Table Friends
func CheckEmailInTableFriendship(dbconn *sql.DB, mailAdress string) (bool, error) {
	var rs *sql.Rows
	var err error
	rs, err = dbconn.Query("select COUNT(*) from public.friends as f where f.user_one_email = $1 or f.user_two_email = $1", mailAdress)
	if err != nil {
		return false, err
	}

	defer rs.Close()

	var count int

	for rs.Next() {
		if err := rs.Scan(&count); err != nil {
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}

}
