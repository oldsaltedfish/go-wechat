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
