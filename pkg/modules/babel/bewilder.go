package babel

import (
	"math/rand"
	"time"

	"github.com/fishykins/trevor/pkg/lang"
)

type message func(adj []lang.Adjective, nouns []lang.Noun, verbs []lang.Verb) string

var messages = []message{
	messageA,
	messageB,
}

func Bewilder() string {
	adjectives, _ := Dict.RandAdjective(3)
	nouns, _ := Dict.RandNoun(3)
	verbs, _ := Dict.RandVerb(3)

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(messages)
	messageFunc := messages[n]
	return messageFunc(adjectives, nouns, verbs)

}

func messageA(adj []lang.Adjective, nouns []lang.Noun, verbs []lang.Verb) string {
	return "Oh no, my " + adj[0].Word + " is " + verbs[0].Word + "ing a " + nouns[0].Word + "!"
}

func messageB(adj []lang.Adjective, nouns []lang.Noun, verbs []lang.Verb) string {
	return "I want to " + verbs[0].Word + " on a " + adj[0].Word + nouns[0].Word + "."
}
