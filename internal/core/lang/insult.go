package lang

import "fmt"

func Insult(name string) (string, error) {
	randNouns, _ := RandNoun(1)
	randAdjectives, _ := RandAdjective(1)

	is := "is"
	if name == "you" {
		is = "are"
	}

	adj := randAdjectives[0]
	noun := randNouns[0]
	return fmt.Sprintf("%s %s a %s %s", name, is, adj.Word, noun.Word), nil
}
