package main

import "testing"

func Test_getYAML(t *testing.T) {
	type args struct {
		input []Config
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test get YAML", args{[]Config{
			{
				Server{"8080"},
				Db{"localhost", "5432", "admin", "password123"},
			},
		}}, "- server:\n    port: \"8080\"\n  db:\n    host: localhost\n    port: \"5432\"\n    user: admin\n    password: password123\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getYAML(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getYAML(): %v, want: %v", got, tt.want)
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
