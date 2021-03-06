package replacer

import "strings"

func NewReplacer() *strings.Replacer {
	replacer := strings.NewReplacer(
		"А", "a", "Б", "b", "В", "v", "Г", "g",
		"Д", "d", "Е", "e", "Ж", "zh", "З", "z",
		"И", "i", "К", "k", "Л", "l", "М", "m",
		"Н", "n", "О", "o", "П", "p", "Р", "r",
		"С", "s", "Т", "t", "У", "u", "Ф", "f",
		"Х", "kh", "Ц", "ts", "Ч", "ch", "Ш", "sh",
		"Щ", "sch", "Э", "e", "Ю", "yu", "Я", "ya",
		"а", "a", "б", "b", "в", "v", "г", "g",
		"д", "d", "е", "e", "ё", "e", "ж", "zh",
		"з", "z", "ий", "y", "ия", "ia", "и", "i",
		"й", "y", "кс", "x", "к", "k", "л", "l",
		"м", "m", "н", "n", "о", "o", "п", "p",
		"р", "r", "с", "s", "т", "t", "у", "u",
		"ф", "f", "х", "kh", "ц", "ts", "ч", "ch",
		"ш", "sh", "щ", "sch", "ъ", "", "ы", "y",
		"ье", "ye", "ь", "", "э", "e", "ю", "yu",
		"я", "ya", " ", "_",
	)
	return replacer
}
