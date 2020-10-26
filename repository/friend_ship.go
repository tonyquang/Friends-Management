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

// DeleteFriendship friendship to DB
func DeleteFriendship(dbconn *sql.DB, UserOne string, UserTwo string) (*model_common.CommonRespone, error) {
	IsFriend, err := CheckIsFriendShip(dbconn, UserOne, UserTwo)

	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}

	if IsFriend == false {
		return nil, errors.New(UserOne + " and " + UserTwo + " are not friends")
	}

	sqlStatement := `delete from friends as f where (f.user_one_email = $1 OR f.user_one_email = $2) AND (f.user_two_email = $1 OR f.user_two_email = $2)`
	_, err = dbconn.Exec(sqlStatement, UserOne, UserTwo)
	if err != nil {
		return nil, errors.New("Check is friends error: " + err.Error())
	}
	return &model_common.CommonRespone{Success: true}, nil
}

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
