package services

import (
	"database/sql"
	model_req "friends_management/models/request"
	model_common "friends_management/models/respone"
	repo "friends_management/repository"
	"net/http"
	"regexp"
	"strconv"
)

type Service interface {
	AddFriend(RequestAddFriend model_req.AddFriendRequest) (*model_common.CommonRespone, *model_common.ResponseError)
	UnFriend(friendship_id string) (*model_common.CommonRespone, *model_common.ResponseError)
	ViewListFriendsByEmail(email string) (*model_common.ListFriendsRespone, *model_common.ResponseError)
	ViewListCommonFriendsByEmail(RequestAddFriend model_req.AddFriendRequest) (*model_common.ListFriendsRespone, *model_common.ResponseError)
}

// Manager is the implementation of recurring service
type Manager struct {
	dbconn *sql.DB
}

// NewManager initializes recurring service
func NewManager(dbconn *sql.DB) *Manager {
	return &Manager{
		dbconn: dbconn,
	}
}

// Add Friends Between Two User
func (m *Manager) AddFriend(RequestAddFriend model_req.AddFriendRequest) (*model_common.CommonRespone, *model_common.ResponseError) {

	if len(RequestAddFriend.Friends) < 2 {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "User One Or User Two Not Allow Empty!"}
	}

	UserOne := RequestAddFriend.Friends[0]
	UserTwo := RequestAddFriend.Friends[1]

	if m.checkIsValidEmail(UserOne) == false || m.checkIsValidEmail(UserTwo) == false {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Format user is not valid, User should be an email address!"}
	}

	commonRespone, err := repo.InsertFriendship(m.dbconn, UserOne, UserTwo)

	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusInternalServerError, Description: err.Error()}
	}

	return commonRespone, nil

}

// UnFriends Between Two User
func (m *Manager) UnFriend(friendship_id string) (*model_common.CommonRespone, *model_common.ResponseError) {

	id, err := strconv.Atoi(friendship_id)
	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Param not valid"}
	}

	commonRespone, err := repo.DeleteFriendship(m.dbconn, id)

	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: err.Error()}
	}

	return commonRespone, nil

}

//View list friends of an user by email
func (m *Manager) ViewListFriendsByEmail(mailAdress string) (*model_common.ListFriendsRespone, *model_common.ResponseError) {
	if m.checkIsValidEmail(mailAdress) == false {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Email not valid"}
	}

	listfriend, err := repo.GetListFriendByEmail(m.dbconn, mailAdress)

	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: err.Error()}
	}

	return listfriend, nil
}

//View list commmon friends of two user
func (m *Manager) ViewListCommonFriendsByEmail(RequestAddFriend model_req.AddFriendRequest) (*model_common.ListFriendsRespone, *model_common.ResponseError) {
	if len(RequestAddFriend.Friends) < 2 {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "User One Or User Two Not Allow Empty!"}
	}

	UserOne := RequestAddFriend.Friends[0]
	UserTwo := RequestAddFriend.Friends[1]

	if m.checkIsValidEmail(UserOne) == false || m.checkIsValidEmail(UserTwo) == false {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Format user is not valid, User should be an email address!"}
	}

	listCommonFriend, err := repo.GetCommonFriendsListByEmail(m.dbconn, UserOne, UserTwo)

	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: err.Error()}
	}

	return listCommonFriend, nil
}

func (m *Manager) checkIsValidEmail(mail string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(mail) < 3 && len(mail) > 254 {
		return false
	}
	return emailRegex.MatchString(mail)
}
