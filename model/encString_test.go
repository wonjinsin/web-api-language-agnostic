package model

import (
	"encoding/json"
	"testing"
)

func TestPassword_String(t *testing.T) {
	password := Password("mysecretpassword")
	expected := "**encrypted**"

	if got := password.String(); got != expected {
		t.Errorf("Password.String() = %v, want %v", got, expected)
	}
}

func TestPassword_MarshalJSON(t *testing.T) {
	password := Password("mysecretpassword")
	expected := `"**encrypted**"`

	got, err := json.Marshal(&password)
	if err != nil {
		t.Errorf("Password.MarshalJSON() error = %v", err)
		return
	}

	if string(got) != expected {
		t.Errorf("Password.MarshalJSON() = %v, want %v", string(got), expected)
	}
}

func TestPassword_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		want    Password
		wantErr bool
	}{
		{
			name:    "valid string",
			json:    `"mypassword"`,
			want:    Password("mypassword"),
			wantErr: false,
		},
		{
			name:    "empty string",
			json:    `""`,
			want:    Password(""),
			wantErr: false,
		},
		{
			name:    "invalid type (number)",
			json:    `123`,
			want:    Password(""),
			wantErr: true,
		},
		{
			name:    "invalid type (boolean)",
			json:    `true`,
			want:    Password(""),
			wantErr: true,
		},
		{
			name:    "invalid json",
			json:    `{"invalid"}`,
			want:    Password(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Password
			err := json.Unmarshal([]byte(tt.json), &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("Password.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("Password.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		password Password
		want     bool
	}{
		{
			name:     "empty password",
			password: Password(""),
			want:     true,
		},
		{
			name:     "non-empty password",
			password: Password("mypassword"),
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.password.IsEmpty(); got != tt.want {
				t.Errorf("Password.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
