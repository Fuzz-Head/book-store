package models

import (
	"regexp"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/lib/pq"

	"github.com/go-playground/validator/v10"
)

const (
	BookTypePaperback BookType = "paperback"
	BookTypeHardcover BookType = "hardcover"
	BookTypeEbook     BookType = "ebook"
)

// Publisher represents book publisher information
type Publisher struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" binding:"required"`
	Location  string     `json:"location"`
	Website   string     `json:"website,omitempty"`
	Founded   *int       `json:"founded,omitempty"` // Year founded
	IsActive  bool       `json:"is_active" grom:"default:true"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" grom:"index"`
	DeletedBy *string    `json:"deleted_by,omitempty"`
	CreatedAt time.Time  `json:"creted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Author represents book author information
type Author struct {
	ID          string     `json:"id" gorm:"primaryKey"`
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Biography   string     `json:"biography,omitempty"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
	DeathDate   *time.Time `json:"death_date,omitempty"`
	Nationality string     `json:"nationality,omitempty"`
	Website     string     `json:"website,omitempty"`
	IsActive    bool       `json:"is_active" grom:"default:true"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" grom:"index"`
	DeletedBy   *string    `json:"deleted_by,omitempty"`
	CreatedAt   time.Time  `json:"creted_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// BookAuthor represents the many-to-many relationship between books and authors
type BookAuthor struct {
	BookID    string     `json:"book_id" gorm:"primaryKey"`
	AuthorID  string     `json:"author_id" gorm:"primaryKey"`
	Role      string     `json:"role"` // "author", "co-author", "editor", "translator", etc.
	Book      Book       `json:"-" gorm:"foreignKey:BookID"`
	Author    Author     `json:"author" gorm:"foreignKey:AuthorID"`
	IsActive  bool       `json:"is_active" grom:"default:true"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" grom:"index"`
	DeletedBy *string    `json:"deleted_by,omitempty"`
	CreatedAt time.Time  `json:"creted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Category represents book categories/genres
type Category struct {
	ID          string  `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"` // For hierarchical categories
}

// Series represents book series information
type Series struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	TotalBooks  *int   `json:"total_books,omitempty"`
}

// Ratings represents book rating information
type Ratings struct {
	Average      float64        `json:"average"`
	Count        int            `json:"count"`
	Distribution map[string]int `json:"distribution,omitempty" gorm:"serializer:json"` // e.g., {"5": 100, "4": 50}
	Sources      []RatingSource `json:"sources,omitempty" gorm:"serializer:json"`
}

// RatingSource represents different rating sources
type RatingSource struct {
	Source  string  `json:"source"` // "goodreads", "amazon", "internal", etc.
	Rating  float64 `json:"rating"`
	Count   int     `json:"count"`
	MaxRate float64 `json:"max_rate"` // Maximum possible rating (5, 10, etc.)
}

// Dimensions represents physical book dimensions
type Dimensions struct {
	Length float64 `json:"length"` // in cm
	Width  float64 `json:"width"`  // in cm
	Height float64 `json:"height"` // in cm
	Weight float64 `json:"weight"` // in grams
}

// Pricing represents different pricing information
type Pricing struct {
	Currency        string             `json:"currency" binding:"required"`
	BasePrice       float64            `json:"base_price"`
	SalePrice       *float64           `json:"sale_price,omitempty"`
	Discounts       []Discount         `json:"discounts,omitempty" gorm:"serializer:json"`
	RegionalPricing map[string]float64 `json:"regional_pricing,omitempty" gorm:"serializer:json"`
}

