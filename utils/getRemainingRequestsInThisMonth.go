package utils

func (c *Client) GetRemainingRequestsInThisMonth() int32 {
	return c.RemainingTimes
}
