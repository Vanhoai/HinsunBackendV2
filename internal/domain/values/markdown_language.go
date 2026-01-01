package values

import "hinsun-backend/internal/core/failure"

type MarkdownLanguageCode string

const (
	Vietnamese MarkdownLanguageCode = "vi"
	English    MarkdownLanguageCode = "en"
	Chinese    MarkdownLanguageCode = "zh"
	Spanish    MarkdownLanguageCode = "es"
	French     MarkdownLanguageCode = "fr"
	German     MarkdownLanguageCode = "de"
	Japanese   MarkdownLanguageCode = "ja"
	Russian    MarkdownLanguageCode = "ru"
	Portuguese MarkdownLanguageCode = "pt"
	Italian    MarkdownLanguageCode = "it"
)

var SupportedMarkdownLanguages = []MarkdownLanguageCode{
	Vietnamese,
	English,
	Chinese,
	Spanish,
	French,
	German,
	Japanese,
	Russian,
	Portuguese,
	Italian,
}

func FromStringToMarkdownLanguageCode(lang string) (MarkdownLanguageCode, error) {
	for _, supportedLang := range SupportedMarkdownLanguages {
		if string(supportedLang) == lang {
			return supportedLang, nil
		}
	}

	return "", failure.NewValidationFailure("Unsupported markdown language code: " + lang)
}

func ConvertStringArrayToMarkdownLanguageCodes(languages []string) ([]MarkdownLanguageCode, error) {
	var result []MarkdownLanguageCode
	for _, str := range languages {
		langCode, err := FromStringToMarkdownLanguageCode(str)
		if err != nil {
			return nil, err
		}

		result = append(result, langCode)
	}

	return result, nil
}

func ConvertMarkdownLanguageCodesToStringArray(languages []MarkdownLanguageCode) []string {
	var result []string
	for _, lang := range languages {
		result = append(result, string(lang))
	}

	return result
}
