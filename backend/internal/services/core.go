package services

import "synthori.space/coffeeTime/internal/database"

func GenerateDatabase() error {
	err := database.CreateDatabase()
	if err != nil {
		return err
	}

	return nil
}
