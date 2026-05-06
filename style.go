package theme

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

// Defines the PurpleClay palette of color shades
var (
	S950 = lipgloss.Color("#07000f")
	S900 = lipgloss.Color("#130027")
	S800 = lipgloss.Color("#1e003f")
	S700 = lipgloss.Color("#280057")
	S600 = lipgloss.Color("#33006f")
	S500 = lipgloss.Color("#3d0087")
	S400 = lipgloss.Color("#48009f")
	S300 = lipgloss.Color("#5712ab")
	S200 = lipgloss.Color("#6725b7")
	S100 = lipgloss.Color("#906ccf")
	S50  = lipgloss.Color("#a980db")
)

// Defines the PurpleClay palette for accent color shades
var (
	Green900 = lipgloss.Color("#087331")
	Green700 = lipgloss.Color("#15803d")
	Amber900 = lipgloss.Color("#bf4102")
	Amber700 = lipgloss.Color("#db4b02")
	Red900   = lipgloss.Color("#ab0513")
	Red700   = lipgloss.Color("#db0f20")
)

var (
	// A defines a PurpleClay themed hyperlink that supports both light and dark terminals
	A = lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Foreground(compat.AdaptiveColor{
			Light: lipgloss.Color("S400"),
			Dark:  lipgloss.Color("S100"),
		})

	h = lipgloss.NewStyle().Padding(0, 1).Bold(true).Foreground(lipgloss.Color("#ffffff"))

	// H1 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H1 = h.
		Background(compat.AdaptiveColor{
			Light: lipgloss.Color("S50"),
			Dark:  lipgloss.Color("S200"),
		})

	// H2 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H2 = h.
		Background(compat.AdaptiveColor{
			Light: lipgloss.Color("S100"),
			Dark:  lipgloss.Color("S300"),
		})

	// H3 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H3 = h.
		Background(compat.AdaptiveColor{
			Light: lipgloss.Color("S200"),
			Dark:  lipgloss.Color("S400"),
		})

	// H4 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H4 = h.
		Background(compat.AdaptiveColor{
			Light: lipgloss.Color("S300"),
			Dark:  lipgloss.Color("S500"),
		})

	// H5 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H5 = h.
		Background(compat.AdaptiveColor{
			Light: lipgloss.Color("S400"),
			Dark:  lipgloss.Color("S600"),
		})

	// H6 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H6 = h.
		Background(compat.AdaptiveColor{
			Light: lipgloss.Color("S500"),
			Dark:  lipgloss.Color("S700"),
		})

	// Mark defines a PurpleClay themed text decoration for highlighting text. Supports
	// both light and dark terminals
	Mark = lipgloss.NewStyle().
		Padding(0, 1).
		Background(compat.AdaptiveColor{
			Light: lipgloss.Color("S50"),
			Dark:  lipgloss.Color("S700"),
		})

	// I defines a PurpleClay themed hyperlink that supports both light and dark terminals
	I = lipgloss.NewStyle().Italic(true)

	// U defines a PurpleClay themed underline text decoration
	U = lipgloss.NewStyle().Underline(true)

	// B defines a PurpleClay themed bold text decoration
	B = lipgloss.NewStyle().Bold(true)

	// S defines a PurpleClay themed strikethrough text decoration
	S = lipgloss.NewStyle().Strikethrough(true)

	// Tick defines a PurpleClay themed glyph ✓ that supports both light and dark terminals
	Tick = lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{
			Light: lipgloss.Color("Green900"),
			Dark:  lipgloss.Color("Green700"),
		}).
		Render("✓")

	// Cross defines a PurpleClay themed glyph x that supports both light and dark terminals
	Cross = lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{
			Light: lipgloss.Color("Red900"),
			Dark:  lipgloss.Color("Red700"),
		}).
		Render("✕")

	// Bang defines a PurpleClay themed glyph ! that supports both light and dark terminals
	Bang = lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{
			Light: lipgloss.Color("Amber900"),
			Dark:  lipgloss.Color("Amber700"),
		}).
		Render("!")

	/*
		// Logging defines a PurpleClay themed [logging] style that supports both light and
		// dark terminals
		//
		// [logging]: https://github.com/charmbracelet/log
		Logging = &log.Styles{
			Timestamp: lipgloss.NewStyle(),
			Caller:    lipgloss.NewStyle().Faint(true),
			Prefix:    lipgloss.NewStyle().Bold(true).Faint(true),
			Message:   lipgloss.NewStyle().MarginRight(2),
			Key: lipgloss.NewStyle().
				Foreground(compat.AdaptiveColor{
					Light: lipgloss.Color("S400"),
					Dark:  lipgloss.Color("S100"),
				}),
			Value:     lipgloss.NewStyle(),
			Separator: lipgloss.NewStyle().Faint(true),
			Levels: map[log.Level]lipgloss.Style{
				log.DebugLevel: lipgloss.NewStyle().
					SetString(Tick).
					Bold(true).
					MaxWidth(2),
				log.InfoLevel: lipgloss.NewStyle().
					SetString(Tick).
					Bold(true).
					MaxWidth(2),
				log.WarnLevel: lipgloss.NewStyle().
					SetString(Bang).
					Bold(true).
					MaxWidth(2),
				log.ErrorLevel: lipgloss.NewStyle().
					SetString(Cross).
					Bold(true).
					MaxWidth(2),
				log.FatalLevel: lipgloss.NewStyle().
					SetString(Cross).
					Bold(true).
					MaxWidth(2),
			},
			Keys: map[string]lipgloss.Style{
				"err": lipgloss.NewStyle().Foreground(compat.AdaptiveColor{
					Light: lipgloss.Color("Red900"),
					Dark:  lipgloss.Color("Red700"),
				}),
				"error": lipgloss.NewStyle().Foreground(compat.AdaptiveColor{
					Light: lipgloss.Color("Red900"),
					Dark:  lipgloss.Color("Red700"),
				}),
			},
			Values: map[string]lipgloss.Style{},
		}
	*/
)
