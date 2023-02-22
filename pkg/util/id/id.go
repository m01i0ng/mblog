package id

import shortid "github.com/jasonsoft/go-short-id"

func GenShortID() string {
	options := shortid.Options{
		Number:        4,
		StartWithYear: true,
		EndWithHost:   false,
	}
	return toLower(shortid.Generate(options))
}

func toLower(ss string) string {
	var lower string
	for _, s := range ss {
		lower += string(s | ' ')
	}

	return lower
}
