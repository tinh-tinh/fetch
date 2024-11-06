package fetch_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch"
)

func Test_Get(t *testing.T) {
	instance := fetch.Create(&fetch.Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
	})

	type Post struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	type Posts []Post
	res, err := instance.Schema(&Posts{}).Get("posts")
	require.Nil(t, err)
	require.Equal(t, 200, res.Status)

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
	res, err = instance.Schema(&Comments{}).Get("comments", &QueryComment{
		PostID: 1,
	})
	require.Nil(t, err)
	require.Equal(t, 200, res.Status)
	data := res.Data.(*Comments)
	if len(*data) > 0 {
		first := (*data)[0]
		require.Equal(t, 1, first.PostID)
	}

	res, err = instance.Post("posts", &Post{
		UserID: 1,
		Title:  "foo",
		Body:   "bar",
	})
	require.Nil(t, err)
	require.Equal(t, 201, res.Status)

	res, err = instance.Put("posts/1", &Post{
		UserID: 1,
		Title:  "foo",
		Body:   "bar",
	})
	require.Nil(t, err)
	require.Equal(t, 200, res.Status)

	res, err = instance.Patch("posts/1", &Post{
		UserID: 1,
		Title:  "foo",
		Body:   "bar",
	})
	require.Nil(t, err)
	require.Equal(t, 200, res.Status)

	res, err = instance.Delete("posts/1")
	require.Nil(t, err)
	require.Equal(t, 200, res.Status)
}

func Test_Timeout(t *testing.T) {
	instance := fetch.Create(&fetch.Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
		Timeout: 10 * time.Millisecond,
	})
	resp, err := instance.Get("comments")
	require.Nil(t, err)
	require.NotNil(t, resp.Error)
}
