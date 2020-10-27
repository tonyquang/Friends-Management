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
	AddFriend(requestAddFriend model_req.AddFriendRequest) (*model_common.CommonRespone, *model_common.ResponseError)
	UnFriend(friendship_id string) (*model_common.CommonRespone, *model_common.ResponseError)
	ViewListFriendsByEmail(email string) (*model_common.ListFriendsRespone, *model_common.ResponseError)
	ViewListCommonFriendsByEmail(requestViewCommonFriend model_req.AddFriendRequest) (*model_common.ListFriendsRespone, *model_common.ResponseError)
	SubscribeUpdate(requestUpdate model_req.HandleUpdateRequest) (*model_common.CommonRespone, *model_common.ResponseError)
	BlockUpdate(requestUpdate model_req.HandleUpdateRequest) (*model_common.CommonRespone, *model_common.ResponseError)
	ViewListFriendsRecvUpdate(email string) (*model_common.ListFriendsRecviceUpdateRespone, *model_common.ResponseError)
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
func (m *Manager) AddFriend(requestAddFriend model_req.AddFriendRequest) (*model_common.CommonRespone, *model_common.ResponseError) {

	if len(requestAddFriend.Friends) < 2 {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "User One Or User Two Not Allow Empty!"}
	}

	UserOne := requestAddFriend.Friends[0]
	UserTwo := requestAddFriend.Friends[1]

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
func (m *Manager) ViewListCommonFriendsByEmail(requestViewCommonFriend model_req.AddFriendRequest) (*model_common.ListFriendsRespone, *model_common.ResponseError) {
	if len(requestViewCommonFriend.Friends) < 2 {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "User One Or User Two Not Allow Empty!"}
	}

	UserOne := requestViewCommonFriend.Friends[0]
	UserTwo := requestViewCommonFriend.Friends[1]

	if m.checkIsValidEmail(UserOne) == false || m.checkIsValidEmail(UserTwo) == false {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Format user is not valid, User should be an email address!"}
	}

	listCommonFriend, err := repo.GetCommonFriendsListByEmail(m.dbconn, UserOne, UserTwo)

	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: err.Error()}
	}

	return listCommonFriend, nil
}

//subscribe to updates from an email address.
func (m *Manager) SubscribeUpdate(requestUpdate model_req.HandleUpdateRequest) (*model_common.CommonRespone, *model_common.ResponseError) {
	requestor := requestUpdate.Requestor
	target := requestUpdate.Target

	if m.checkIsValidEmail(requestor) == false || m.checkIsValidEmail(target) == false {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Format user is not valid, User should be an email address!"}
	}

	commonRes, err := repo.UpdateSubscribeStatusFriend(m.dbconn, requestor, target)
	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: err.Error()}
	}

	return commonRes, nil
}

//Block to updates from an email address.
func (m *Manager) BlockUpdate(requestUpdate model_req.HandleUpdateRequest) (*model_common.CommonRespone, *model_common.ResponseError) {
	requestor := requestUpdate.Requestor
	target := requestUpdate.Target

	if m.checkIsValidEmail(requestor) == false || m.checkIsValidEmail(target) == false {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Format user is not valid, User should be an email address!"}
	}

	commonRes, err := repo.BlockStatusFriend(m.dbconn, requestor, target)
	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: err.Error()}
	}

	return commonRes, nil
}

//retrieve all email addresses that can receive updates from an email address.
func (m *Manager) ViewListFriendsRecvUpdate(mail string) (*model_common.ListFriendsRecviceUpdateRespone, *model_common.ResponseError) {
	if m.checkIsValidEmail(mail) == false {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: "Format user is not email"}
	}

	listFriendsRecvUpdate, err := repo.GetAllEmailReceiveUpdate(m.dbconn, mail)

	if err != nil {
		return nil, &model_common.ResponseError{Code: http.StatusBadRequest, Description: err.Error()}
	}

	return listFriendsRecvUpdate, nil
}

func (m *Manager) checkIsValidEmail(mail string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(mail) < 3 && len(mail) > 254 {
		return false
	}
	return emailRegex.MatchString(mail)
}
