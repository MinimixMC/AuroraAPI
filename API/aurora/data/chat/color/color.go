package color

// Not very useful right now

type Color struct {
	Name string
	Code string
	Hex  string
}

var (
	DARK_RED = &Color{
		Name: "dark_red",
		Code: "4",
		Hex:  "AA0000",
	}
	RED = &Color{
		Name: "red",
		Code: "c",
		Hex:  "FF5555",
	}
	GOLD = &Color{
		Name: "gold",
		Code: "6",
		Hex:  "FFAA00",
	}
	YELLOW = &Color{
		Name: "yellow",
		Code: "e",
		Hex:  "FFFF55",
	}
	DARK_GREEN = &Color{
		Name: "dark_green",
		Code: "2",
		Hex:  "00AA00",
	}
	GREEN = &Color{
		Name: "green",
		Code: "a",
		Hex:  "55FF55",
	}
	AQUA = &Color{
		Name: "aqua",
		Code: "b",
		Hex:  "55FFFF",
	}
	DARK_AQUA = &Color{
		Name: "dark_aqua",
		Code: "3",
		Hex:  "00AAAA",
	}
	DARK_BLUE = &Color{
		Name: "dark_blue",
		Code: "1",
		Hex:  "0000AA",
	}
	BLUE = &Color{
		Name: "blue",
		Code: "9",
		Hex:  "5555FF",
	}
	LIGHT_PURPLE = &Color{
		Name: "light_purple",
		Code: "d",
		Hex:  "FF55FF",
	}
	DARK_PURPLE = &Color{
		Name: "dark_purple",
		Code: "5",
		Hex:  "AA00AA",
	}
	WHITE = &Color{
		Name: "white",
		Code: "f",
		Hex:  "FFFFFF",
	}
	GRAY = &Color{
		Name: "gray",
		Code: "7",
		Hex:  "AAAAAA",
	}
	DARK_GRAY = &Color{
		Name: "dark_gray",
		Code: "8",
		Hex:  "555555",
	}
	BLACK = &Color{
		Name: "black",
		Code: "0",
		Hex:  "000000",
	}
)
