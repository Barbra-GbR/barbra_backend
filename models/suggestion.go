package models

import (
	"github.com/Barbra-GbR/barbra-backend/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	suggestionCollectionName = "suggestions"
)

// The Suggestion model is generated by the nlp-layer
type Suggestion struct {
	Provider string        `json:"provider"	bson:"provider" binding:"required"`
	URL      string        `json:"url"     	bson:"url"      binding:"required"`
	Kind     string        `json:"kind"    	bson:"kind"     binding:"required"`
	Title    string        `json:"title"   	bson:"title"    binding:"required"`
	Category string        `json:"category"	bson:"category" binding:"required"`
	Tags     []string      `json:"tags"     bson:"tags"    	binding:"required"`
	Content  string        `json:"content"  bson:"content"  binding:"required"`
	ID       bson.ObjectId `json:"id"       bson:"_id"      binding:"required"`
}

// NewSuggestion creates a new suggestion with the specified data
// The ID will be set automatically
func NewSuggestion(url string, kind string, title string, category string, provider string, tags []string, content string) *Suggestion {
	return &Suggestion{
		ID:       bson.NewObjectId(),
		Content:  content,
		Kind:     kind,
		Tags:     tags,
		Title:    title,
		Category: category,
		URL:      url,
		Provider: provider,
	}
}

// GetSuggestion searches for the suggestion in the database, if no matches are found a new one will be created
// and saved to the Database.
func GetSuggestion(url string, kind string, title string, provider string, category string, tags []string, content string) (*Suggestion, error) {
	collection := db.GetDB().C(suggestionCollectionName)

	suggestion := new(Suggestion)
	err := collection.Find(bson.M{
		"url":      url,
		"kind":     kind,
		"title":    title,
		"category": category,
		"tags":     tags,
		"content":  content,
		"provider": provider,
	}).One(suggestion)
	if err == mgo.ErrNotFound {
		suggestion = NewSuggestion(url, kind, title, category, provider, tags, content)
		err = suggestion.Save()
	}

	return suggestion, err
}

// GetSuggestionsByID returns all to the ids matching suggestions from the Database
// Non existing objects will be ignored
func GetSuggestionsByID(ids []bson.ObjectId) (*[]Suggestion, error) {
	collection := db.GetDB().C(suggestionCollectionName)
	var suggestions []Suggestion
	err := collection.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&suggestions)
	return &suggestions, err
}

// GetSuggestionByID returns a suggestion a the matching id from the Database
func GetSuggestionByID(id bson.ObjectId) (*Suggestion, error) {
	collection := db.GetDB().C(suggestionCollectionName)
	suggestion := new(Suggestion)
	err := collection.FindId(id).One(suggestion)
	return suggestion, err
}

// SuggestionExists checks if the specified suggestion exists in the Database
// If errors occur thrown false will be returned
func SuggestionExists(id bson.ObjectId) bool {
	collection := db.GetDB().C(suggestionCollectionName)
	count, err := collection.FindId(id).Limit(1).Count()
	return count > 0 && err == nil
}

// Save inserts the suggestion into the database. If it already exists it will be updated
func (suggestion *Suggestion) Save() error {
	collection := db.GetDB().C(suggestionCollectionName)
	_, err := collection.UpsertId(suggestion.ID, suggestion)
	return err
}
