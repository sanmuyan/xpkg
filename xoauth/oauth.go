package xoauth

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"golang.org/x/oauth2"
)

type UserInfo interface {
	GetEmail() string
}

type Client struct {
	config *oauth2.Config
}

func NewClient(config *oauth2.Config) *Client {
	return &Client{
		config: config,
	}
}

// GetAuthURL 构建授权 URL
func (c *Client) GetAuthURL(state string) string {
	return c.config.AuthCodeURL(state)
}

// GetToken 获取 token
func (c *Client) GetToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.config.Exchange(ctx, code)
}

// GetUserInfo 根据 token 获取用户信息并解析用户信息
func (c *Client) GetUserInfo(ctx context.Context, token *oauth2.Token, userInfoURL string, user UserInfo) error {
	client := c.config.Client(ctx, token)
	resp, err := client.Get(userInfoURL)
	if err != nil {
		return err
	}
	userRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(userRaw))
	}
	return json.Unmarshal(userRaw, user)
}
