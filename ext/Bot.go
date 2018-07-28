package ext

import (
	"encoding/json"
	"strconv"
	"net/url"
	"github.com/pkg/errors"
)

type Bot struct {
	Token string
	Id int
	FirstName string
	UserName string
}

func (b Bot) GetMe() (*User, error) {
	v := url.Values{}

	r, err := Get(b, "getMe", v)
	if err != nil {
		return nil, errors.Wrapf(err, "could not getMe")
	}
	if !r.Ok {
		return nil, errors.New("invalid getMe request")
	}

	var u User
	json.Unmarshal(r.Result, &u)
	return &u, nil
}

func (b Bot) GetUserProfilePhotos(userId int) (*UserProfilePhotos, error) {
	v := url.Values{}
	v.Add("user_id", strconv.Itoa(userId))


	r, err := Get(b, "getUserProfilePhotos", v)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get user profile photos")
	}
	if !r.Ok {
		return nil, errors.New("invalid getUserProfilePhotos request")
	}

	var userProfilePhotos UserProfilePhotos
	json.Unmarshal(r.Result, &userProfilePhotos)

	return &userProfilePhotos, nil
}


func (b Bot) GetFile(fileId string) (*File, error) {
	v := url.Values{}
	v.Add("file_id", fileId)

	r, err := Get(b, "getFile", v)
	if err != nil {
		return nil, errors.Wrapf(err, "could not complete getFile request")
	}
	if !r.Ok {
		return nil, errors.New("invalid getFile request")
	}

	var f File
	json.Unmarshal(r.Result, &f)
	return &f, nil
}

// TODO: options here
// TODO: r.OK or unmarshal??
func (b Bot) AnswerCallbackQuery(callbackQueryId string) (bool, error) {
	v := url.Values{}
	v.Add("callback_query_id", callbackQueryId)

	r, err := Get(b, "answerCallbackQuery", v)
	if err != nil {
		return false, errors.Wrapf(err, "could not complete getFile request")
	}
	if !r.Ok {
		return false, errors.New("invalid answerCallbackQuery request")
	}

	var bb bool
	json.Unmarshal(r.Result, &bb)
	return bb, nil
}
