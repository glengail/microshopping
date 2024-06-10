package cart

import "gorm.io/gorm"

func (item *Item) AfterUpdate(tx *gorm.DB) error {
	if item.Count <= 0 {
		//Unscoped不带软删除的的删除条件
		return tx.Unscoped().Delete(&item).Error
	}
	return nil
}
