package main

import (
	"reflect"
	"testing"
)

func Test_filterWords(t *testing.T) {
	type args struct {
		text      string
		censorMap map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test filter words 1", args{"Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!",
			map[string]string{
				"крипта":   "фрукты",
				"крипту":   "фрукты",
				"крипты":   "фруктов",
				"биткоин":  "яблоки",
				"лайткоин": "яблоки",
				"эфир":     "яблоки",
			}}, "Внимание! Покупай срочно фрукты только у нас! Яблоки по низким ценам! Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!"},
		{"Test filter words 2", args{"",
			map[string]string{
				"крипта":   "фрукты",
				"крипту":   "фрукты",
				"крипты":   "фруктов",
				"биткоин":  "яблоки",
				"лайткоин": "яблоки",
				"эфир":     "яблоки",
			}}, "!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterWords(tt.args.text, tt.args.censorMap); got != tt.want {
				t.Errorf("filterWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordsToSentence(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test words to sentence 1", args{[]string{"Яблоки", "", "", "по", "низким", "ценам"}}, "Яблоки по низким ценам!"},
		{"Test words to sentence 2", args{[]string{""}}, "!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordsToSentence(tt.args.words); got != tt.want {
				t.Errorf("WordsToSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUpper(t *testing.T) {
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test check upper 1", args{"крипту", "фрукты"}, "фрукты"},
		{"Test check upper 2", args{"Биткоин", "яблоки"}, "Яблоки"},
		{"Test check upper 3", args{"Биткоин", ""}, ""},
		{"Test check upper 4", args{"", "яблоки"}, "яблоки"},
		{"Test check upper 5", args{"", ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUpper(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("CheckUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitSenctences(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Test split sentences 1", args{"Покупай срочно фрукты только у нас! Яблоки по низким ценам! Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!"},
			[]string{"Покупай срочно фрукты только у нас", " Яблоки по низким ценам", " Беги, успевай стать финансово независимым с помощью фруктов", " Фрукты будущее финансового мира"}},
		{"Test split sentences 2", args{"Внимание! Спасибо за внимание!"}, []string{"Внимание!    Спасибо за внимание!"}},
		{"Test split sentences 3", args{"Внимание. Спасибо за внимание."}, []string{"Внимание. Спасибо за внимание."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitSenctences(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSenctences() = %v, want %v", got, tt.want)
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
