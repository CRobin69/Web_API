package book

import "encoding/json"

type BookInput struct {
	Title string      `json:"title" binding:"required"`        //binding json, required = wajib diisi
	Price json.Number `json:"price" binding:"required,number"` //ketentuan ketika input

}
