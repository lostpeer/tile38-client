package t38c

import (
	"fmt"
)

// Ping the server.
func (client *Client) Ping() error {
	var resp struct {
		Ping string `json:"ping"`
	}

	err := client.jExecute(&resp, "PING")
	if err != nil {
		return err
	}

	if resp.Ping != "pong" {
		return fmt.Errorf("bad ping response: %v", resp)
	}

	return nil
}
