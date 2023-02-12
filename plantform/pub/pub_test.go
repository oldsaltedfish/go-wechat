package pub

import (
	"context"
	"fmt"
	"github.com/oldsaltedfish/go-wechat/cache"
	"github.com/oldsaltedfish/go-wechat/plantform/pub/config"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestClient_GetAccessToken(t *testing.T) {
	type fields struct {
		Config *config.Config
		Cache  cache.Cache
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"1",
			fields{
				Config: &config.Config{
					AppID:          "wx438b11a615731893",
					AppSecret:      "cd2349ea5060a8875d0b8f6b7c0150f6",
					Token:          "40DJcmTpwM21hFocS6cR",
					EncodingAESKey: "cF2RcQ9nVycnDKR4CNUtPMmiXUMdU0JMT3bVbXnjT1f",
				},
				Cache: cache.NewRedis(redis.NewClient(&redis.Options{
					Addr:     "101.43.21.215:6379",
					Password: "lkalsdjaoekm", // no password set
					DB:       1,              // use default DB
				})),
			},
			args{
				ctx: context.Background(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(WithConfig(tt.fields.Config), WithCache(tt.fields.Cache))
			got, err := c.GetAccessToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

//
//"/wechat/officialaccount?signature=add6dd4932eecad5f4f2aa78b4ee7ba0f1057c1f&echostr=8291736311866310623&timestamp=1676184542&nonce=208879378"
//"/wechat/officialaccount?signature=1b324d71c6c4fc43c40641edc8a93bbfa2e95a27&echostr=1628780342031613631&timestamp=1676185655&nonce=1610602073"
//

// signature=add6dd4932eecad5f4f2aa78b4ee7ba0f1057c1f&echostr=8291736311866310623&timestamp=1676184542&nonce=208879378
func TestClient_CheckSignature(t *testing.T) {
	type fields struct {
		config *config.Config
		cache  cache.Cache
	}
	type args struct {
		signature string
		timestamp string
		nonce     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"1",
			fields{
				config: &config.Config{
					AppID:          "wxd04d2cc131841294",
					AppSecret:      "8548dea0475e6b0427b9a2283a07e415",
					Token:          "40DJcmTpwM21hFocS6cR",
					EncodingAESKey: "cF2RcQ9nVycnDKR4CNUtPMmiXUMdU0JMT3bVbXnjT1f",
				},
				cache: cache.NewRedis(redis.NewClient(&redis.Options{
					Addr:     "101.43.21.215:6379",
					Password: "lkalsdjaoekm", // no password set
					DB:       1,              // use default DB
				})),
			},
			args{
				signature: "1b324d71c6c4fc43c40641edc8a93bbfa2e95a27",
				timestamp: "1676185655",
				nonce:     "1610602073",
			},
			true,
		},
		{
			"2",
			fields{
				config: &config.Config{
					AppID:          "wxd04d2cc131841294",
					AppSecret:      "8548dea0475e6b0427b9a2283a07e415",
					Token:          "40DJcmTpwM21hFocS6cR",
					EncodingAESKey: "cF2RcQ9nVycnDKR4CNUtPMmiXUMdU0JMT3bVbXnjT1f",
				},
				cache: cache.NewRedis(redis.NewClient(&redis.Options{
					Addr:     "101.43.21.215:6379",
					Password: "lkalsdjaoekm", // no password set
					DB:       1,              // use default DB
				})),
			},
			args{
				signature: "1b324d123171c6c4fc43c40641edc8a93bbfa2e95a27",
				timestamp: "16761856511235",
				nonce:     "1610602073",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config: tt.fields.config,
				cache:  tt.fields.cache,
			}
			if got := c.CheckSignature(tt.args.signature, tt.args.timestamp, tt.args.nonce); got != tt.want {
				t.Errorf("CheckSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
