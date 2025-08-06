package handlers

import (
	"time"

	"github.com/lib/pq"
)

//type CreateBookInput struct {
//	Title  string  `json:"title" binding:"required"`
//	Author string  `json:"author" binding:"required"`
//	Price  float64 `json:"price" binding:"required,gt=0"`
//	ISBN   string  `json:"isbn" binding:"required"`
//}
//
//type UpdateBookInput struct {
//	Title  string  `json:"title" binding:"required"`
//	Author string  `json:"author" binding:"required"`
//	Price  float64 `json:"price" binding:"required,gt=0"`
//	ISBN   string  `json:"isbn" binding:"required"`
//}

// Input structs for API requests
type CreateBookInput struct {
	Title               string            `json:"title" binding:"required"`
	Subtitle            string            `json:"subtitle,omitempty"`
	OriginalTitle       string            `json:"original_title,omitempty"`
	Description         string            `json:"description,omitempty"`
	ShortDescription    string            `json:"short_description,omitempty"`
	PublisherID         string            `json:"publisher_id,omitempty"`
	PublicationDate     *time.Time        `json:"publication_date,omitempty"`
	OriginalPublishDate *time.Time        `json:"original_publish_date,omitempty"`
	Edition             string            `json:"edition,omitempty"`
	SeriesID            *string           `json:"series_id,omitempty"`
	SeriesNumber        *float64          `json:"series_number,omitempty"`
	Language            string            `json:"language" binding:"required"`
	OriginalLanguage    string            `json:"original_language,omitempty"`
	TranslatedFrom      string            `json:"translated_from,omitempty"`
	AgeRating           string            `json:"age_rating,omitempty"`
	ContentWarnings     pq.StringArray    `json:"content_warnings,omitempty"`
	Keywords            pq.StringArray    `json:"keywords,omitempty"`
	Tags                pq.StringArray    `json:"tags,omitempty"`
	CategoryIDs         []string          `json:"category_ids,omitempty"`
	Genres              pq.StringArray    `json:"genres,omitempty"`
	Subjects            pq.StringArray    `json:"subjects,omitempty"`
	DeweyDecimal        string            `json:"dewey_decimal,omitempty"`
	LibraryOfCongress   string            `json:"library_of_congress,omitempty"`
	Status              string            `json:"status,omitempty"`
	Featured            bool              `json:"featured,omitempty"`
	Bestseller          bool              `json:"bestseller,omitempty"`
	NewRelease          bool              `json:"new_release,omitempty"`
	StaffPick           bool              `json:"staff_pick,omitempty"`
	SEOTitle            string            `json:"seo_title,omitempty"`
	SEODescription      string            `json:"seo_description,omitempty"`
	MetaKeywords        pq.StringArray    `json:"meta_keywords,omitempty"`
	Slug                string            `json:"slug,omitempty"`
	ExternalIDs         map[string]string `json:"external_ids,omitempty"`
	AuthorIDs           []string          `json:"author_ids,omitempty"`
}

type UpdateBookInput struct {
	Title               string            `json:"title" binding:"required"`
	Subtitle            string            `json:"subtitle,omitempty"`
	OriginalTitle       string            `json:"original_title,omitempty"`
	Description         string            `json:"description,omitempty"`
	ShortDescription    string            `json:"short_description,omitempty"`
	PublisherID         string            `json:"publisher_id,omitempty"`
	PublicationDate     *time.Time        `json:"publication_date,omitempty"`
	OriginalPublishDate *time.Time        `json:"original_publish_date,omitempty"`
	Edition             string            `json:"edition,omitempty"`
	SeriesID            *string           `json:"series_id,omitempty"`
	SeriesNumber        *float64          `json:"series_number,omitempty"`
	Language            string            `json:"language" binding:"required"`
	OriginalLanguage    string            `json:"original_language,omitempty"`
	TranslatedFrom      string            `json:"translated_from,omitempty"`
	AgeRating           string            `json:"age_rating,omitempty"`
	ContentWarnings     pq.StringArray    `json:"content_warnings,omitempty"`
	Keywords            pq.StringArray    `json:"keywords,omitempty"`
	Tags                pq.StringArray    `json:"tags,omitempty"`
	CategoryIDs         []string          `json:"category_ids,omitempty"`
	Genres              pq.StringArray    `json:"genres,omitempty"`
	Subjects            pq.StringArray    `json:"subjects,omitempty"`
	DeweyDecimal        string            `json:"dewey_decimal,omitempty"`
	LibraryOfCongress   string            `json:"library_of_congress,omitempty"`
	Status              string            `json:"status,omitempty"`
	Featured            bool              `json:"featured,omitempty"`
	Bestseller          bool              `json:"bestseller,omitempty"`
	NewRelease          bool              `json:"new_release,omitempty"`
	StaffPick           bool              `json:"staff_pick,omitempty"`
	SEOTitle            string            `json:"seo_title,omitempty"`
	SEODescription      string            `json:"seo_description,omitempty"`
	MetaKeywords        pq.StringArray    `json:"meta_keywords,omitempty"`
	Slug                string            `json:"slug,omitempty"`
	ExternalIDs         map[string]string `json:"external_ids,omitempty"`
	AuthorIDs           []string          `json:"author_ids,omitempty"`
	UpdatedAt           *time.Time        `json:"updated_at"`
}