// Discount represents discount information
type Discount struct {
	Type        string     `json:"type"` // "percentage", "fixed_amount"
	Value       float64    `json:"value"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Description string     `json:"description,omitempty"`
}

// Availability represents book availability information
type Availability struct {
	InStock          bool       `json:"in_stock"`
	StockCount       *int       `json:"stock_count,omitempty"`
	PreOrder         bool       `json:"pre_order"`
	ReleaseDate      *time.Time `json:"release_date,omitempty"`
	EstimatedRestock *time.Time `json:"estimated_restock,omitempty"`
	Discontinued     bool       `json:"discontinued"`
}

type BookType string

// Edition represents different editions of a book
type Edition struct {
	ID           string         `json:"id" gorm:"primaryKey"`
	BookID       string         `json:"book_id"`
	Type         BookType       `json:"type" binding:"required" validate:"booktype"` // "hardcover", "paperback", "ebook", "audiobook"
	Format       string         `json:"format,omitempty"`                            // "PDF", "EPUB", "MOBI", "MP3", etc.
	Language     string         `json:"language" binding:"required"`
	PageCount    *int           `json:"page_count,omitempty"`
	Dimensions   *Dimensions    `json:"dimensions,omitempty" gorm:"embedded;embeddedPrefix:dim_"`
	ISBN10       string         `json:"isbn10,omitempty"`
	ISBN13       string         `json:"isbn13,omitempty"`
	ASIN         string         `json:"asin,omitempty"` // Amazon Standard Identification Number
	EAN          string         `json:"ean,omitempty"`  // European Article Number
	Pricing      Pricing        `json:"pricing" gorm:"embedded;embeddedPrefix:pricing_"`
	Availability Availability   `json:"availability" gorm:"embedded;embeddedPrefix:avail_"`
	CoverImages  pq.StringArray `json:"cover_images,omitempty" gorm:"type:text[]"`
	SamplePages  pq.StringArray `json:"sample_pages,omitempty" gorm:"type:text[]"`
}

// Review represents user reviews
type Review struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	BookID    string    `json:"book_id"`
	UserID    string    `json:"user_id"`
	Rating    float64   `json:"rating" binding:"required,min=1,max=5"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Verified  bool      `json:"verified"` // Verified purchase
	Helpful   int       `json:"helpful"`  // Helpful votes count
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Award represents literary awards
type Award struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Category    string `json:"category,omitempty"`
	Year        int    `json:"year"`
	Winner      bool   `json:"winner"` // true if won, false if nominated
	Description string `json:"description,omitempty"`
}

// BookAward represents the many-to-many relationship between books and awards
type BookAward struct {
	BookID  string `json:"book_id" gorm:"primaryKey"`
	AwardID string `json:"award_id" gorm:"primaryKey"`
	Book    Book   `json:"-" gorm:"foreignKey:BookID"`
	Award   Award  `json:"award" gorm:"foreignKey:AwardID"`
}

// Enhanced main Book struct
type Book struct {
	ID               string `json:"id" gorm:"primaryKey"`
	Title            string `json:"title" binding:"required"`
	Subtitle         string `json:"subtitle,omitempty"`
	OriginalTitle    string `json:"original_title,omitempty"`
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`

	// Publication Information
	PublisherID         string     `json:"publisher_id"`
	Publisher           Publisher  `json:"publisher" gorm:"foreignKey:PublisherID"`
	PublicationDate     *time.Time `json:"publication_date,omitempty"`
	OriginalPublishDate *time.Time `json:"original_publish_date,omitempty"`
	Edition             string     `json:"edition,omitempty"` // "1st", "2nd", "Revised", etc.

	// Series Information
	SeriesID     *string  `json:"series_id,omitempty"`
	Series       *Series  `json:"series,omitempty" gorm:"foreignKey:SeriesID"`
	SeriesNumber *float64 `json:"series_number,omitempty"` // Float to handle 1.5, etc.

	// Content Information
	Language         string         `json:"language" binding:"required"`
	OriginalLanguage string         `json:"original_language,omitempty"`
	TranslatedFrom   string         `json:"translated_from,omitempty"`
	AgeRating        string         `json:"age_rating,omitempty"` // "G", "PG", "PG-13", "R", "Adult"
	ContentWarnings  pq.StringArray `json:"content_warnings,omitempty" gorm:"type:text[]"`
	Keywords         pq.StringArray `json:"keywords,omitempty" gorm:"type:text[]"`
	Tags             pq.StringArray `json:"tags,omitempty" gorm:"type:text[]"`

	// Classification
	Categories        []Category     `json:"categories,omitempty" gorm:"many2many:book_categories;"`
	Genres            pq.StringArray `json:"genres,omitempty" gorm:"type:text[]"`
	Subjects          pq.StringArray `json:"subjects,omitempty" gorm:"type:text[]"`
	DeweyDecimal      string         `json:"dewey_decimal,omitempty"`
	LibraryOfCongress string         `json:"library_of_congress,omitempty"`

	// Relationships
	Authors  []BookAuthor `json:"authors,omitempty" gorm:"foreignKey:BookID"`
	Editions []Edition    `json:"editions,omitempty" gorm:"foreignKey:BookID"`
	Awards   []BookAward  `json:"awards,omitempty" gorm:"foreignKey:BookID"`
	Reviews  []Review     `json:"reviews,omitempty" gorm:"foreignKey:BookID"`

	// Ratings and Reviews
	Ratings     Ratings `json:"ratings" gorm:"embedded;embeddedPrefix:ratings_"`
	ReviewCount int     `json:"review_count"`

	// Status and Metadata
	Status     string `json:"status"` // "published", "upcoming", "out_of_print", "discontinued"
	Featured   bool   `json:"featured"`
	Bestseller bool   `json:"bestseller"`
	NewRelease bool   `json:"new_release"`
	StaffPick  bool   `json:"staff_pick"`

	// SEO and Marketing
	SEOTitle       string         `json:"seo_title,omitempty"`
	SEODescription string         `json:"seo_description,omitempty"`
	MetaKeywords   pq.StringArray `json:"meta_keywords,omitempty" gorm:"type:text[]"`
	Slug           string         `json:"slug" gorm:"uniqueIndex"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	// External References
	ExternalIDs map[string]string `json:"external_ids,omitempty" gorm:"serializer:json"` // goodreads_id, amazon_id, etc.

	// Analytics (if needed)
	ViewCount     int `json:"view_count"`
	PurchaseCount int `json:"purchase_count"`
	WishlistCount int `json:"wishlist_count"`
}

