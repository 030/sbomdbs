package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	ID     uint `gorm:"primaryKey"`
	Number int
	Type   string // aws
	Images []Image
}

type Image struct {
	ID                    uint `gorm:"primaryKey"`
	AccountID             uint // Foreign key
	Date                  time.Time
	Name, Version, Digest string
	Components            []Component
}

type Component struct {
	ID            uint `gorm:"primaryKey"`
	ImageID       uint // Foreign key
	Name, Version string
	Severity      string // critical, high, medium, low, unknown
}

func main() {
	// Replace with your PostgreSQL connection string
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=9998 sslmode=disable TimeZone=UTC"

	// Establish a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&Account{}, &Image{}, &Component{})
	if err != nil {
		log.Fatal(err)
	}

	// Check if a user with the same email has been registered today
	var count int64
	today := time.Now().Truncate(24 * time.Hour) // Truncate to get start of the day
	db.Model(&Image{}).Where("date >= ? AND date < ? AND name = 'alpine' AND version = '3.12.1' AND digest = 'sha256:c0e9560cda118f9ec63ddefb4a173a2b2a0347082d7dff7dc14272e7841a5b5a'", today, today.Add(24*time.Hour)).Count(&count)

	if count > 0 {
		fmt.Println("A record for today already exists. Not inserting.")
		return
	}

	newImage := Account{
		Number: 123456789,
		Type:   "aws",
		Images: []Image{
			{
				Date:    time.Now(),
				Name:    "alpine",
				Version: "3.12.1",
				Digest:  "sha256:c0e9560cda118f9ec63ddefb4a173a2b2a0347082d7dff7dc14272e7841a5b5a",
				Components: []Component{
					{Name: "zlib", Version: "1.2.3", Severity: "critical"},
					{Name: "lib", Version: "1.3.4", Severity: "high"},
				},
			},
		},
	}

	// newImage := Image{
	// 	Date:    time.Now(),
	// 	Name:    "nginx",
	// 	Version: "1.25.1",
	// 	Digest:  "sha256:67f9a4f10d147a6e04629340e6493c9703300ca23a2f7f3aa56fe615d75d31ca",
	// }

	// Insert the new user into the database
	result := db.Create(&newImage)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("User inserted successfully")
}
