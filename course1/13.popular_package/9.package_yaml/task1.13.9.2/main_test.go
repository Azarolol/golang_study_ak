package main

import (
	"reflect"
	"testing"
)

func Test_getConfigFromYAML(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{"Test get config from YAML 1", args{[]byte(`server: 
 port: "8080"
db:
 host: localhost
 port: "5432"
 user: admin
 password: password123`)}, Config{
			Server{"8080"},
			Db{"localhost", "5432", "admin", "password123"},
		}, false},
		{"Test get config from YAML 2", args{[]byte(`random text`)}, Config{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfigFromYAML(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfigFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfigFromYAML() = %v, want %v", got, tt.want)
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
