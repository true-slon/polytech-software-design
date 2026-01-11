package cr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
)

type Client struct {
	http    *http.Client
	baseURL string
	token   string
}

func NewClient(token string) *Client {
	return &Client{
		http:    &http.Client{Timeout: 5 * time.Second},
		baseURL: "https://api.clashroyale.com/v1",
		token:   token,
	}
}

func (c *Client) get(endpoint string, v interface{}) error {
	return retry.Do(
		func() error {
			return c.makeRequest(endpoint, v)
		},
		retry.Attempts(5),
		retry.Delay(1*time.Second),
	)
}

func (c *Client) makeRequest(endpoint string, v interface{}) error {
	req, err := http.NewRequest("GET", c.baseURL+endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("API error: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) GetPlayer(tag string) (*Player, error) {
	tag = strings.TrimPrefix(tag, "#")

	var player Player
	err := c.get("/players/%23"+tag, &player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (c *Client) GetClan(tag string) (*Clan, error) {
	tag = strings.TrimPrefix(tag, "#")

	var clan Clan
	err := c.get("/clans/%23"+tag, &clan)
	if err != nil {
		return nil, err
	}

	return &clan, nil
}

func (c *Client) GetBattleLog(tag string) (*BattleList, error) {
	tag = strings.TrimPrefix(tag, "#")

	var battleList BattleList
	err := c.get("/players/%23"+tag+"/battlelog", &battleList)
	if err != nil {
		return nil, err
	}

	return &battleList, nil
}
