package main

import (
	"reflect"
	"testing"
)

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

func Test_censorAds(t *testing.T) {
	type args struct {
		ads    []Ad
		censor map[string]string
	}
	tests := []struct {
		name string
		args args
		want []Ad
	}{
		{"Test censor ads", args{
			[]Ad{
				{
					Title:    "Куплю велосипед MeRiDa",
					Описание: "Куплю велосипед meriDA в хорошем состоянии.",
				},
				{
					Title:    "Продам ВаЗ 2101",
					Описание: "Продам ваз 2101 в хорошем состоянии.",
				},
				{
					Title:    "Продам БМВ",
					Описание: "Продам бМв в хорошем состоянии.",
				},
				{
					Title:    "Продам macBook pro",
					Описание: "Продам macBook PRO в хорошем состоянии.",
				},
			},
			map[string]string{
				"велосипед merida": "телефон Apple",
				"ваз":              "ВАЗ",
				"бмв":              "BMW",
				"macbook pro":      "Macbook Pro",
			},
		},
			[]Ad{
				{
					Title:    "Куплю телефон Apple",
					Описание: "Куплю телефон Apple в хорошем состоянии.",
				},
				{
					Title:    "Продам ВАЗ 2101",
					Описание: "Продам ВАЗ 2101 в хорошем состоянии.",
				},
				{
					Title:    "Продам BMW",
					Описание: "Продам BMW в хорошем состоянии.",
				},
				{
					Title:    "Продам Macbook Pro",
					Описание: "Продам Macbook Pro в хорошем состоянии.",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := censorAds(tt.args.ads, tt.args.censor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("censorAds() = %v, want %v", got, tt.want)
			}
		})
	}
}
