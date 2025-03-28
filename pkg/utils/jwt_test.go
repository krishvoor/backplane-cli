package utils

import (
	"testing"
)

func TestGetFieldFromJWT(t *testing.T) {
	type testCase struct {
		name    string
		token   string
		field   string
		want    string
		wantErr bool
	}
	tests := []testCase{
		{
			name:  "Get string field",
			token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", // notsecret
			field: "sub",
			want:  "1234567890",
		},
		{
			name:    "Get number field",
			token:   "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjAsImV4cCI6MTcxNjY1MDA3MSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSJ9._CyJxncO4NBOH6a-Q_2oIVelCRZKJh9YiPBm4XEBZgI", // notsecret
			field:   "iat",
			wantErr: true,
		},
		{
			name:    "Get field that doesn't exist",
			token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", // notsecret
			field:   "foo",
			wantErr: true,
		},
		{
			name:    "Invalid token",
			token:   "abcdefg", // notsecret
			field:   "foo",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStringFieldFromJWT(tt.token, tt.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStringFieldFromJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStringFieldFromJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsernameFromJWT(t *testing.T) {
	type testCase struct {
		name  string
		token string
		want  string
	}
	tests := []testCase{
		{
			name:  "Get username",
			token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyZWRoYXQuY29tIiwiZXhwIjoxMTIwODI4MzQ0LCJ1c2VybmFtZSI6InRlc3R1c2VyIn0.2uBp-c/dIUtipUsnT1J6zjkJNVlIE640ZbuCvWevWRQ", // notsecret
			want:  "testuser",
		},
		{
			name:  "Get username when username field is missing",
			token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjAsImV4cCI6MTcxNjY1MDA3MSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSJ9._CyJxncO4NBOH6a-Q_2oIVelCRZKJh9YiPBm4XEBZgI", // notsecret
			want:  "anonymous",
		},
		{
			name:  "Invalid token",
			token: "abcdefg", // notsecret
			want:  "anonymous",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetUsernameFromJWT(tt.token)
			if got != tt.want {
				t.Errorf("GetUsernameFromJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO figure out this code and safely introduce sound error logic
// These test got added when we moved the function out of login.go to reuse in remediation cmd
// needs a careful revisit as to not break our flows
func TestGetContextNickname(t *testing.T) {
	type testCase struct {
		name        string
		namespace   string
		clusterNick string
		userNick    string
		want        string
	}
	tests := []testCase{
		{
			name:        "GetContextNickname",
			namespace:   "testNamespace",
			clusterNick: "testClusterNick",
			userNick:    "testUserNick",
			want:        "testNamespace" + "/" + "testClusterNick" + "/" + "testUserNick",
		},
		{
			name:        "GetContextNickname with empty userNick",
			namespace:   "testNamespace",
			clusterNick: "testClusterNick",
			userNick:    "",
			want:        "testNamespace" + "/" + "testClusterNick" + "/",
		},
		{
			name:        "GetContextNickname with empty userNick and empty clusterNick",
			namespace:   "testNamespace",
			clusterNick: "",
			userNick:    "",
			want:        "testNamespace" + "/" + "/",
		},
		{
			name:        "GetContextNickname with empty userNick and empty clusterNick and empty namespace",
			namespace:   "",
			clusterNick: "",
			userNick:    "",
			want:        "/" + "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetContextNickname(tt.namespace, tt.clusterNick, tt.userNick)
			if got != tt.want {
				t.Errorf("GetContextNickname() got = %v, want %v", got, tt.want)
			}
		})
	}
}
