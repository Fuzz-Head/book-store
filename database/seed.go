package database

import (
	"log"

	"github.com/Fuzz-Head/domain/models"
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
		{ID: uuid.New().String(), Title: "The Road Vol 1", Author: "Cormac McCarthy", Price: 9.71, ISBN: "978000000001"},
		{ID: uuid.New().String(), Title: "Sapiens Vol 2", Author: "Yuval Noah Harari", Price: 13.82, ISBN: "978000000002"},
		{ID: uuid.New().String(), Title: "Educated Vol 3", Author: "Tara Westover", Price: 12.47, ISBN: "978000000003"},
		{ID: uuid.New().String(), Title: "The Silent Patient Vol 4", Author: "Alex Michaelides", Price: 11.25, ISBN: "978000000004"},
		{ID: uuid.New().String(), Title: "Atomic Habits Vol 5", Author: "James Clear", Price: 11.68, ISBN: "978000000005"},
		{ID: uuid.New().String(), Title: "The Midnight Library Vol 6", Author: "Matt Haig", Price: 13.22, ISBN: "978000000006"},
		{ID: uuid.New().String(), Title: "Jane Eyre Vol 7", Author: "Charlotte Brontë", Price: 10.34, ISBN: "978000000007"},
		{ID: uuid.New().String(), Title: "The Great Gatsby Vol 8", Author: "F. Scott Fitzgerald", Price: 9.63, ISBN: "978000000008"},
		{ID: uuid.New().String(), Title: "Pride and Prejudice Vol 9", Author: "Jane Austen", Price: 9.12, ISBN: "978000000009"},
		{ID: uuid.New().String(), Title: "Fahrenheit 451 Vol 10", Author: "Ray Bradbury", Price: 8.91, ISBN: "978000000010"},
		{ID: uuid.New().String(), Title: "1984 Vol 11", Author: "George Orwell", Price: 10.74, ISBN: "978000000011"},
		{ID: uuid.New().String(), Title: "The Hobbit Vol 12", Author: "J.R.R. Tolkien", Price: 14.55, ISBN: "978000000012"},
		{ID: uuid.New().String(), Title: "Brave New World Vol 13", Author: "Aldous Huxley", Price: 10.07, ISBN: "978000000013"},
		{ID: uuid.New().String(), Title: "The Alchemist Vol 14", Author: "Paulo Coelho", Price: 9.99, ISBN: "978000000014"},
		{ID: uuid.New().String(), Title: "The Book Thief Vol 15", Author: "Markus Zusak", Price: 11.66, ISBN: "978000000015"},
		{ID: uuid.New().String(), Title: "Where the Crawdads Sing Vol 16", Author: "Delia Owens", Price: 12.89, ISBN: "978000000016"},
		{ID: uuid.New().String(), Title: "Becoming Vol 17", Author: "Michelle Obama", Price: 14.98, ISBN: "978000000017"},
		{ID: uuid.New().String(), Title: "Sapiens Vol 18", Author: "Yuval Noah Harari", Price: 13.77, ISBN: "978000000018"},
		{ID: uuid.New().String(), Title: "Educated Vol 19", Author: "Tara Westover", Price: 13.17, ISBN: "978000000019"},
		{ID: uuid.New().String(), Title: "The Silent Patient Vol 20", Author: "Alex Michaelides", Price: 12.36, ISBN: "978000000020"},
		{ID: uuid.New().String(), Title: "Atomic Habits Vol 21", Author: "James Clear", Price: 11.98, ISBN: "978000000021"},
		{ID: uuid.New().String(), Title: "The Midnight Library Vol 22", Author: "Matt Haig", Price: 12.88, ISBN: "978000000022"},
		{ID: uuid.New().String(), Title: "Jane Eyre Vol 23", Author: "Charlotte Brontë", Price: 10.55, ISBN: "978000000023"},
		{ID: uuid.New().String(), Title: "The Great Gatsby Vol 24", Author: "F. Scott Fitzgerald", Price: 9.80, ISBN: "978000000024"},
		{ID: uuid.New().String(), Title: "Pride and Prejudice Vol 25", Author: "Jane Austen", Price: 8.75, ISBN: "978000000025"},
		{ID: uuid.New().String(), Title: "Fahrenheit 451 Vol 26", Author: "Ray Bradbury", Price: 9.40, ISBN: "978000000026"},
		{ID: uuid.New().String(), Title: "1984 Vol 27", Author: "George Orwell", Price: 10.11, ISBN: "978000000027"},
		{ID: uuid.New().String(), Title: "The Hobbit Vol 28", Author: "J.R.R. Tolkien", Price: 13.10, ISBN: "978000000028"},
		{ID: uuid.New().String(), Title: "Brave New World Vol 29", Author: "Aldous Huxley", Price: 10.03, ISBN: "978000000029"},
		{ID: uuid.New().String(), Title: "The Alchemist Vol 30", Author: "Paulo Coelho", Price: 9.50, ISBN: "978000000030"},
		{ID: uuid.New().String(), Title: "The Book Thief Vol 31", Author: "Markus Zusak", Price: 12.40, ISBN: "978000000031"},
		{ID: uuid.New().String(), Title: "Where the Crawdads Sing Vol 32", Author: "Delia Owens", Price: 11.67, ISBN: "978000000032"},
		{ID: uuid.New().String(), Title: "Becoming Vol 33", Author: "Michelle Obama", Price: 14.19, ISBN: "978000000033"},
		{ID: uuid.New().String(), Title: "Sapiens Vol 34", Author: "Yuval Noah Harari", Price: 13.28, ISBN: "978000000034"},
		{ID: uuid.New().String(), Title: "Educated Vol 35", Author: "Tara Westover", Price: 12.88, ISBN: "978000000035"},
		{ID: uuid.New().String(), Title: "The Silent Patient Vol 36", Author: "Alex Michaelides", Price: 12.60, ISBN: "978000000036"},
		{ID: uuid.New().String(), Title: "Atomic Habits Vol 37", Author: "James Clear", Price: 11.55, ISBN: "978000000037"},
		{ID: uuid.New().String(), Title: "The Midnight Library Vol 38", Author: "Matt Haig", Price: 13.44, ISBN: "978000000038"},
		{ID: uuid.New().String(), Title: "Jane Eyre Vol 39", Author: "Charlotte Brontë", Price: 10.79, ISBN: "978000000039"},
		{ID: uuid.New().String(), Title: "The Great Gatsby Vol 40", Author: "F. Scott Fitzgerald", Price: 9.94, ISBN: "978000000040"},
		{ID: uuid.New().String(), Title: "Pride and Prejudice Vol 41", Author: "Jane Austen", Price: 8.66, ISBN: "978000000041"},
		{ID: uuid.New().String(), Title: "Fahrenheit 451 Vol 42", Author: "Ray Bradbury", Price: 9.27, ISBN: "978000000042"},
		{ID: uuid.New().String(), Title: "1984 Vol 43", Author: "George Orwell", Price: 10.63, ISBN: "978000000043"},
		{ID: uuid.New().String(), Title: "The Hobbit Vol 44", Author: "J.R.R. Tolkien", Price: 14.44, ISBN: "978000000044"},
		{ID: uuid.New().String(), Title: "Brave New World Vol 45", Author: "Aldous Huxley", Price: 9.93, ISBN: "978000000045"},
		{ID: uuid.New().String(), Title: "The Alchemist Vol 46", Author: "Paulo Coelho", Price: 10.02, ISBN: "978000000046"},
		{ID: uuid.New().String(), Title: "The Book Thief Vol 47", Author: "Markus Zusak", Price: 12.67, ISBN: "978000000047"},
		{ID: uuid.New().String(), Title: "Where the Crawdads Sing Vol 48", Author: "Delia Owens", Price: 13.21, ISBN: "978000000048"},
		{ID: uuid.New().String(), Title: "Becoming Vol 49", Author: "Michelle Obama", Price: 15.11, ISBN: "978000000049"},
		{ID: uuid.New().String(), Title: "Sapiens Vol 50", Author: "Yuval Noah Harari", Price: 13.55, ISBN: "978000000050"},
		{ID: uuid.New().String(), Title: "Educated Vol 51", Author: "Tara Westover", Price: 12.67, ISBN: "978000000051"},
		{ID: uuid.New().String(), Title: "The Silent Patient Vol 52", Author: "Alex Michaelides", Price: 11.92, ISBN: "978000000052"},
		{ID: uuid.New().String(), Title: "Atomic Habits Vol 53", Author: "James Clear", Price: 11.43, ISBN: "978000000053"},
		{ID: uuid.New().String(), Title: "The Midnight Library Vol 54", Author: "Matt Haig", Price: 12.72, ISBN: "978000000054"},
		{ID: uuid.New().String(), Title: "Jane Eyre Vol 55", Author: "Charlotte Brontë", Price: 10.01, ISBN: "978000000055"},
		{ID: uuid.New().String(), Title: "The Great Gatsby Vol 56", Author: "F. Scott Fitzgerald", Price: 9.22, ISBN: "978000000056"},
		{ID: uuid.New().String(), Title: "Pride and Prejudice Vol 57", Author: "Jane Austen", Price: 8.41, ISBN: "978000000057"},
		{ID: uuid.New().String(), Title: "Fahrenheit 451 Vol 58", Author: "Ray Bradbury", Price: 9.38, ISBN: "978000000058"},
		{ID: uuid.New().String(), Title: "1984 Vol 59", Author: "George Orwell", Price: 10.38, ISBN: "978000000059"},
		{ID: uuid.New().String(), Title: "The Hobbit Vol 60", Author: "J.R.R. Tolkien", Price: 13.77, ISBN: "978000000060"},
		{ID: uuid.New().String(), Title: "Brave New World Vol 61", Author: "Aldous Huxley", Price: 10.12, ISBN: "978000000061"},
		{ID: uuid.New().String(), Title: "The Alchemist Vol 62", Author: "Paulo Coelho", Price: 10.88, ISBN: "978000000062"},
		{ID: uuid.New().String(), Title: "The Book Thief Vol 63", Author: "Markus Zusak", Price: 12.35, ISBN: "978000000063"},
		{ID: uuid.New().String(), Title: "Where the Crawdads Sing Vol 64", Author: "Delia Owens", Price: 11.89, ISBN: "978000000064"},
		{ID: uuid.New().String(), Title: "Becoming Vol 65", Author: "Michelle Obama", Price: 15.66, ISBN: "978000000065"},
		{ID: uuid.New().String(), Title: "Sapiens Vol 66", Author: "Yuval Noah Harari", Price: 13.84, ISBN: "978000000066"},
		{ID: uuid.New().String(), Title: "Educated Vol 67", Author: "Tara Westover", Price: 13.18, ISBN: "978000000067"},
		{ID: uuid.New().String(), Title: "The Silent Patient Vol 68", Author: "Alex Michaelides", Price: 11.76, ISBN: "978000000068"},
		{ID: uuid.New().String(), Title: "Atomic Habits Vol 69", Author: "James Clear", Price: 11.89, ISBN: "978000000069"},
		{ID: uuid.New().String(), Title: "The Midnight Library Vol 70", Author: "Matt Haig", Price: 12.53, ISBN: "978000000070"},
		{ID: uuid.New().String(), Title: "Jane Eyre Vol 71", Author: "Charlotte Brontë", Price: 10.65, ISBN: "978000000071"},
		{ID: uuid.New().String(), Title: "The Great Gatsby Vol 72", Author: "F. Scott Fitzgerald", Price: 9.75, ISBN: "978000000072"},
		{ID: uuid.New().String(), Title: "Pride and Prejudice Vol 73", Author: "Jane Austen", Price: 8.98, ISBN: "978000000073"},
		{ID: uuid.New().String(), Title: "Fahrenheit 451 Vol 74", Author: "Ray Bradbury", Price: 9.91, ISBN: "978000000074"},
		{ID: uuid.New().String(), Title: "1984 Vol 75", Author: "George Orwell", Price: 10.92, ISBN: "978000000075"},
		{ID: uuid.New().String(), Title: "The Hobbit Vol 76", Author: "J.R.R. Tolkien", Price: 13.32, ISBN: "978000000076"},
		{ID: uuid.New().String(), Title: "Brave New World Vol 77", Author: "Aldous Huxley", Price: 10.26, ISBN: "978000000077"},
		{ID: uuid.New().String(), Title: "The Alchemist Vol 78", Author: "Paulo Coelho", Price: 10.59, ISBN: "978000000078"},
		{ID: uuid.New().String(), Title: "The Book Thief Vol 79", Author: "Markus Zusak", Price: 12.48, ISBN: "978000000079"},
		{ID: uuid.New().String(), Title: "Where the Crawdads Sing Vol 80", Author: "Delia Owens", Price: 11.92, ISBN: "978000000080"},
		{ID: uuid.New().String(), Title: "Becoming Vol 81", Author: "Michelle Obama", Price: 14.85, ISBN: "978000000081"},
		{ID: uuid.New().String(), Title: "Sapiens Vol 82", Author: "Yuval Noah Harari", Price: 13.44, ISBN: "978000000082"},
		{ID: uuid.New().String(), Title: "Educated Vol 83", Author: "Tara Westover", Price: 13.09, ISBN: "978000000083"},
		{ID: uuid.New().String(), Title: "The Silent Patient Vol 84", Author: "Alex Michaelides", Price: 12.24, ISBN: "978000000084"},
		{ID: uuid.New().String(), Title: "Atomic Habits Vol 85", Author: "James Clear", Price: 11.37, ISBN: "978000000085"},
		{ID: uuid.New().String(), Title: "The Midnight Library Vol 86", Author: "Matt Haig", Price: 13.13, ISBN: "978000000086"},
		{ID: uuid.New().String(), Title: "Jane Eyre Vol 87", Author: "Charlotte Brontë", Price: 10.88, ISBN: "978000000087"},
		{ID: uuid.New().String(), Title: "The Great Gatsby Vol 88", Author: "F. Scott Fitzgerald", Price: 9.97, ISBN: "978000000088"},
		{ID: uuid.New().String(), Title: "Pride and Prejudice Vol 89", Author: "Jane Austen", Price: 8.90, ISBN: "978000000089"},
		{ID: uuid.New().String(), Title: "Fahrenheit 451 Vol 90", Author: "Ray Bradbury", Price: 9.75, ISBN: "978000000090"},
		{ID: uuid.New().String(), Title: "1984 Vol 91", Author: "George Orwell", Price: 10.33, ISBN: "978000000091"},
		{ID: uuid.New().String(), Title: "The Hobbit Vol 92", Author: "J.R.R. Tolkien", Price: 13.44, ISBN: "978000000092"},
		{ID: uuid.New().String(), Title: "Brave New World Vol 93", Author: "Aldous Huxley", Price: 10.41, ISBN: "978000000093"},
		{ID: uuid.New().String(), Title: "The Alchemist Vol 94", Author: "Paulo Coelho", Price: 10.78, ISBN: "978000000094"},
		{ID: uuid.New().String(), Title: "The Book Thief Vol 95", Author: "Markus Zusak", Price: 12.88, ISBN: "978000000095"},
		{ID: uuid.New().String(), Title: "Where the Crawdads Sing Vol 96", Author: "Delia Owens", Price: 12.12, ISBN: "978000000096"},
		{ID: uuid.New().String(), Title: "Becoming Vol 97", Author: "Michelle Obama", Price: 15.33, ISBN: "978000000097"},
		{ID: uuid.New().String(), Title: "Sapiens Vol 98", Author: "Yuval Noah Harari", Price: 13.21, ISBN: "978000000098"},
		{ID: uuid.New().String(), Title: "Educated Vol 99", Author: "Tara Westover", Price: 12.45, ISBN: "978000000099"},
		{ID: uuid.New().String(), Title: "The Silent Patient Vol 100", Author: "Alex Michaelides", Price: 12.11, ISBN: "978000000100"},
	}

	for _, book := range books {
		book.Prepare()
		if err := DB.Create(&book).Error; err != nil {
			log.Printf("Failed to seed book: %s - %v", book.Title, err)
		}
	}

	log.Println("Book seeding complete.")
}
