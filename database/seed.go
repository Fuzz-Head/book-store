package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Fuzz-Head/domain/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func SeedCompleteBookData() { // []models.Book {

	var bookCount int64
	if err := DB.Model(&models.Book{}).Count(&bookCount).Error; err != nil {
		log.Printf("Failed to check book count: %v", err)
		return
	}

	var publisherCount int64
	if err := DB.Model(&models.Publisher{}).Count(&publisherCount).Error; err != nil {
		log.Printf("Failed to check publisher count: %v", err)
		return
	}

	var categoriesCount int64
	if err := DB.Model(&models.Category{}).Count(&categoriesCount).Error; err != nil {
		log.Printf("Failed to check category count: %v", err)
		return
	}

	var seriesCount int64
	if err := DB.Model(&models.Series{}).Count(&seriesCount).Error; err != nil {
		log.Printf("Failed to check series count: %v", err)
		return
	}

	var authorCount int64
	if err := DB.Model(&models.Author{}).Count(&authorCount).Error; err != nil {
		log.Printf("Failed to check author count: %v", err)
		return
	}

	var authorBookCount int64
	if err := DB.Model(&models.BookAuthor{}).Count(&authorBookCount).Error; err != nil {
		log.Printf("Failed to check book author count: %v", err)
		return
	}

	var awardBook int64
	if err := DB.Model(&models.BookAward{}).Count(&awardBook).Error; err != nil {
		log.Printf("Failed to check book award count: %v", err)
		return
	}

	var editionCount int64
	if err := DB.Model(&models.Edition{}).Count(&editionCount).Error; err != nil {
		log.Printf("Failed to check edition count: %v", err)
		return
	}

	var reviewCount int64
	if err := DB.Model(&models.Review{}).Count(&reviewCount).Error; err != nil {
		log.Printf("Failed to check review count: %v", err)
		return
	}

	var awardCount int64
	if err := DB.Model(&models.Award{}).Count(&awardCount).Error; err != nil {
		log.Printf("Failed to check award count: %v", err)
		return
	}

	if bookCount+publisherCount+categoriesCount+seriesCount+authorCount+authorBookCount+awardBook+editionCount+reviewCount+awardBook > 115 {
		log.Println("Book and other data already seeded — skipping seeding.")
		return
	}

	log.Println("Seeding initial books...")

	// First, create publishers
	publishers := []models.Publisher{
		{
			ID:       uuid.New().String(),
			Name:     "Vintage Books",
			Location: "New York, NY, USA",
			Website:  "https://www.penguinrandomhouse.com/publishers/vintage-books/",
			Founded:  intPtr(1954),
		},
		{
			ID:       uuid.New().String(),
			Name:     "Harper",
			Location: "New York, NY, USA",
			Website:  "https://www.harpercollins.com/",
			Founded:  intPtr(1817),
		},
		{
			ID:       uuid.New().String(),
			Name:     "Random House",
			Location: "New York, NY, USA",
			Website:  "https://www.penguinrandomhouse.com/",
			Founded:  intPtr(1927),
		},
		{
			ID:       uuid.New().String(),
			Name:     "Canongate Books",
			Location: "Edinburgh, Scotland",
			Website:  "https://canongate.co.uk/",
			Founded:  intPtr(1973),
		},
		{
			ID:       uuid.New().String(),
			Name:     "Crown Publishing Group",
			Location: "New York, NY, USA",
			Website:  "https://crownpublishing.com/",
			Founded:  intPtr(1933),
		},
	}

	// Create authors
	authors := []models.Author{
		{
			ID:          uuid.New().String(),
			FirstName:   "Cormac",
			LastName:    "McCarthy",
			Biography:   "Cormac McCarthy is an American novelist, playwright, and screenwriter. He has written twelve novels, two plays, two screenplays, and three short stories, spanning the Western and post-apocalyptic genres.",
			Nationality: "American",
			BirthDate:   timePtr(time.Date(1933, 7, 20, 0, 0, 0, 0, time.UTC)),
			Website:     "https://www.cormacmccarthy.com/",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "George",
			LastName:    "Orwell",
			Biography:   "Eric Arthur Blair, known by his pen name George Orwell, was an English novelist, essayist, journalist and critic. His work is characterised by lucid prose, biting social criticism, opposition to totalitarianism, and outspoken support of democratic socialism.",
			BirthDate:   timePtr(time.Date(1903, 6, 25, 0, 0, 0, 0, time.UTC)),
			DeathDate:   timePtr(time.Date(1950, 1, 21, 0, 0, 0, 0, time.UTC)),
			Nationality: "British",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "J.R.R.",
			LastName:    "Tolkien",
			Biography:   "John Ronald Reuel Tolkien was an English writer, poet, philologist, and academic, best known as the author of the high fantasy works The Hobbit and The Lord of the Rings.",
			BirthDate:   timePtr(time.Date(1892, 1, 3, 0, 0, 0, 0, time.UTC)),
			DeathDate:   timePtr(time.Date(1973, 9, 2, 0, 0, 0, 0, time.UTC)),
			Nationality: "British",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "Jane",
			LastName:    "Austen",
			Biography:   "Jane Austen was an English novelist known primarily for her six major novels, which interpret, critique and comment upon the British landed gentry at the end of the 18th century.",
			BirthDate:   timePtr(time.Date(1775, 12, 16, 0, 0, 0, 0, time.UTC)),
			DeathDate:   timePtr(time.Date(1817, 7, 18, 0, 0, 0, 0, time.UTC)),
			Nationality: "British",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "F. Scott",
			LastName:    "Fitzgerald",
			Biography:   "Francis Scott Key Fitzgerald was an American novelist, essayist, and short story writer. He is best known for his novels depicting the flamboyance and excess of the Jazz Age.",
			BirthDate:   timePtr(time.Date(1896, 9, 24, 0, 0, 0, 0, time.UTC)),
			DeathDate:   timePtr(time.Date(1940, 12, 21, 0, 0, 0, 0, time.UTC)),
			Nationality: "American",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "Ray",
			LastName:    "Bradbury",
			Biography:   "Ray Douglas Bradbury was an American author and screenwriter. One of the most celebrated 20th-century American writers, he worked in a variety of modes, including fantasy, science fiction, horror, mystery, and realistic fiction.",
			BirthDate:   timePtr(time.Date(1920, 8, 22, 0, 0, 0, 0, time.UTC)),
			DeathDate:   timePtr(time.Date(2012, 6, 5, 0, 0, 0, 0, time.UTC)),
			Nationality: "American",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "Yuval Noah",
			LastName:    "Harari",
			Biography:   "Yuval Noah Harari is an Israeli public intellectual, historian and professor in the Department of History at the Hebrew University of Jerusalem.",
			BirthDate:   timePtr(time.Date(1976, 2, 24, 0, 0, 0, 0, time.UTC)),
			Nationality: "Israeli",
			Website:     "https://www.ynharari.com/",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "James",
			LastName:    "Clear",
			Biography:   "James Clear is an American author, entrepreneur, and speaker. He is known for his book Atomic Habits, which has sold over 5 million copies worldwide.",
			BirthDate:   timePtr(time.Date(1986, 1, 1, 0, 0, 0, 0, time.UTC)),
			Nationality: "American",
			Website:     "https://jamesclear.com/",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "Michelle",
			LastName:    "Obama",
			Biography:   "Michelle LaVaughn Robinson Obama is an American attorney and author who served as the first lady of the United States from 2009 to 2017.",
			BirthDate:   timePtr(time.Date(1964, 1, 17, 0, 0, 0, 0, time.UTC)),
			Nationality: "American",
		},
		{
			ID:          uuid.New().String(),
			FirstName:   "Matt",
			LastName:    "Haig",
			Biography:   "Matt Haig is an English author and journalist. He has written both fiction and non-fiction books for children and adults, and is a advocate for mental health awareness.",
			BirthDate:   timePtr(time.Date(1975, 7, 3, 0, 0, 0, 0, time.UTC)),
			Nationality: "British",
			Website:     "https://www.matthaig.com/",
		},
	}

	// Create categories
	categories := []models.Category{
		{
			ID:          uuid.New().String(),
			Name:        "Literary Fiction",
			Description: "Fiction that is regarded as having literary merit",
		},
		{
			ID:          uuid.New().String(),
			Name:        "Science Fiction",
			Description: "Fiction dealing with futuristic concepts",
		},
		{
			ID:          uuid.New().String(),
			Name:        "Fantasy",
			Description: "Fiction involving magical or supernatural elements",
		},
		{
			ID:          uuid.New().String(),
			Name:        "Self-Help",
			Description: "Books designed to help readers improve their lives",
		},
		{
			ID:          uuid.New().String(),
			Name:        "Biography/Memoir",
			Description: "Non-fiction accounts of someone's life",
		},
		{
			ID:          uuid.New().String(),
			Name:        "History",
			Description: "Non-fiction books about historical events and periods",
		},
	}

	// Create series
	series := []models.Series{
		{
			ID:          uuid.New().String(),
			Name:        "Middle-earth",
			Description: "J.R.R. Tolkien's fantasy universe",
			TotalBooks:  intPtr(4),
		},
	}

	// Create awards
	awards := []models.Award{
		{
			ID:          uuid.New().String(),
			Name:        "Pulitzer Prize for Fiction",
			Category:    "Literature",
			Year:        2007,
			Winner:      true,
			Description: "Awarded for distinguished fiction by an American author",
		},
		{
			ID:          uuid.New().String(),
			Name:        "Hugo Award for Best Novel",
			Category:    "Science Fiction",
			Year:        1954,
			Winner:      true,
			Description: "Awarded annually for the best science fiction or fantasy novel",
		},
	}

	for _, publisher := range publishers {
		DB.FirstOrCreate(&publisher, models.Publisher{ID: publisher.ID})
	}

	for _, author := range authors {
		DB.FirstOrCreate(&author, models.Author{ID: author.ID})
	}

	for _, category := range categories {
		DB.FirstOrCreate(&category, models.Category{ID: category.ID})
	}

	for _, s := range series {
		DB.FirstOrCreate(&s, models.Series{ID: s.ID})
	}

	for _, award := range awards {
		DB.FirstOrCreate(&award, models.Award{ID: award.ID})
	}

	now := time.Now()

	// Create complete books
	books := []models.Book{
		{
			// Book 1: The Road by Cormac McCarthy
			ID:                  uuid.New().String(),
			Title:               "The Road",
			Subtitle:            "",
			OriginalTitle:       "The Road",
			Description:         "A father and his son walk alone through burned America. Nothing moves in the ravaged landscape save the ash on the wind. It is cold enough to crack stones, and when the snow falls it is gray. The sky is dark. Their destination is the coast, although they don't know what, if anything, awaits them there.",
			ShortDescription:    "A post-apocalyptic tale of a father and son's journey through a devastated world.",
			PublisherID:         publishers[0].ID,
			Publisher:           publishers[0],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(2006, 9, 26, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(2006, 9, 26, 0, 0, 0, 0, time.UTC)),
			Edition:             "1st",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "Adult",
			ContentWarnings:     pq.StringArray{"Violence", "Disturbing themes"},
			Keywords:            pq.StringArray{"post-apocalyptic", "survival", "father-son", "dystopian"},
			Tags:                pq.StringArray{"literary fiction", "dystopian", "award winner"},
			Genres:              pq.StringArray{"Literary Fiction", "Post-Apocalyptic", "Dystopian"},
			Subjects:            pq.StringArray{"Survival", "Family relationships", "End times"},
			DeweyDecimal:        "813.54",
			LibraryOfCongress:   "PS3563.C337 R63 2006",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "The Road by Cormac McCarthy - Post-Apocalyptic Literary Fiction",
			SEODescription:      "Read The Road, Cormac McCarthy's Pulitzer Prize-winning post-apocalyptic novel about a father and son's journey through a devastated world.",
			MetaKeywords:        pq.StringArray{"Cormac McCarthy", "The Road", "post-apocalyptic", "Pulitzer Prize"},
			Slug:                "the-road-cormac-mccarthy",
			ExternalIDs:         map[string]string{"goodreads_id": "6288", "amazon_asin": "0307387899"},
			ViewCount:           15432,
			PurchaseCount:       3421,
			WishlistCount:       8765,
			Ratings: models.Ratings{
				Average:      4.2,
				Count:        125673,
				Distribution: map[string]int{"5": 42000, "4": 38000, "3": 28000, "2": 12000, "1": 5673},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.2, Count: 125673, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.1, Count: 8934, MaxRate: 5.0},
				},
			},
			ReviewCount: 8934,
			CreatedAt:   now.AddDate(0, -6, 0),
			UpdatedAt:   now.AddDate(0, -1, 0),
		},
		{
			// Book 2: 1984 by George Orwell
			ID:                  uuid.New().String(),
			Title:               "1984",
			Subtitle:            "",
			OriginalTitle:       "Nineteen Eighty-Four",
			Description:         "Winston Smith works for the Ministry of Truth in London, chief city of Airstrip One. Big Brother stares out from every poster, the Thought Police uncover every act of betrayal. When Winston finds love with Julia, he discovers that life does not have to be dull and deadening, and awakens to new possibilities.",
			ShortDescription:    "George Orwell's dystopian masterpiece about totalitarian control and the power of truth.",
			PublisherID:         publishers[1].ID,
			Publisher:           publishers[1],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(1949, 6, 8, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(1949, 6, 8, 0, 0, 0, 0, time.UTC)),
			Edition:             "Modern Classic",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "Young Adult",
			ContentWarnings:     pq.StringArray{"Political themes", "Torture", "Psychological manipulation"},
			Keywords:            pq.StringArray{"dystopian", "totalitarian", "surveillance", "big brother"},
			Tags:                pq.StringArray{"classic", "dystopian", "political fiction"},
			Genres:              pq.StringArray{"Dystopian Fiction", "Political Fiction", "Science Fiction"},
			Subjects:            pq.StringArray{"Totalitarianism", "Surveillance", "Political control"},
			DeweyDecimal:        "823.912",
			LibraryOfCongress:   "PR6029.R8 N5 1949",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "1984 by George Orwell - Classic Dystopian Fiction",
			SEODescription:      "Read George Orwell's 1984, the classic dystopian novel about Big Brother and totalitarian control.",
			MetaKeywords:        pq.StringArray{"George Orwell", "1984", "dystopian", "Big Brother"},
			Slug:                "1984-george-orwell",
			ExternalIDs:         map[string]string{"goodreads_id": "5470", "amazon_asin": "0452284236"},
			ViewCount:           28945,
			PurchaseCount:       6789,
			WishlistCount:       12456,
			Ratings: models.Ratings{
				Average:      4.4,
				Count:        234567,
				Distribution: map[string]int{"5": 98000, "4": 89000, "3": 35000, "2": 8000, "1": 4567},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.4, Count: 234567, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.3, Count: 15432, MaxRate: 5.0},
				},
			},
			ReviewCount: 15432,
			CreatedAt:   now.AddDate(-1, -3, 0),
			UpdatedAt:   now.AddDate(0, -2, 0),
		},
		{
			// Book 3: The Hobbit by J.R.R. Tolkien
			ID:                  uuid.New().String(),
			Title:               "The Hobbit",
			Subtitle:            "or There and Back Again",
			OriginalTitle:       "The Hobbit",
			Description:         "Bilbo Baggins is a hobbit who enjoys a comfortable, unambitious life, rarely traveling any farther than his pantry or cellar. But his contentment is disturbed when the wizard Gandalf and a company of dwarves arrive on his doorstep one day to whisk him away on an adventure.",
			ShortDescription:    "The classic fantasy adventure that started Middle-earth, following Bilbo Baggins on his unexpected journey.",
			PublisherID:         publishers[2].ID,
			Publisher:           publishers[2],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(1937, 9, 21, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(1937, 9, 21, 0, 0, 0, 0, time.UTC)),
			Edition:             "75th Anniversary",
			SeriesID:            &series[0].ID,
			Series:              &series[0],
			SeriesNumber:        float64Ptr(1.0),
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "All Ages",
			ContentWarnings:     pq.StringArray{"Fantasy violence"},
			Keywords:            pq.StringArray{"fantasy", "adventure", "hobbits", "dragons", "middle-earth"},
			Tags:                pq.StringArray{"fantasy", "classic", "adventure", "children's literature"},
			Genres:              pq.StringArray{"Fantasy", "Adventure", "Children's Literature"},
			Subjects:            pq.StringArray{"Hobbits", "Dragons", "Dwarves", "Adventure"},
			DeweyDecimal:        "823.912",
			LibraryOfCongress:   "PZ7.T574 Ho 1938",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "The Hobbit by J.R.R. Tolkien - Classic Fantasy Adventure",
			SEODescription:      "Read The Hobbit, J.R.R. Tolkien's beloved fantasy adventure about Bilbo Baggins and his unexpected journey.",
			MetaKeywords:        pq.StringArray{"J.R.R. Tolkien", "The Hobbit", "Middle-earth", "fantasy"},
			Slug:                "the-hobbit-jrr-tolkien",
			ExternalIDs:         map[string]string{"goodreads_id": "5907", "amazon_asin": "054792822X"},
			ViewCount:           35678,
			PurchaseCount:       8934,
			WishlistCount:       15673,
			Ratings: models.Ratings{
				Average:      4.6,
				Count:        189234,
				Distribution: map[string]int{"5": 89000, "4": 67000, "3": 25000, "2": 6000, "1": 2234},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.6, Count: 189234, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.5, Count: 12345, MaxRate: 5.0},
				},
			},
			ReviewCount: 12345,
			CreatedAt:   now.AddDate(-2, -1, 0),
			UpdatedAt:   now.AddDate(0, -1, 0),
		},
		{
			// Book 4: Pride and Prejudice by Jane Austen
			ID:                  uuid.New().String(),
			Title:               "Pride and Prejudice",
			Subtitle:            "",
			OriginalTitle:       "Pride and Prejudice",
			Description:         "Elizabeth Bennet is the second eldest in a family of five daughters. Although their mother is eager to see them all married to wealthy men, Elizabeth won't give up her independence without a fight. When she meets the proud and seemingly arrogant Mr. Darcy, Elizabeth's first impressions of him are entirely negative.",
			ShortDescription:    "Jane Austen's beloved romance about Elizabeth Bennet and the proud Mr. Darcy.",
			PublisherID:         publishers[0].ID,
			Publisher:           publishers[0],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(1813, 1, 28, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(1813, 1, 28, 0, 0, 0, 0, time.UTC)),
			Edition:             "Penguin Classics",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "General",
			ContentWarnings:     pq.StringArray{},
			Keywords:            pq.StringArray{"romance", "regency", "marriage", "social class"},
			Tags:                pq.StringArray{"classic", "romance", "british literature"},
			Genres:              pq.StringArray{"Romance", "Classic Literature", "Historical Fiction"},
			Subjects:            pq.StringArray{"Marriage", "Social class", "English society"},
			DeweyDecimal:        "823.7",
			LibraryOfCongress:   "PR4034.P7",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "Pride and Prejudice by Jane Austen - Classic Romance Novel",
			SEODescription:      "Read Pride and Prejudice, Jane Austen's timeless romance about Elizabeth Bennet and Mr. Darcy.",
			MetaKeywords:        pq.StringArray{"Jane Austen", "Pride and Prejudice", "romance", "classic"},
			Slug:                "pride-and-prejudice-jane-austen",
			ExternalIDs:         map[string]string{"goodreads_id": "1885", "amazon_asin": "0141439513"},
			ViewCount:           42156,
			PurchaseCount:       9876,
			WishlistCount:       18234,
			Ratings: models.Ratings{
				Average:      4.3,
				Count:        198765,
				Distribution: map[string]int{"5": 78000, "4": 76000, "3": 32000, "2": 9000, "1": 3765},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.3, Count: 198765, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.2, Count: 14567, MaxRate: 5.0},
				},
			},
			ReviewCount: 14567,
			CreatedAt:   now.AddDate(-1, -8, 0),
			UpdatedAt:   now.AddDate(0, -3, 0),
		},
		{
			// Book 5: The Great Gatsby by F. Scott Fitzgerald
			ID:                  uuid.New().String(),
			Title:               "The Great Gatsby",
			Subtitle:            "",
			OriginalTitle:       "The Great Gatsby",
			Description:         "Jay Gatsby is a man with a past. Nobody knows where he comes from, what he does, or how he made his fortune. But everyone goes to his parties. Set in the 1920s, this is a story of impossible love, dreams, and tragedy set against the decadence of the Jazz Age.",
			ShortDescription:    "F. Scott Fitzgerald's masterpiece about the American Dream and the Jazz Age.",
			PublisherID:         publishers[1].ID,
			Publisher:           publishers[1],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(1925, 4, 10, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(1925, 4, 10, 0, 0, 0, 0, time.UTC)),
			Edition:             "Scribner Classic",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "Young Adult",
			ContentWarnings:     pq.StringArray{"Adult themes", "Alcohol use"},
			Keywords:            pq.StringArray{"jazz age", "american dream", "love", "tragedy"},
			Tags:                pq.StringArray{"classic", "american literature", "jazz age"},
			Genres:              pq.StringArray{"Literary Fiction", "Classic Literature", "Historical Fiction"},
			Subjects:            pq.StringArray{"American Dream", "Social class", "1920s America"},
			DeweyDecimal:        "813.52",
			LibraryOfCongress:   "PS3511.I9 G7",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "The Great Gatsby by F. Scott Fitzgerald - American Classic",
			SEODescription:      "Read The Great Gatsby, F. Scott Fitzgerald's iconic novel about Jay Gatsby and the American Dream.",
			MetaKeywords:        pq.StringArray{"F. Scott Fitzgerald", "The Great Gatsby", "American Dream", "Jazz Age"},
			Slug:                "the-great-gatsby-f-scott-fitzgerald",
			ExternalIDs:         map[string]string{"goodreads_id": "4671", "amazon_asin": "0743273567"},
			ViewCount:           38721,
			PurchaseCount:       7654,
			WishlistCount:       16789,
			Ratings: models.Ratings{
				Average:      3.9,
				Count:        156789,
				Distribution: map[string]int{"5": 45000, "4": 52000, "3": 38000, "2": 15000, "1": 6789},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 3.9, Count: 156789, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.0, Count: 11234, MaxRate: 5.0},
				},
			},
			ReviewCount: 11234,
			CreatedAt:   now.AddDate(-1, -5, 0),
			UpdatedAt:   now.AddDate(0, -2, 0),
		},
		{
			// Book 6: Fahrenheit 451 by Ray Bradbury
			ID:                  uuid.New().String(),
			Title:               "Fahrenheit 451",
			Subtitle:            "",
			OriginalTitle:       "Fahrenheit 451",
			Description:         "Guy Montag is a fireman. His job is to destroy the most illegal of commodities, the printed book, along with the houses in which they are hidden. Montag never questions the destruction and ruin his actions produce, returning each day to his bland life and wife, Mildred, who spends all day with her television 'family.'",
			ShortDescription:    "Ray Bradbury's classic dystopian novel about censorship and the power of books.",
			PublisherID:         publishers[2].ID,
			Publisher:           publishers[2],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(1953, 10, 19, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(1953, 10, 19, 0, 0, 0, 0, time.UTC)),
			Edition:             "60th Anniversary",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "Young Adult",
			ContentWarnings:     pq.StringArray{"Book burning", "Dystopian themes"},
			Keywords:            pq.StringArray{"dystopian", "censorship", "books", "fireman"},
			Tags:                pq.StringArray{"science fiction", "dystopian", "classic"},
			Genres:              pq.StringArray{"Science Fiction", "Dystopian Fiction", "Classic Literature"},
			Subjects:            pq.StringArray{"Censorship", "Technology", "Book burning"},
			DeweyDecimal:        "813.54",
			LibraryOfCongress:   "PS3503.R167 F3",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "Fahrenheit 451 by Ray Bradbury - Classic Dystopian Science Fiction",
			SEODescription:      "Read Fahrenheit 451, Ray Bradbury's powerful novel about censorship and the importance of books.",
			MetaKeywords:        pq.StringArray{"Ray Bradbury", "Fahrenheit 451", "dystopian", "censorship"},
			Slug:                "fahrenheit-451-ray-bradbury",
			ExternalIDs:         map[string]string{"goodreads_id": "13079982", "amazon_asin": "1451673319"},
			ViewCount:           29876,
			PurchaseCount:       6543,
			WishlistCount:       13245,
			Ratings: models.Ratings{
				Average:      4.1,
				Count:        143521,
				Distribution: map[string]int{"5": 58000, "4": 54000, "3": 23000, "2": 6000, "1": 2521},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.1, Count: 143521, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.2, Count: 9876, MaxRate: 5.0},
				},
			},
			ReviewCount: 9876,
			CreatedAt:   now.AddDate(-1, -7, 0),
			UpdatedAt:   now.AddDate(0, -1, 0),
		},
		{
			// Book 7: Sapiens by Yuval Noah Harari
			ID:                  uuid.New().String(),
			Title:               "Sapiens",
			Subtitle:            "A Brief History of Humankind",
			OriginalTitle:       "Sapiens: A Brief History of Humankind",
			Description:         "From a renowned historian comes a groundbreaking narrative of humanity's creation and evolution—a #1 international bestseller—that explores the ways in which biology and history have defined us and enhanced our understanding of what it means to be \"human.\"",
			ShortDescription:    "Yuval Noah Harari's groundbreaking exploration of human history and civilization.",
			PublisherID:         publishers[3].ID,
			Publisher:           publishers[3],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(2014, 9, 4, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC)),
			Edition:             "International Edition",
			Language:            "en",
			OriginalLanguage:    "he",
			TranslatedFrom:      "Hebrew",
			AgeRating:           "Adult",
			ContentWarnings:     pq.StringArray{"Complex historical themes"},
			Keywords:            pq.StringArray{"history", "anthropology", "evolution", "civilization"},
			Tags:                pq.StringArray{"non-fiction", "history", "bestseller", "popular science"},
			Genres:              pq.StringArray{"History", "Anthropology", "Popular Science"},
			Subjects:            pq.StringArray{"Human evolution", "Civilization", "Agriculture", "Religion"},
			DeweyDecimal:        "909",
			LibraryOfCongress:   "D21 .H265 2015",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "Sapiens by Yuval Noah Harari - A Brief History of Humankind",
			SEODescription:      "Read Sapiens, Yuval Noah Harari's bestselling exploration of human history from the Stone Age to the present.",
			MetaKeywords:        pq.StringArray{"Yuval Noah Harari", "Sapiens", "human history", "evolution"},
			Slug:                "sapiens-yuval-noah-harari",
			ExternalIDs:         map[string]string{"goodreads_id": "23692271", "amazon_asin": "0062316095"},
			ViewCount:           52341,
			PurchaseCount:       12456,
			WishlistCount:       23789,
			Ratings: models.Ratings{
				Average:      4.5,
				Count:        298765,
				Distribution: map[string]int{"5": 134000, "4": 112000, "3": 38000, "2": 10000, "1": 4765},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.5, Count: 298765, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.4, Count: 18234, MaxRate: 5.0},
				},
			},
			ReviewCount: 18234,
			CreatedAt:   now.AddDate(0, -10, 0),
			UpdatedAt:   now.AddDate(0, -1, 0),
		},
		{
			// Book 8: Atomic Habits by James Clear
			ID:                  uuid.New().String(),
			Title:               "Atomic Habits",
			Subtitle:            "An Easy & Proven Way to Build Good Habits & Break Bad Ones",
			OriginalTitle:       "Atomic Habits",
			Description:         "No matter your goals, Atomic Habits offers a proven framework for improving--every day. James Clear, one of the world's leading experts on habit formation, reveals practical strategies that will teach you exactly how to form good habits, break bad ones, and master the tiny behaviors that lead to remarkable results.",
			ShortDescription:    "James Clear's practical guide to building good habits and breaking bad ones.",
			PublisherID:         publishers[4].ID,
			Publisher:           publishers[4],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(2018, 10, 16, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(2018, 10, 16, 0, 0, 0, 0, time.UTC)),
			Edition:             "1st",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "General",
			ContentWarnings:     pq.StringArray{},
			Keywords:            pq.StringArray{"habits", "self-improvement", "productivity", "psychology"},
			Tags:                pq.StringArray{"self-help", "productivity", "psychology", "bestseller"},
			Genres:              pq.StringArray{"Self-Help", "Psychology", "Personal Development"},
			Subjects:            pq.StringArray{"Habit formation", "Behavioral psychology", "Personal improvement"},
			DeweyDecimal:        "158.1",
			LibraryOfCongress:   "BF335 .C54 2018",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "Atomic Habits by James Clear - Build Good Habits & Break Bad Ones",
			SEODescription:      "Read Atomic Habits, James Clear's bestselling guide to building good habits and breaking bad ones with proven strategies.",
			MetaKeywords:        pq.StringArray{"James Clear", "Atomic Habits", "habit formation", "self-help"},
			Slug:                "atomic-habits-james-clear",
			ExternalIDs:         map[string]string{"goodreads_id": "40121378", "amazon_asin": "0735211299"},
			ViewCount:           67543,
			PurchaseCount:       15678,
			WishlistCount:       28934,
			Ratings: models.Ratings{
				Average:      4.7,
				Count:        187654,
				Distribution: map[string]int{"5": 109000, "4": 58000, "3": 15000, "2": 3000, "1": 2654},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.7, Count: 187654, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.6, Count: 23456, MaxRate: 5.0},
				},
			},
			ReviewCount: 23456,
			CreatedAt:   now.AddDate(0, -8, 0),
			UpdatedAt:   now.AddDate(0, -1, 0),
		},
		{
			// Book 9: Becoming by Michelle Obama
			ID:                  uuid.New().String(),
			Title:               "Becoming",
			Subtitle:            "",
			OriginalTitle:       "Becoming",
			Description:         "In a life filled with meaning and accomplishment, Michelle Obama has emerged as one of the most iconic and compelling women of our era. As First Lady of the United States of America—the first African American to serve in that role—she helped create the most welcoming and inclusive White House in history.",
			ShortDescription:    "Michelle Obama's powerful memoir about her journey from the South Side of Chicago to the White House.",
			PublisherID:         publishers[4].ID,
			Publisher:           publishers[4],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(2018, 11, 13, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(2018, 11, 13, 0, 0, 0, 0, time.UTC)),
			Edition:             "1st",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "General",
			ContentWarnings:     pq.StringArray{},
			Keywords:            pq.StringArray{"memoir", "first lady", "politics", "inspiration"},
			Tags:                pq.StringArray{"memoir", "biography", "politics", "inspiration"},
			Genres:              pq.StringArray{"Biography", "Memoir", "Politics"},
			Subjects:            pq.StringArray{"First Ladies", "African American women", "Politics"},
			DeweyDecimal:        "973.932092",
			LibraryOfCongress:   "E909.O24 A3 2018",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "Becoming by Michelle Obama - Powerful Memoir",
			SEODescription:      "Read Becoming, Michelle Obama's inspiring memoir about her journey from Chicago to the White House.",
			MetaKeywords:        pq.StringArray{"Michelle Obama", "Becoming", "memoir", "First Lady"},
			Slug:                "becoming-michelle-obama",
			ExternalIDs:         map[string]string{"goodreads_id": "38746485", "amazon_asin": "1524763136"},
			ViewCount:           89234,
			PurchaseCount:       18765,
			WishlistCount:       34567,
			Ratings: models.Ratings{
				Average:      4.8,
				Count:        456789,
				Distribution: map[string]int{"5": 298000, "4": 123000, "3": 25000, "2": 7000, "1": 3789},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.8, Count: 456789, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.7, Count: 34567, MaxRate: 5.0},
				},
			},
			ReviewCount: 34567,
			CreatedAt:   now.AddDate(0, -9, 0),
			UpdatedAt:   now.AddDate(0, -1, 0),
		},
		{
			// Book 10: The Midnight Library by Matt Haig
			ID:                  uuid.New().String(),
			Title:               "The Midnight Library",
			Subtitle:            "",
			OriginalTitle:       "The Midnight Library",
			Description:         "Between life and death there is a library, and within that library, the shelves go on forever. Every book provides a chance to try another life you could have lived. To see how things would be if you had made other choices... Would you have done anything different, if you had the chance to undo your regrets?",
			ShortDescription:    "Matt Haig's philosophical novel about life, regret, and infinite possibilities.",
			PublisherID:         publishers[3].ID,
			Publisher:           publishers[3],
			Categories:          []models.Category{},
			PublicationDate:     timePtr(time.Date(2020, 8, 13, 0, 0, 0, 0, time.UTC)),
			OriginalPublishDate: timePtr(time.Date(2020, 8, 13, 0, 0, 0, 0, time.UTC)),
			Edition:             "1st",
			Language:            "en",
			OriginalLanguage:    "en",
			AgeRating:           "Adult",
			ContentWarnings:     pq.StringArray{"Suicide themes", "Depression"},
			Keywords:            pq.StringArray{"philosophy", "life choices", "regret", "parallel lives"},
			Tags:                pq.StringArray{"literary fiction", "philosophy", "contemporary"},
			Genres:              pq.StringArray{"Literary Fiction", "Philosophy", "Contemporary Fiction"},
			Subjects:            pq.StringArray{"Life choices", "Regret", "Parallel universes", "Mental health"},
			DeweyDecimal:        "823.92",
			LibraryOfCongress:   "PR6108.A344 M53 2020",
			Status:              "published",
			Featured:            true,
			Bestseller:          true,
			NewRelease:          false,
			StaffPick:           true,
			SEOTitle:            "The Midnight Library by Matt Haig - Philosophical Fiction",
			SEODescription:      "Read The Midnight Library, Matt Haig's thought-provoking novel about life, regret, and infinite possibilities.",
			MetaKeywords:        pq.StringArray{"Matt Haig", "The Midnight Library", "philosophy", "life choices"},
			Slug:                "the-midnight-library-matt-haig",
			ExternalIDs:         map[string]string{"goodreads_id": "52578297", "amazon_asin": "0525559477"},
			ViewCount:           45672,
			PurchaseCount:       9876,
			WishlistCount:       19234,
			Ratings: models.Ratings{
				Average:      4.2,
				Count:        234567,
				Distribution: map[string]int{"5": 89000, "4": 98000, "3": 35000, "2": 9000, "1": 3567},
				Sources: []models.RatingSource{
					{Source: "goodreads", Rating: 4.2, Count: 234567, MaxRate: 5.0},
					{Source: "amazon", Rating: 4.1, Count: 16789, MaxRate: 5.0},
				},
			},
			ReviewCount: 16789,
			CreatedAt:   now.AddDate(0, -4, 0),
			UpdatedAt:   now.AddDate(0, -1, 0),
		},
	}

	bookToCategoryMap := []int{
		0, // The Road -> Literary Fiction
		1, // 1984 -> Science Fiction
		2, // The Hobbit -> Fantasy
		0, // Pride and Prejudice -> Literary Fiction
		0, // The Great Gatsby -> Literary Fiction
		1, // Fahrenheit 451 -> Science Fiction
		5, // Sapiens -> History
		3, // Atomic Habits -> Self-Help
		4, // Becoming -> Biography/Memoir
		0, // The Midnight Library -> Literary Fiction
	}

	bookToAuthorMap := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	// Now create editions for each book
	for i, book := range books {
		DB.Create(&books[i])

		// Create multiple editions for each book
		editions := createEditionsForBook(book.ID, i)
		for _, edition := range editions {
			DB.Create(&edition)
		}
		books[i].Editions = editions

		// Create authors relationship
		if i < len(bookToAuthorMap) {
			authorIndex := bookToAuthorMap[i]
			bookAuthors := createBookAuthors(book.ID, i, authors[authorIndex])
			for _, bookAuthor := range bookAuthors {
				DB.Create(&bookAuthor)
			}
			books[i].Authors = bookAuthors
		}

		// Add categories
		if i < len(bookToCategoryMap) {
			categoryIndex := bookToCategoryMap[i]
			books[i].Categories = []models.Category{categories[categoryIndex]}
		}

		// For GORM many2many, you might need to explicitly save the association
		if len(books[i].Categories) > 0 {
			DB.Model(&books[i]).Association("Categories").Replace(books[i].Categories)
		}
		// Create some reviews
		reviews := createReviewsForBook(book.ID, i)
		for _, review := range reviews {
			DB.Create(&review)
		}
		books[i].Reviews = reviews

		// Create awards for some books - only for books that have corresponding awards
		if i < len(awards) {
			bookAwards := createBookAwards(book.ID, awards[i])
			for _, bookAward := range bookAwards {
				DB.Create(&bookAward)
			}
			books[i].Awards = bookAwards
		}
	}

}

