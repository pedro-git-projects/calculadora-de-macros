package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// global variables
const CaloriesPerGramOfProtein = 4

const CaloriesPerGramOfFat = 9

const CaloriesPerGramOfCarb = 4

// KgToLbs takes a pointer to a float and a float and alters the value of a variable from kg to lbs
func KgToLbs(kgPtr *float64, kg float64) float64 {
	*kgPtr = kg * 2.2046
	return *kgPtr
}

// NumberOfCalories calculates the amount of calories from the bodywight in lbs
func NumberOfCalories(weight float64) float64 {
	calories := weight * 15
	return calories
}

// WeightToCarbs calculates the amount of carbs from the weight in lbs
func WeightToCarbs(weight float64) float64 {
	carbs := weight * 2
	return carbs
}

// CaloriesFromProtein calculates how many calories have been spent with protein
func CaloriesFromProtein(protein float64) float64 {
	proteinCalories := protein * 4
	return proteinCalories
}

// CaloriesFromCarbs calculates how many calories have been spent with carbs
func CaloriesFromCarbs(carbs float64) float64 {
	carbsCalories := carbs * 4
	return carbsCalories
}

// CalculateFat calculates the recommended fat intake in grams
func CalculateFat(remainingCalories float64) float64 {
	fat := remainingCalories / 9
	return fat
}

// CaloriesFromFat calculates the amount of calories  from fat intake in gram
func CaloriesFromFat(fat float64) float64 {
	fatCalories := fat * 9
	return fatCalories
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	title1 := figure.NewFigure("Calculadora", "", true)
	title2 := figure.NewFigure("de", "", true)
	title3 := figure.NewFigure("Macros", "", true)
	title1.Print()
	title2.Print()
	title3.Print()

	fmt.Printf("\n\n\n\n\n\n")
	fmt.Println("Insira o seu peso em quilos marcando eventuais casas decimais com ponto.")
	fmt.Printf("Por exemplo, \"88.50\"\n")
	fmt.Printf("-> ")

	input, err := reader.ReadString('\n')

	number, err := strconv.ParseFloat(strings.TrimSuffix(input, "\n"), 64)

	if err != nil {
		fmt.Println("Um erro ocorreu: ", err.Error())
	}

	// converting from kg to lbs
	KgToLbs(&number, number)
	weight := number

	fmt.Printf("O seu peso em libras é %.2f\n", number)
	fmt.Printf("O consumo recomendado de calorias é 15 calorias por libra.\n")

	// calculating the number of calories
	calories := NumberOfCalories(number)
	fmt.Printf("A quantidade recomendada de calorias para este peso é %0.fcal\n", calories)

	// amount of protein
	protein := weight
	fmt.Printf("O consumo recomndado de proteína é 1 grama pro libra, portanto seu consumo recomendado é de %.0fg\n", protein)

	// amount of carbs
	carbs := WeightToCarbs(weight)
	fmt.Printf("O consumo recomndado de carboidratos é 2 gramas pro libra, portanto seu consumo recomendado é de %.0fg\n", carbs)

	proteinCalories := CaloriesFromProtein(protein)
	carbsCalories := CaloriesFromCarbs(carbs)
	caloriesBeforeFat := proteinCalories + carbsCalories

	fmt.Printf("As calorias somadas até agora são %.0f\n", caloriesBeforeFat)

	remainingCalories := calories - caloriesBeforeFat
	fmt.Printf("As calorias restantes são %.0f\n", remainingCalories)

	fat := CalculateFat(remainingCalories)
	fmt.Printf("O consumo recomendado de gordura é de %.0fg\n", fat)

	fatCalories := CaloriesFromFat(fat)
	totalCalories := proteinCalories + carbsCalories + fatCalories
	fmt.Printf("A quantidade prevista de calorias foi %0.f, a quantidade obtida de calorias foi %0.f\n", calories, totalCalories)

	fmt.Printf(`
	+--------------+----------+---------+
	| carboidratos | proteína | gordura |
	+--------------+----------+---------+
	|      %.0fg    |    %.0fg  |    %.0fg  |
	+--------------+----------+---------+
	
	`, carbs, protein, fat)
}
