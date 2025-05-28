package fetch_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch/v2"
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

	res := instance.Get("posts")
	require.Nil(t, res.Error)
	require.Equal(t, 200, res.Status)

	type QueryComment struct {
		PostID int `query:"postId"`
	}

	res = instance.Post("posts", &Post{
		UserID: 1,
		Title:  "foo",
		Body:   "bar",
	}, &QueryComment{
		PostID: 1,
	})
	require.Nil(t, res.Error)
	require.Equal(t, 201, res.Status)

	res = instance.Put("posts/1", &Post{
		UserID: 1,
		Title:  "foo",
		Body:   "bar",
	}, &QueryComment{
		PostID: 1,
	})
	require.Nil(t, res.Error)
	require.Equal(t, 200, res.Status)

	res = instance.Patch("posts/1", &Post{
		UserID: 1,
		Title:  "foo",
		Body:   "bar",
	}, &QueryComment{
		PostID: 1,
	})
	require.Nil(t, res.Error)
	require.Equal(t, 200, res.Status)

	res = instance.Delete("posts/1", &QueryComment{
		PostID: 1,
	})
	require.Nil(t, res.Error)
	require.Equal(t, 200, res.Status)
}

func Test_Timeout(t *testing.T) {
	instance := fetch.Create(&fetch.Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
		Timeout: 10 * time.Millisecond,
	})
	resp := instance.Get("comments")
	require.NotNil(t, resp.Error)
}

func Test_Cookies(t *testing.T) {
	instance := fetch.Create(&fetch.Config{
		BaseUrl:         "https://google.com",
		WithCredentials: true,
	})
	resp := instance.Get("")
	require.Nil(t, resp.Error)

	req, err := instance.GetConfig("GET", "", nil)
	require.Nil(t, err)
	require.NotEmpty(t, req.Cookies())
}
