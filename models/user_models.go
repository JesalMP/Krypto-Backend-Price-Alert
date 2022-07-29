package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDat struct {
	Id          primitive.ObjectID
	Name, Email string
	Alerts      []Alert
}

type Alert struct {
	Id                   primitive.ObjectID
	AlertPrice           float64
	PriceAtAlertCreation float64
	AlertState           string
	AlertCreationTime    string
	AlertTriggerTime     string
}
