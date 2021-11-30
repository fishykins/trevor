package lang

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// A handler for the mongodb language collection. You can have multiple of these kicking around your app no worries!
type Dictionary struct {
	collection *mongo.Collection
}

// Creates a new dictionary handler.
func NewDictionary(collection *mongo.Collection) Dictionary {
	return Dictionary{collection: collection}
}

// Gets a random verb that has a description.
func (d *Dictionary) RandVerb(num int) ([]Verb, error) {
	words, err := d.RandWord(num, "verb")
	if err != nil {
		return nil, err
	}
	verbs := make([]Verb, num)
	for i, word := range words {
		verbs[i] = NewVerb(word)
	}
	return verbs, nil
}

// Gets a random noun that has a description.
func (d *Dictionary) RandNoun(num int) ([]Noun, error) {
	words, err := d.RandWord(num, "noun")
	if err != nil {
		return nil, err
	}
	nouns := make([]Noun, num)
	for i, word := range words {
		nouns[i] = NewNoun(word)
	}
	return nouns, nil
}

// Gets a random adjective that has a description.
func (d *Dictionary) RandAdjective(num int) ([]Adjective, error) {
	words, err := d.RandWord(num, "adjective")
	if err != nil {
		return nil, err
	}
	adjectives := make([]Adjective, num)
	for i, word := range words {
		adjectives[i] = NewAdjective(word)
	}
	return adjectives, nil
}

// Gets a random word of given type. Only words with a description are returned.
func (d *Dictionary) RandWord(num int, wordType string) ([]string, error) {
	pipeline := []bson.M{{"$match": bson.M{"type": wordType}}, {"$match": bson.M{"description": bson.M{"$ne": ""}}}, {"$sample": bson.M{"size": num}}}
	cursor, err := d.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	var words []string

	for cursor.Next(context.Background()) {
		var word Word
		err := cursor.Decode(&word)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		words = append(words, word.Inner)
	}
	return words, nil
}

// Finds the given word. Returns nil if not found, and errors if data is missing.
func (d *Dictionary) LookupWord(word string) (*Word, error) {
	var result Word
	searchResult := d.collection.FindOne(context.Background(), bson.M{"word": strings.ToLower(word)})
	fmt.Println(searchResult)
	err := searchResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	if result.Inner == "" {
		return nil, fmt.Errorf("Word \"%s\" not found", word)
	}
	if result.Type == "" {
		return &result, fmt.Errorf("Word \"%s\" has no type", word)
	}
	if result.Description == "" {
		return &result, fmt.Errorf("Word \"%s\" has no description", word)
	}
	return &result, nil
}
