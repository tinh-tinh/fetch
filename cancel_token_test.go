package fetch_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tinh-tinh/fetch/v2"
)

func Test_CancelToken(t *testing.T) {
	// Create a cancel token
	token := fetch.NewCancelToken()

	instance := fetch.Create(&fetch.Config{
		BaseUrl:     "https://httpbin.org",
		CancelToken: token.Context(),
	})

	// Start the request
	go func() {
		resp := instance.Get("/delay/5")
		if resp.Error != nil {
			fmt.Println("HTTP error:", resp.Error)
		}
	}()

	// Cancel after 2 seconds
	time.Sleep(1 * time.Second)
	fmt.Println("Canceling...")
	token.Cancel()

	// Wait a bit to see result
	time.Sleep(3 * time.Second)
}
