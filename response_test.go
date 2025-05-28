package fetch_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch/v2"
)

func Test_Schema(t *testing.T) {
	instance := fetch.Create(&fetch.Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
	})

	type Comment struct {
		PostID int    `json:"postId"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}
	type Comments []Comment
	type QueryComment struct {
		PostID int `query:"postId"`
	}
	var comments Comments
	res := instance.Get("comments", &QueryComment{
		PostID: 1,
	}).Format(&comments)

	require.Nil(t, res.Error)
	require.Equal(t, 200, res.Status)
	if len(comments) > 0 {
		first := (comments)[0]
		require.Equal(t, 1, first.PostID)
	}

	type Post struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	res = instance.Get("comments").Format(&Post{})
	require.NotNil(t, res.Error)
}