// Helper functions
func intPtr(i int) *int {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func float64Ptr(f float64) *float64 {
	return &f
}

func createEditionsForBook(bookID string, bookIndex int) []models.Edition {
	//now := time.Now()

	editions := []models.Edition{
		{
			ID:        uuid.New().String(),
			BookID:    bookID,
			Type:      models.BookTypePaperback,
			Language:  "en",
			PageCount: intPtr(320 + bookIndex*50),
			ISBN10:    fmt.Sprintf("054792%04d", bookIndex+1),
			ISBN13:    fmt.Sprintf("978000000%03d", bookIndex+1),
			Dimensions: &models.Dimensions{
				Length: 19.8,
				Width:  12.9,
				Height: 2.1,
				Weight: 280.5,
			},
			Pricing: models.Pricing{
				Currency:  "USD",
				BasePrice: 9.99 + float64(bookIndex)*0.5,
				SalePrice: float64Ptr(8.99 + float64(bookIndex)*0.5),
			},
			Availability: models.Availability{
				InStock:    true,
				StockCount: intPtr(150 + bookIndex*10),
				PreOrder:   false,
			},
			CoverImages: pq.StringArray{
				fmt.Sprintf("https://images.example.com/covers/%s-paperback.jpg", bookID),
			},
		},
		{
			ID:        uuid.New().String(),
			BookID:    bookID,
			Type:      models.BookTypeHardcover,
			Language:  "en",
			PageCount: intPtr(320 + bookIndex*50),
			ISBN13:    fmt.Sprintf("978000001%03d", bookIndex+1),
			Dimensions: &models.Dimensions{
				Length: 23.4,
				Width:  15.6,
				Height: 3.2,
				Weight: 580.2,
			},
			Pricing: models.Pricing{
				Currency:  "USD",
				BasePrice: 24.99 + float64(bookIndex)*1.0,
			},
			Availability: models.Availability{
				InStock:    true,
				StockCount: intPtr(75 + bookIndex*5),
				PreOrder:   false,
			},
			CoverImages: pq.StringArray{
				fmt.Sprintf("https://images.example.com/covers/%s-hardcover.jpg", bookID),
			},
		},
		{
			ID:       uuid.New().String(),
			BookID:   bookID,
			Type:     models.BookTypeEbook,
			Format:   "EPUB",
			Language: "en",
			ISBN13:   fmt.Sprintf("978000002%03d", bookIndex+1),
			Pricing: models.Pricing{
				Currency:  "USD",
				BasePrice: 7.99 + float64(bookIndex)*0.3,
			},
			Availability: models.Availability{
				InStock:  true,
				PreOrder: false,
			},
			CoverImages: pq.StringArray{
				fmt.Sprintf("https://images.example.com/covers/%s-ebook.jpg", bookID),
			},
		},
	}

	return editions
}

func createBookAuthors(bookID string, bookIndex int, author models.Author) []models.BookAuthor {
	// This is simplified - in reality you'd match with actual author IDs
	return []models.BookAuthor{
		{
			BookID:   bookID,
			AuthorID: uuid.New().String(),
			Role:     "author",
			Author:   author,
		},
	}
}

func createReviewsForBook(bookID string, bookIndex int) []models.Review {
	now := time.Now()

	reviews := []models.Review{
		{
			ID:        uuid.New().String(),
			BookID:    bookID,
			UserID:    uuid.New().String(),
			Rating:    4.5,
			Title:     "Excellent read!",
			Content:   "This book exceeded my expectations. The writing is engaging and the story is compelling.",
			Verified:  true,
			Helpful:   23,
			CreatedAt: now.AddDate(0, -2, -5),
			UpdatedAt: now.AddDate(0, -2, -5),
		},
		{
			ID:        uuid.New().String(),
			BookID:    bookID,
			UserID:    uuid.New().String(),
			Rating:    4.0,
			Title:     "Good book",
			Content:   "Well written and thought-provoking. Recommended for fans of the genre.",
			Verified:  true,
			Helpful:   15,
			CreatedAt: now.AddDate(0, -1, -10),
			UpdatedAt: now.AddDate(0, -1, -10),
		},
		{
			ID:        uuid.New().String(),
			BookID:    bookID,
			UserID:    uuid.New().String(),
			Rating:    2.5,
			Title:     "typical read",
			Content:   "This book has nothing new or great typical stuff all condensed in one place.",
			Verified:  true,
			Helpful:   223,
			CreatedAt: now.AddDate(0, -3, -5),
			UpdatedAt: now.AddDate(0, -7, -13),
		},
		{
			ID:        uuid.New().String(),
			BookID:    bookID,
			UserID:    uuid.New().String(),
			Rating:    4.0,
			Title:     "Good read",
			Content:   "Well written specially for a debut novel.",
			Verified:  true,
			Helpful:   15,
			CreatedAt: now.AddDate(0, 0, -10),
			UpdatedAt: now.AddDate(0, 0, -10),
		},
	}

	return reviews
}

func createBookAwards(bookID string, award models.Award) []models.BookAward {
	return []models.BookAward{
		{
			BookID:  bookID,
			AwardID: uuid.New().String(),
			Award:   award,
		},
	}
}
