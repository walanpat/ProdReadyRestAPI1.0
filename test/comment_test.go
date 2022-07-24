package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComments(t *testing.T) {
	fmt.Println("running e2e test for health check endpoint")
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "TESTINGAUTHOR","body":"Hello Post Test"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

//func TestGetComment(t *testing.T) {
//	client := resty.New()
//	resp, err := client.R().
//		SetBody(`{"slug:" "/", "author:" "TESTINGAUTHOR","body:""Hello Post Test"`).
//		Post(BASE_URL + "/api/comment")
//	assert.NoError(t, err)
//	assert.Equal(t, 200, resp.StatusCode())
//}
