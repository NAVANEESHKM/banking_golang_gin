package models

import("time")

type Customer struct{

	BankID int  `json:"bankid" bson:"bankid"`
	Customer_Name string `json:"customer_name" bson:"customer_name"`
	DOB time.Time `json:"dob" bson:"dob"`
	Password []byte `json:"pasword" bson:"password"`
	Customer_ID int  `json:"customer_id" bson:"customer_id"`
	Balance int  `json:"balance" bson:"balance"`
}

