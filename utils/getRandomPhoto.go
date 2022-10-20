package utils

import (
	"math/rand"
	"time"
)

func (c *Client) GetRandomPhoto() (*Photo, error) {

	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1001)
	result, err := c.CuratedPhotos(1, randNum)
	if err == nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}
	return nil, err
}
