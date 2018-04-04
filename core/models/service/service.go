package service

import (
	"time"
	"log"
	"github.com/kangbb/todolist/core/models/entities"
)

func NewItem(label string, isFinished bool, t string) entities.ItemInfo {
	item := entities.ItemInfo {
		Label: label,
		IsFinished: isFinished,
		CreateAt: t,
	}
	return item
}

// CreateItem insert a new Item to the db
func CreateItem(label string, isFinished bool, t string) (bool, error) {
	item := NewItem(label, isFinished, t)
	_, err := entities.Engine.InsertOne(item)
	if err != nil {
		return false, err
	}
	return true, nil
}
func GetAllItems() []entities.ItemInfo {
	items := make([]entities.ItemInfo, 0)
	entities.Engine.SQL("select * from item_info").Find(&items)
	return items
}
func GetItems(label string, t string) []entities.ItemInfo {
	items := make([]entities.ItemInfo, 0)
	entities.Engine.Where("items.create_at = ? AND items.label = ?", t, label).Find(&items)
	return items
}

func UpdateItem(label string, isFinished bool, t string) (bool, error) {
	items := GetItems(label, t);
	item := NewItem(label, isFinished, t)
	affected, err := entities.Engine.Id(items[0].ItemID).Cols("age").Update(&item)
	if err != nil {
		log.Println("err: ", err)
		log.Println("affaected: ", affected)
		return false, err
	}
	return true, nil
}

func DeleteItem(label string, isFinished bool, t string) (bool, error) {
  items := GetItems(label, t);
	item := new(entities.ItemInfo)
	affected, err := entities.Engine.Id(items[0].ItemID).Delete(item)
	if err != nil {
		log.Println("err: ", err)
		log.Println("affaected: ", affected)
		return false, err
	}
	return true, nil
}
