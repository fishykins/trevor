package babel

import (
	"math/rand"
	"time"
)

var yes = []string{
	"Yes",
	"Affirmative",
	"It is done my lord",
	"I think you will find the outcome of this escapade to your liking",
	"Please accept this result, I am sure you will find it to be satisfactory",
	"yes daddy",
	"Ohhhhh yeaaaaaa baby",
	"Done",
	"Bish bash bosh",
	"Hippedy hop, my cpu is hot",
}

var no = []string{
	"No",
	"Negative",
	"It is not done my lord",
	"I dont think this is what you wanted",
	"Oh no, I am afraid this is not what you wanted",
	"No daddy",
	"Roses are red, violets are blue, I've burned a hole in my cpu",
}

func Affirmative() string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(yes)
	return yes[n]
}

func Negative() string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(no)
	return no[n]
}
