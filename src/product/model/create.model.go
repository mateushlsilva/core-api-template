package product_model

type Create struct {
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    Ingredients string  `json:"ingredients"`
}