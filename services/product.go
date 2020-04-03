package services

type Quantity string

const (
	Piece Quantity = "шт"
	KG    Quantity = "кг"
	Litre Quantity = "л"
	Empty Quantity = ""
)

func ParseQuantity(s string) Quantity {
	switch os := s; os {
	case "шт":
		return Piece
	case "кг":
		return KG
	case "л":
		return Litre
	default:
		return Empty
	}
}

type Product struct {
	Visible     bool
	Name        string
	Description string
	Quantity    Quantity
	Price       float64
	Image       string
}

type ViewData struct {
	Name  string
	Items []Product
}

type Category struct {
	Title   string
	Enabled bool
}

type CategoriesMap map[string]Category

var Categories = CategoriesMap{
	"backery": {
		Title:   "Пекарня",
		Enabled: true,
	},
	"bread": {
		Title:   "Бакалія",
		Enabled: true,
	},

	"zamorozheni": {
		Title:   "Заморожені",
		Enabled: false,
	},
	"souse": {
		Title:   "Соуси",
		Enabled: false,
	},
	"conserv": {
		Title:   "Консерви, Олія, Оцет",
		Enabled: false,
	},
	"sweet": {
		Title:   "Солодощі, Чіпси",
		Enabled: false,
	},
	"drink": {
		Title:   "Напої",
		Enabled: false,
	},
	"child": {
		Title:   "Товари для дітей",
		Enabled: false,
	},
}

func GetTitle(s string) string {
	var cat = Categories[s]
	return cat.Title
}

type ProductFetcher interface {
	Fetch(category string) ViewData
	mapper(rawData interface{}) []Product
}
