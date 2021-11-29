package lang

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Dictionary *mongo.Client

func init() {
	c, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Dictionary = c

}

func Open() error {
	err := Dictionary.Connect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	Dictionary.Disconnect(context.Background())
}

func English() *mongo.Collection {
	babelfish := Dictionary.Database("babelfish")
	return babelfish.Collection("english")
}

func RandVerb(num int) ([]Verb, error) {
	words, err := RandWord(num, "verb")
	if err != nil {
		return nil, err
	}
	verbs := make([]Verb, num)
	for i, word := range words {
		verbs[i] = NewVerb(word)
	}
	return verbs, nil
}

func RandNoun(num int) ([]Noun, error) {
	words, err := RandWord(num, "noun")
	if err != nil {
		return nil, err
	}
	nouns := make([]Noun, num)
	for i, word := range words {
		nouns[i] = NewNoun(word)
	}
	return nouns, nil
}

func RandAdjective(num int) ([]Adjective, error) {
	words, err := RandWord(num, "adjective")
	if err != nil {
		return nil, err
	}
	adjectives := make([]Adjective, num)
	for i, word := range words {
		adjectives[i] = NewAdjective(word)
	}
	return adjectives, nil
}

func RandWord(num int, wordType string) ([]string, error) {
	pipeline := []bson.M{{"$match": bson.M{"type": wordType}}, {"$match": bson.M{"description": bson.M{"$ne": ""}}}, {"$sample": bson.M{"size": num}}}
	cursor, err := English().Aggregate(context.Background(), pipeline)
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
		words = append(words, word.Word)
	}
	return words, nil
}

func LookupWord(word string) (*Word, error) {
	var result Word
	searchResult := English().FindOne(context.Background(), bson.M{"word": strings.ToLower(word)})
	fmt.Println(searchResult)
	err := searchResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	if result.Word == "" || result.WordType == "" || result.Description == "" {
		return nil, fmt.Errorf("Word \"%s\" not found", word)
	}
	return &result, nil
}
