package models

import (
	u "api-go-deploy-heroku/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Contact for struktur model
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_iD"` 
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/

func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message("01", false, "Contact name should be on the payload"), false
	}
	if contact.Phone == "" {
		return u.Message("01", false, "Phone number should be on the payload"), false
	}
	if contact.UserID <= 0 {
		return u.Message("01", false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message("00", true, "success"), true
}

func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}
	GetDB().Create(contact)
	resp := u.Message("00", true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}
func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}
