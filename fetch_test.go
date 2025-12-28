package fetch_test

import (
	"testing"
	"time"

	"github.com/sony/gobreaker/v2"
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

func Test_CircuitBreaker(t *testing.T) {
	instance := fetch.Create(&fetch.Config{
		BaseUrl: "https://jsonplaceholder.com",
		CBSettings: &gobreaker.Settings{
			Name: "CB",
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= 3 && failureRatio >= 0.6
			},
		},
		Timeout: 5 * time.Second,
	})

	var resp *fetch.Response
	// Intentionally cause failures
	for range 5 {
		resp = instance.Get("abc")
	}
	require.NotNil(t, resp.Error)
	require.Equal(t, resp.Error, gobreaker.ErrOpenState)

	instance2 := fetch.Create(&fetch.Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
		CBSettings: &gobreaker.Settings{
			Name: "CB2",
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= 3 && failureRatio >= 0.6
			},
		},
		Timeout: 5 * time.Second,
	})

	resp = instance2.Get("posts")
	require.Nil(t, resp.Error)
	require.Equal(t, 200, resp.Status)
}
