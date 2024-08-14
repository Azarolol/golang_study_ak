package main

import "testing"

func Test_writeJSON(t *testing.T) {
	type args struct {
		filePath string
		data     interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test write JSON", args{"test/users.json", []map[string]interface{}{
			{
				"name": "Elliot",
				"age":  25,
			},
			{
				"name": "Fraser",
				"age":  30,
			},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeJSON(tt.args.filePath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("writeJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
