package database

import (
	"log"

	"github.com/Fuzz-Head/models"
	"github.com/google/uuid"
)

func SeedBooks() {
	var count int64
	if err := DB.Model(&models.Book{}).Count(&count).Error; err != nil {
		log.Printf("Failed to check book count: %v", err)
		return
	}

	if count > 0 {
		log.Println("Books already seeded — skipping seeding.")
		return
	}

	log.Println("Seeding initial books...")

	books := []models.Book{
		{ID: uuid.New().String(), Title: "1984", Author: "George Orwell", Price: 9.99, ISBN: "9780452284234"},
		{ID: uuid.New().String(), Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 10.99, ISBN: "9780061120084"},
		{ID: uuid.New().String(), Title: "Pride and Prejudice", Author: "Jane Austen", Price: 8.95, ISBN: "9780141439518"},
		{ID: uuid.New().String(), Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Price: 9.49, ISBN: "9780743273565"},
		{ID: uuid.New().String(), Title: "The Catcher in the Rye", Author: "J.D. Salinger", Price: 9.79, ISBN: "9780316769488"},
		{ID: uuid.New().String(), Title: "The Hobbit", Author: "J.R.R. Tolkien", Price: 12.99, ISBN: "9780547928227"},
		{ID: uuid.New().String(), Title: "Fahrenheit 451", Author: "Ray Bradbury", Price: 9.89, ISBN: "9781451673319"},
		{ID: uuid.New().String(), Title: "Jane Eyre", Author: "Charlotte Brontë", Price: 11.50, ISBN: "9780141441146"},
		{ID: uuid.New().String(), Title: "Brave New World", Author: "Aldous Huxley", Price: 10.25, ISBN: "9780060850524"},
		{ID: uuid.New().String(), Title: "The Alchemist", Author: "Paulo Coelho", Price: 10.99, ISBN: "9780061122415"},
		{ID: uuid.New().String(), Title: "The Book Thief", Author: "Markus Zusak", Price: 11.95, ISBN: "9780375842207"},
		{ID: uuid.New().String(), Title: "The Road", Author: "Cormac McCarthy", Price: 9.50, ISBN: "9780307387899"},
		{ID: uuid.New().String(), Title: "Sapiens", Author: "Yuval Noah Harari", Price: 14.99, ISBN: "9780062316097"},
		{ID: uuid.New().String(), Title: "Educated", Author: "Tara Westover", Price: 13.95, ISBN: "9780399590504"},
		{ID: uuid.New().String(), Title: "The Silent Patient", Author: "Alex Michaelides", Price: 12.00, ISBN: "9781250301697"},
		{ID: uuid.New().String(), Title: "Where the Crawdads Sing", Author: "Delia Owens", Price: 13.45, ISBN: "9780735219106"},
		{ID: uuid.New().String(), Title: "Atomic Habits", Author: "James Clear", Price: 11.79, ISBN: "9780735211292"},
		{ID: uuid.New().String(), Title: "The Midnight Library", Author: "Matt Haig", Price: 12.49, ISBN: "9780525559474"},
		{ID: uuid.New().String(), Title: "The Subtle Art of Not Giving a F*ck", Author: "Mark Manson", Price: 10.59, ISBN: "9780062457714"},
		{ID: uuid.New().String(), Title: "Becoming", Author: "Michelle Obama", Price: 15.99, ISBN: "9781524763138"},
	}

	for _, book := range books {
		book.Prepare()
		if err := DB.Create(&book).Error; err != nil {
			log.Printf("Failed to seed book: %s - %v", book.Title, err)
		}
	}

	log.Println("Book seeding complete.")
}
