package moderation

import (
	"net/http"
	"net/url"
	"testing"
)

func Test_moderationClient_ValidAuth(t *testing.T) {
	type fields struct {
		appid          string
		secretKey      string
		textBusinessId string
		conf           *config
		httpClient     *http.Client
	}
	type args struct {
		values url.Values
	}

	client := &moderationClient{
		appid:     "appId",
		secretKey: "secret",
	}

	v1 := client.Auth("businessId")

	values := url.Values{}
	values.Add("businessId", v1.Get("businessId"))
	values.Add("nonce", v1.Get("nonce"))
	values.Add("appId", v1.Get("appId"))
	values.Add("timestamp", v1.Get("timestamp"))
	values.Add("signature", v1.Get("signature"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test1", fields{secretKey: client.secretKey}, args{values: values}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &moderationClient{
				appid:          tt.fields.appid,
				secretKey:      tt.fields.secretKey,
				textBusinessId: tt.fields.textBusinessId,
				conf:           tt.fields.conf,
				httpClient:     tt.fields.httpClient,
			}

			if err := client.ValidAuth(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("moderationClient.ValidAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
