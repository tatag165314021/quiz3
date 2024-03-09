package structs

type Segitiga struct {
	Alas   float64 `json:"alas"`
	Tinggi float64 `json:"tinggi"`
}

type Persegi struct {
	Sisi int `json:"sisi"`
}

type PersegiPanjang struct {
	Panjang int `json:"panjang"`
	Lebar   int `json:"lebar"`
}

type Lingkaran struct {
	JariJari float64 `json:"jariJari"`
}

type Category struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Books struct {
	ID          int64  `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageurl"`
	ReleaseYear int    `json:"releaseyear"`
	Price       string `json:"price"`
	TotalPage   int    `json:"totalpage"`
	Thickness   string `json:"thickness"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
	CategoryID  int    `json:"category_id"`
}
