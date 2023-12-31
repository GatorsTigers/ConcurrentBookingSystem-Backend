package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateScreens(screens *[]models.Screen) error {
	if txn := DbInstance.Db.Create(screens); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func ShowScreens(screens *[]models.Screen) error {
	if txn := DbInstance.Db.Find(screens); txn.Error != nil {
		return txn.Error
	}
	return nil
}
