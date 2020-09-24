package markdown

import (
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
)

func Render(text, style string) (string, error) {
	// Glamour rendering preserves carriage return characters in code blocks, but
	// we need to ensure that no such characters are present in the output.
	text = strings.ReplaceAll(text, "\r\n", "\n")

	tr, err := glamour.NewTermRenderer(
		glamour.WithStylePath(style),
		// glamour.WithBaseURL(""),  // TODO: make configurable
		// glamour.WithWordWrap(80), // TODO: make configurable
	)
	if err != nil {
		return "", err
	}

	return tr.Render(text)
}

func GetStyle(terminalTheme string) string {
	style := os.Getenv("GLAMOUR_STYLE")
	if style != "" && style != "auto" {
		return style
	}

	if terminalTheme == "light" || terminalTheme == "dark" {
		return terminalTheme
	}

	return "notty"
}
