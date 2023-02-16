package chat

import "testing"

func TestInitChatClient(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"empty token", args{token: ""},
		},
		{
			"ok token", args{token: "abc"},
		},
		{
			"error token", args{token: "error."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitChatClient(tt.args.token)
		})
	}
}
