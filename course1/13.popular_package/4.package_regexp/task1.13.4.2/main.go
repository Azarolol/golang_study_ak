package main

import (
	"fmt"
	"regexp"
)

type Ad struct {
	Title    string
	Описание string
}

func main() {
	ads := []Ad{
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
	}

	ads = censorAds(ads, map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	})

	for _, ad := range ads {
		fmt.Println(ad.Title)
		fmt.Println(ad.Описание)
		fmt.Println()
	}
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	newAds := make([]Ad, len(ads))
	copy(newAds, ads)
	for i := 0; i < len(ads); i++ {
		for k, v := range censor {
			re := regexp.MustCompile("(?i)" + k)
			newAds[i].Title = re.ReplaceAllString(newAds[i].Title, v)
			newAds[i].Описание = re.ReplaceAllString(newAds[i].Описание, v)
		}
	}
	return newAds
}
