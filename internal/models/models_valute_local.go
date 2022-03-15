package models

const (
	ValuteUSD = "USD"
	ValuteEUR = "EUR"
	ValuteGBP = "GBP"
)

type Valute struct {
	Name  string `json:"Name"`
	Emoji string `json:"Emoji"`
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

func GetValuteItemName(valute string) string {
	return valuteMap[valute].Name
}

func GetValuteItemNameEmoji(valute string) string {
	return valute + " " + valuteMap[valute].Emoji
}
