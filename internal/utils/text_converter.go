package converters

import (
	"strings"
)

type TextConverter struct{}

func NewTextConverter() *TextConverter {
	return &TextConverter{}
}

func (c *TextConverter) Convert(text string) string {
	// 将所有换行符转换为 HTML 换行符
	convertedText := strings.ReplaceAll(text, "\n", "<br>")
	// 将所有 URL 转换为 HTML 链接
	convertedText = ConvertUrlsToLinks(convertedText)
	return convertedText
}

func ConvertUrlsToLinks(text string) string {
	words := strings.Split(text, " ")
	for i, word := range words {
		if strings.HasPrefix(word, "http") {
			words[i] = "<a href='" + word + "'>" + word + "</a>"
		}
	}
	return strings.Join(words, " ")
}
