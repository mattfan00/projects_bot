package models

type Project struct {
	Id          string `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	Url         string `bson:"url"`
	Description string `bson:"description"`
	CreatedBy   string `bson:"created_by"`
}
