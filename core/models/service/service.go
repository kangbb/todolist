package service

import (
	"log"

	"github.com/kangbb/todolist/core/models/entities"
)

func NewItem(label string, isFinished bool, t string) entities.ItemInfo {
	item := entities.ItemInfo{
		Label:      label,
		IsFinished: isFinished,
		CreateAt:   t,
	}
	return item
}

// CreateItem insert a new Item to the db
func CreateItem(label string, isFinished bool, t string) (bool, error) {
	item := NewItem(label, isFinished, t)
	_, err := entities.MasterEngine.InsertOne(item)
	if err != nil {
		return false, err
	}
	return true, nil
}
func GetAllItems() []entities.ItemInfo {
	items := make([]entities.ItemInfo, 0)
	err := entities.SlaveEngine.Find(&items)
	if err != nil {
		log.Println(err)
	}
	return items
}
func GetItems(label string, t string) []entities.ItemInfo {
	items := make([]entities.ItemInfo, 0)
	err := entities.SlaveEngine.Where("items.create_at = ? AND items.label = ?", t, label).Find(&items)
	if err != nil {
		log.Println(err)
		return nil
	}
	return items
}

func UpdateItem(label string, isFinished bool, t string) (bool, error) {
	items := GetItems(label, t)
	log.Println(items)
	item := NewItem(label, isFinished, t)
	affected, err := entities.MasterEngine.Id(items[0].ItemID).Cols("is_finished").Update(&item)
	if err != nil {
		log.Println("err: ", err)
		log.Println("affaected: ", affected)
		return false, err
	}
	return true, nil
}

func DeleteItem(label string, isFinished bool, t string) (bool, error) {
	items := GetItems(label, t)
	item := new(entities.ItemInfo)
	affected, err := entities.MasterEngine.Id(items[0].ItemID).Delete(item)
	if err != nil {
		log.Println("err: ", err)
		log.Println("affaected: ", affected)
		return false, err
	}
	return true, nil
}
