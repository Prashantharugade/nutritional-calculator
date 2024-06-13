package main

import "fmt"

func main() {

	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(500),
		Sugars:              SugarGram(24),
		SaturatedFattyAcids: SaturatedFattyAcids(5),
		Sodium:              SodiumMiligram(700),
		Fruits:              FruitsPercent(50),
		Fibre:               FibreGram(6),
		Protein:             ProteinGram(10),
	}, Food)

	fmt.Printf("Nutritional Score: %d \n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
