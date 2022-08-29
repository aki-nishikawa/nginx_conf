package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {

	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatorURL)
	testChatUser := &chatUser{User: testUser}

	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatorURL {
		t.Error(
			"When no value set to avatar_url, authAvatar.GetAvatarURL() must return ErrNoAvatorURL\n",
			"Returned error: ", err,
		)
	}

	testUrl := "http://url-to-avatar"
	testUser = &gomniauthtest.TestUser{}
	testChatUser = &chatUser{User: testUser}
	testUser.On("AvatarURL").Return(testUrl, nil)

	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error(
			"When value set to avatar_url, authAvatar.GetAvatarURL() must not return any error\n",
			"Returned error: ", err,
		)
	}
	if url != testUrl {
		t.Error(
			"authAvatar.GetAvatarURL() returned invaliable url\n",
			"Returned url: ", url,
		)
	}

}

func TestGravatarAvatar(t *testing.T) {

	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}

	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("gravatarAvatar.GetAvatarURL() must not return any error")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Error("gravatarAvatar.GetAvatarURL() returned incorrect value: ", url)
	}

}

func TestFilesystemAvatar(t *testing.T) {

	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var filesystemAvatar FilesystemAvatar
	user := &chatUser{uniqueID: "abc"}

	url, err := filesystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("filesystemAvatar.GetAvatarURL() must not return any error")
	}
	if url != "/avatars/abc.jpg" {
		t.Error("filesystemAvatar.GetAvatarURL() returned incorrect url: ", url)
	}

}
