package cargo

import (
	"strconv"
	"testing"
)

func TestEndpoint_GetActive(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			if got := e.GetActive(); got != tt.want {
				t.Errorf("GetActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_GetActiveUrl(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
	}
	type args struct {
		mode Mode
	}
	tests := []struct {
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			if got := e.GetActiveUrl(tt.args.mode); got != tt.want {
				t.Errorf("GetActiveUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_GetDevelopment(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			if got := e.GetDevelopment(); got != tt.want {
				t.Errorf("GetDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_GetProduction(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			if got := e.GetProduction(); got != tt.want {
				t.Errorf("GetProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_GetTitle(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			if got := e.GetTitle(); got != tt.want {
				t.Errorf("GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_SetActive(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			e.SetActive(tt.args.isActive)
		})
	}
}

func TestEndpoint_SetDevelopment(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
	}
	type args struct {
		development string
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			e.SetDevelopment(tt.args.development)
		})
	}
}

func TestEndpoint_SetProduction(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
	}
	type args struct {
		production string
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			e.SetProduction(tt.args.production)
		})
	}
}

func TestEndpoint_SetTitle(t *testing.T) {
	type fields struct {
		title       string
		production  string
		development string
		isActive    bool
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
			e := &Endpoint{
				title:       tt.fields.title,
				production:  tt.fields.production,
				development: tt.fields.development,
				isActive:    tt.fields.isActive,
			}
			e.SetTitle(tt.args.title)
		})
	}
}
