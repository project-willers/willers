package config_test

import (
	"testing"
	"willers-api/config"
)

func TestPort(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		want string
	}{
		{
			name: "Get Port Correct",
			want: "8080",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := config.Port(); got != tc.name {
				t.Errorf("Port() = %s, want = %s", got, tc.want)
			}
		})
	}
}

func TestDSN(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		want string
	}{
		{
			name: "Get DSN Correct",
			want: "test:test@tcp(localhost:3306)/test?parseTime=true&loc=Asia%2FTokyo&collation=utf8mb4_bin",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := config.DSN(); got != tc.want {
				t.Errorf("DSN() = %s, want = %s", got, tc.want)
			}
		})
	}
}
