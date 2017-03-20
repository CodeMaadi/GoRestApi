package main 

type User struct {
    Id string `json:"Id,omitempty"`
	UserName string `json:"UserName,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	MiddleName string `json:"MiddleName,omitempty"`
	LastName string `json:"LastName,omitempty"`
	UserImage string `json:"UserImage,omitempty"`
	UserLink string `json:"UserLink,omitempty"`
	Address string `json:"Address,omitempty"`
	City string `json:"City,omitempty"`
	State string `json:"State,omitempty"`
	ZipCode string `json:"ZipCode,omitempty"`
	Country string `json:"Country,omitempty"`
	AadharID string `json:"AadharId,omitempty"`
	PanCardNumber string `json:"PanCardNumber,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	Email string `json:"Email,omitempty"`
	Role string `json:"Role,omitempty"`
}
