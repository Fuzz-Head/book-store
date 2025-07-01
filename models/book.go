package models

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Book struct {
	ID     string  `json:"id" grom:"primaryKey" binding:"required"`
	ISBN   string  `json:"isbn" binding:"required,isbn"`
	Title  string  `json:"title" binding:"required"`
	Author string  `json:"author" binding:"required"`
	Price  float64 `json:"price" binding:"required,gt=0"`
}

func IsbnValidator(fl validator.FieldLevel) bool {
	isbn := fl.Field().String()
	regex := `^(?:ISBN(?:-13)?:? )?(?=[0-9]{13}$|(?=(?:[0-9]+[- ]){4})[- 0-9]{17}$)97[89][- ]?[0-9]{1,5}[- ]?[0-9]+[- ]?[0-9]+[- ]?[0-9]$`
	return regexp.MustCompile(regex).MatchString(isbn)
}

func (b *Book) Prepare() {
	b.Title = strings.TrimSpace(b.Title)
	b.Author = strings.TrimSpace(b.Author)
}
