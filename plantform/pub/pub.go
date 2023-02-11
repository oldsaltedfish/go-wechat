package pub

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/oldsaltedfish/go-wechat/cache"
	"github.com/oldsaltedfish/go-wechat/plantform/pub/config"
	"sort"
	"time"
)

type Client struct {
	config *config.Config
	cache  cache.Cache
}

type Option func(d *Client)

func WithConfig(config *config.Config) Option {
	return func(c *Client) {
		c.config = config
	}
}

func WithCache(cache cache.Cache) Option {
	return func(c *Client) {
		c.cache = cache
	}
}

func NewClient(option ...Option) *Client {
	client := &Client{}
	for _, o := range option {
		o(client)
	}
	return client
}

func (c *Client) getAccessTokenKey() string {
	return c.config.Prefix + ":" + c.config.AppID + ":" + "AccessToken"
}

func (c *Client) GetAccessToken(ctx context.Context) (*AccessToken, error) {
	var (
		accessToken = new(AccessToken)
	)

	err := c.cache.Get(ctx, c.getAccessTokenKey(), accessToken)
	if err == nil &&
		accessToken.AccessToken != "" &&
		accessToken.ExpireTime > time.Now().Unix() {
		return accessToken, nil
	}
	if err != nil && err != cache.ErrCacheNil {
		return nil, err
	}
	accessToken, err = getAccessToken(c.config.AppID, c.config.AppSecret)
	if err != nil {
		return nil, err
	}
	accessTokenStr, err := json.Marshal(accessToken)
	if err != nil {
		return nil, err
	}
	err = c.cache.Save(ctx, c.getAccessTokenKey(), accessTokenStr)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (c *Client) CheckSignature(signature string, timestamp string, nonce string) bool {
	if c.config == nil {
		panic("wechat public account config is nil")
	}
	tmpArr := []string{signature, timestamp, nonce}
	sort.Slice(tmpArr, func(i, j int) bool {
		return i > j
	})
	var tmpStr string
	for _, item := range tmpArr {
		tmpStr += item
	}
	if signature == tmpStr {
		return true
	} else {
		return false
	}
}

func (c *Client) Receive(xmlData []byte) (*ReceiveMsg, error) {
	msg := new(ReceiveMsg)
	err := xml.Unmarshal(xmlData, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *Client) SendTxt(msg *SendMsg) ([]byte, error) {
	xmlData, err := xml.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return xmlData, nil
}
