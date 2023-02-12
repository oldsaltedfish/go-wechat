package pub

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
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
	tmpArr := []string{c.config.Token, timestamp, nonce}
	sort.Slice(tmpArr, func(i, j int) bool {

		return tmpArr[i] < tmpArr[j]
	})
	var tmpStr string
	for _, item := range tmpArr {
		tmpStr += item
	}
	h := sha1.New()
	h.Write([]byte(tmpStr))
	newSignature := hex.EncodeToString(h.Sum(nil))
	if signature == newSignature {
		return true
	} else {
		return false
	}
}

func (c *Client) UnmarshallReceiveMsg(xmlData []byte) (*ReceiveMsg, error) {
	msg := new(ReceiveMsg)
	err := xml.Unmarshal(xmlData, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *Client) MarshallSendMsg(msg *SendMsg) ([]byte, error) {
	xmlData, err := xml.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return xmlData, nil
}
