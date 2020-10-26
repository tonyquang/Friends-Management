package repository

import (
	"database/sql"
	"errors"
	model_common "friends_management/models/respone"
	"log"
)

// InsertFriendship insert new friendship to DB
func InsertFriendship(dbconn *sql.DB, UserOne string, UserTwo string) (model_common.CommonRespone, error) {
	tx, err := dbconn.Begin()
	if err != nil {
		log.Fatal(err)
		return model_common.CommonRespone{Success: false}, err
	}

	IsFriend, err := CheckIsFriendShip(dbconn, UserOne, UserTwo)

	if err != nil {
		return model_common.CommonRespone{Success: false}, errors.New("Check is friends error: " + err.Error())
	}

	if IsFriend == true {
		return model_common.CommonRespone{Success: false}, errors.New(UserOne + " and " + UserTwo + " was friend")
	}

	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare("INSERT INTO public.friends(user_one_email, user_two_email, update_status) VALUES (?, ?, ?)")
	if err != nil {
		return model_common.CommonRespone{Success: false}, errors.New("Fail Insert Friendship with error: " + err.Error())
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	//When be Friend, Both Users can retrive update
	if _, err := stmt.Exec(UserOne, UserTwo, true); err != nil {
		return model_common.CommonRespone{Success: false}, err
	}
	if err := tx.Commit(); err != nil {
		return model_common.CommonRespone{Success: false}, err
	} else {
		return model_common.CommonRespone{Success: true}, nil
	}
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
