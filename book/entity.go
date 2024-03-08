package book


type Book struct {
	ID    int    		`json:"id"`
	Title string 		`json:"title"`
	Description  string `json:"desc"`
	Price int    		`json:"price"`
	Rating int   		`json:"rating"`
	Discount int 		`json:"discount"`
	CreateAt string 	`json:"create_at"`
	UpdateAt string 	`json:"update_at"`
}
