package person

import (
	"fmt"
	"strings"
	"time"
)

type Person struct {
	Name           string
	Birthday       time.Time
	PhysicalTraits PhysicalTraits
	Residence      string
	WeddingDay     time.Time
	Pets           []string
	Relatives      map[string]string
	Occupations    []string
}

type PhysicalTraits struct {
	Hair    string
	Eyes    string
	Height  string
	Weight  string
	Comfort string
}

var Katy = Person{
	Name:     "Katy Rose Behlok, nee Miner",
	Birthday: time.Date(1996, time.March, 15, 0, 0, 0, 0, time.UTC),
	PhysicalTraits: PhysicalTraits{
		Hair:    "Brown, wavy, frizzy when humid",
		Eyes:    "Very pretty, greenish, grayish blue",
		Height:  "5'1",
		Weight:  "165lbs",
		Comfort: "very uncomfortable with body at the moment (would like to be 140)",
	},
	Residence:  "2851 Chili Avenue, Rochester NY, 14624",
	WeddingDay: time.Date(2023, time.July, 7, 0, 0, 0, 0, time.UTC),
	Pets: []string{
		"Dog: Genny Kolsch Behlok",
		"Dog: Stella Artois Behlok",
	},
	Relatives: map[string]string{
		"Husband":         "Alex Behlok",
		"Father":          "Jon Miner",
		"Mother":          "Angela Miner",
		"Brother":         "David Miner",
		"BrotherWife":     "Kaitlyn Miner (nee Miskovsky)",
		"UncleLarry":      "Larry Molinelli",
		"UncleMoe":        "Anthony Molinelli",
		"UncleDan":        "Dan Miner",
		"UncleJeff":       "Jeff Miner",
		"AuntAnnie":       "Annie Molinelli (Larry's wife)",
		"AuntDarlene":     "Darlene (Moe's wife)",
		"AuntJoanne":      "Joanne (Dan's wife)",
		"FatherInLaw":     "Eli Behlok",
		"MotherInLaw":     "Kelly Behlok",
		"SisterInLaw":     "Katie Behlok",
		"SisterInLaw2":    "Caitlyn Miner",
		"UncleInLawRay":   "Ray Behlok",
		"UncleInLawRalph": "Ralph Behlok",
		"CousinInLawClay": "Clay Behlok",
		"CousinInLawMatt": "Matt Behlok",
		"SisterBfInLaw":   "Andy Romero",
		"MatGrandpa":      "Frank Molinelli (deceased)",
		"MatGrandma":      "Sarah Molinelli (deceased)",
		"PatGrandpa":      "Roger Miner (deceased)",
		"PatGrandmaBio":   "Donna Miner (divorced Roger 1995)",
		"PatStepGrandma":  "Joan Miner (Roger's 2nd wife)",
		"Brooke":          "Brooke Chatterton, one of katy's 3 best friends. Co-head bridesmaid with Bianca",
		"Bianca":          "Bianca , one of katy's 3 best friends. Co-head bridesmaid with Brooke",
		"Hannah":          "Katy's other best friend. Biomedical engineer based in Boston, but living in Israel until recently.",
	},
	Occupations: []string{
		"4th grade teacher at Churchville Chili Elementary School (Churchville, NY), also teaches dance to children ages 3-17",
		"Dance Instructor for children roughly ages 8-14 at Ashford Dance Studio in the Village Gate, Rochester, NY",
	},
}

// buildMapString takes a map and converts it into a string with each key-value pair on a new line.
func BuildMapString(m map[string]string) string {
	pairs := make([]string, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, fmt.Sprintf("%s: %s", k, v))
	}
	return strings.Join(pairs, "; ")
}

// BuildPersonString converts the Person data into a string suitable for passing to an LLM.
func BuildPersonString(p Person) string {
	traits := p.PhysicalTraits
	traitsStr := fmt.Sprintf("Hair: %s, Eyes: %s, Height: %s, Weight: %s, Comfort: %s",
		traits.Hair, traits.Eyes, traits.Height, traits.Weight, traits.Comfort)

	petsStr := strings.Join(p.Pets, "; ")
	relativesStr := BuildMapString(p.Relatives)
	occupationsStr := strings.Join(p.Occupations, "; ")

	info := fmt.Sprintf(`Name: %s
Birthday: %s
PhysicalTraits: %s
Residence: %s
WeddingDay: %s
Pets: %s
Relatives: %s
Occupations: %s`,
		p.Name,
		p.Birthday.Format("2006-01-02"),
		traitsStr,
		p.Residence,
		p.WeddingDay.Format("2006-01-02"),
		petsStr,
		relativesStr,
		occupationsStr,
	)
	return info
}
