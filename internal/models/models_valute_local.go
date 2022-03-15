package models

const (
	ValuteUSD = "USD"
	ValuteEUR = "EUR"
	ValuteGBP = "GBP"
)

type Valute struct {
	CharCode string `json:"CharCode"`
	Name     string `json:"Name"`
	Emoji    string `json:"Emoji"`
}

var valuteMap = map[string]Valute{
	ValuteUSD: {
		Name:  "Доллара США",
		Emoji: "🇺🇸",
	},
	ValuteEUR: {
		Name:  "Евро",
		Emoji: "🇪🇺",
	},
	ValuteGBP: {
		Name:  "Фунта стерлингов",
		Emoji: "🇬🇧",
	},
}

func GetValuteItem(valute string) Valute {
	return valuteMap[valute]
}

func GetValuteItemFullName(valute string) string {
	item := valuteMap[valute]

	return item.Name + " " + item.Emoji
}

func GetValuteItemShortName(valute string) string {
	return valute + " " + valuteMap[valute].Emoji
}