// Additional helper structs for complex queries

// BookSearchResult for search API responses
type BookSearchResult struct {
	Book
	Relevance  float64             `json:"relevance,omitempty"`
	Highlights map[string][]string `json:"highlights,omitempty"`
}

// BookSummary for list endpoints
type BookSummary struct {
	ID              string     `json:"id"`
	Title           string     `json:"title"`
	Authors         []string   `json:"authors"` // Simplified author names
	CoverImage      string     `json:"cover_image,omitempty"`
	AverageRating   float64    `json:"average_rating"`
	Price           float64    `json:"price"`
	PublicationDate *time.Time `json:"publication_date,omitempty"`
}

// BookStats for analytics
type BookStats struct {
	BookID         string    `json:"book_id"`
	TotalViews     int       `json:"total_views"`
	TotalPurchases int       `json:"total_purchases"`
	TotalWishlists int       `json:"total_wishlists"`
	TotalReviews   int       `json:"total_reviews"`
	AverageRating  float64   `json:"average_rating"`
	LastUpdated    time.Time `json:"last_updated"`
}

// todo add to utils
func IsbnValidator(fl validator.FieldLevel) bool {
	isbn := fl.Field().String()
	regex := `^97[89][- ]?\d{1,5}[- ]?\d+[- ]?\d+[- ]?\d$`
	// regex := `^(?:ISBN(?:-13)?:? )?(?=[0-9]{13}$|(?=(?:[0-9]+[- ]){4})[- 0-9]{17}$)97[89][- ]?[0-9]{1,5}[- ]?[0-9]+[- ]?[0-9]+[- ]?[0-9]$`
	return regexp.MustCompile(regex).MatchString(isbn)
}

func (b *Book) Prepare() {
	// b.Title = strings.TrimSpace(b.Title)
	// b.Author = strings.TrimSpace(b.Author)
}

// ValidBookTypes contains all valid book types
var ValidBookTypes = []BookType{
	BookTypePaperback,
	BookTypeHardcover,
	BookTypeEbook,
}

// IsValid checks if the book type is valid
func (bt BookType) IsValid() bool {
	for _, validType := range ValidBookTypes {
		if bt == validType {
			return true
		}
	}
	return false
}

// Register custom validator with gin
func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("booktype", validBookType)
	}
}

// Custom validator function
func validBookType(fl validator.FieldLevel) bool {
	bookType := BookType(fl.Field().String())
	return bookType.IsValid()
}
