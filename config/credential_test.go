package cargo

import (
	"strconv"
	"testing"
)

func TestCredential_GetActive(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	tests := []struct {
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			if got := c.GetActive(); got != tt.want {
				t.Errorf("GetActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredential_GetPassword(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	tests := []struct {
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			if got := c.GetPassword(); got != tt.want {
				t.Errorf("GetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredential_GetTitle(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	tests := []struct {
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			if got := c.GetTitle(); got != tt.want {
				t.Errorf("GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredential_GetUsername(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	tests := []struct {
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			if got := c.GetUsername(); got != tt.want {
				t.Errorf("GetUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredential_SetActive(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	type args struct {
		isActive bool
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			c.SetActive(tt.args.isActive)
		})
	}
}

func TestCredential_SetPassword(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	type args struct {
		password string
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			c.SetPassword(tt.args.password)
		})
	}
}

func TestCredential_SetTitle(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	type args struct {
		title string
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			c.SetTitle(tt.args.title)
		})
	}
}

func TestCredential_SetUsername(t *testing.T) {
	type fields struct {
		title    string
		username string
		password string
		isActive bool
	}
	type args struct {
		username string
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Credential{
				title:    tt.fields.title,
				username: tt.fields.username,
				password: tt.fields.password,
				isActive: tt.fields.isActive,
			}
			c.SetUsername(tt.args.username)
		})
	}
}
