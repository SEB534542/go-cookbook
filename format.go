package main

import (
	"strconv"
	"strings"
)

/* textToLines takes a string, splits the string into a slice for each new line
and removes all non text characters and empty lines. It returns the slice.*/
func textToLines(s string) []string {
	// Split string into lines
	lines := strings.Split(s, "\n")
	// Remove all non text characters
	for i, _ := range lines {
		for j := uint8(0); j < 32; j++ {
			lines[i] = strings.Replace(lines[i], string(j), "", -1)
		}
	}
	// Remove empty lines
	newLines := []string{}
	for _, line := range lines {
		if line != "" {
			newLines = append(newLines, line)
		}
	}
	return newLines
}

func textToIngrds(s string) []Ingrd {
	lines := textToLines(s)
	ingrs := []Ingrd{}
	// Convert each line to an ingredient
	for _, line := range lines {
		xs := strings.Split(line, " ")
		for i, s := range xs {
			amount, err := strconv.ParseFloat(s, 64)
			if err == nil {
				// Check if a unit is included directly behind the float
				units := map[string]func(float64) (string, float64){
					"gram":        func(f float64) (string, float64) { return gram, f },
					"grams":       func(f float64) (string, float64) { return gram, f },
					"gr":          func(f float64) (string, float64) { return gram, f },
					"g":           func(f float64) (string, float64) { return gram, f },
					"kilogram":    func(f float64) (string, float64) { return gram, f * 1000 },
					"kg":          func(f float64) (string, float64) { return gram, f * 1000 },
					"cup":         func(f float64) (string, float64) { return cup, f },
					"cups":        func(f float64) (string, float64) { return cup, f },
					"ml":          func(f float64) (string, float64) { return ml, f },
					"milliliter":  func(f float64) (string, float64) { return ml, f },
					"liter":       func(f float64) (string, float64) { return ml, f * 1000 },
					"tbsp":        func(f float64) (string, float64) { return tbsp, f },
					"tablespoon":  func(f float64) (string, float64) { return tbsp, f },
					"tablespoons": func(f float64) (string, float64) { return tbsp, f },
					"eetlepel":    func(f float64) (string, float64) { return tbsp, f },
					"eetlepels":   func(f float64) (string, float64) { return tbsp, f },
					"tsp":         func(f float64) (string, float64) { return tsp, f },
					"teaspoon":    func(f float64) (string, float64) { return tsp, f },
					"teaspoons":   func(f float64) (string, float64) { return tsp, f },
					"theelepel":   func(f float64) (string, float64) { return tsp, f },
					"theelepels":  func(f float64) (string, float64) { return tsp, f },
					"stuk":        func(f float64) (string, float64) { return pcs, f },
					"stuks":       func(f float64) (string, float64) { return pcs, f },
					"pieces":      func(f float64) (string, float64) { return pcs, f },
				}
				offset := 0
				unit := pcs
				_, ok := units[xs[i+1]]
				if ok {
					unit, amount = units[xs[i+1]](amount)
					offset += len(xs[i+1])
				}
				ingr := Ingrd{
					Amount: amount,
					Unit:   unit,
					Item:   line[strings.Index(line, s)+len(s)+offset+1:], // assuming item is directly after amount in the text
					Notes:  line[:strings.Index(line, s)],                 // assuming an text before the float is additional notes
				}
				ingrs = append(ingrs, ingr)
			}
		}
	}
	return ingrs
}
