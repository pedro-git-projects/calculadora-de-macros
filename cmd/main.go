package main

import (
	"bufio"
	"fmt"
	"marco-calculator/pkg/helpers"
	"os"
	"regexp"
	"strconv"

	"github.com/common-nighthawk/go-figure"
)

// constant declaration
const CaloriesPerGramOfProtein = 4

const CaloriesPerGramOfFat = 9

const CaloriesPerGramOfCarb = 4

func main() {

	reader := bufio.NewReader(os.Stdin)
	title1 := figure.NewFigure("Calculadora", "", true)
	title2 := figure.NewFigure("de", "", true)
	title3 := figure.NewFigure("Macros", "", true)
	title1.Print()
	title2.Print()
	title3.Print()

	fmt.Printf("\n\n\n\n\n\n")
	fmt.Println("Insira o seu peso alvo em quilos marcando eventuais casas decimais com ponto.")
	fmt.Printf("Por exemplo, \"88.50\"\n")
	fmt.Printf("-> ")

	input, err := reader.ReadString('\n')

	// regular expression that matches all non 0-9 inclusive characters
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}

	// parse input after all non numerical characters removal
	number, err := strconv.ParseFloat(reg.ReplaceAllLiteralString(input, ""), 64)

	if err != nil {
		panic(err)
	}

	// converting from kg to lbs
	helpers.KgToLbs(&number, number)
	weight := number

	fmt.Printf("O seu peso em libras é %.2f\n", number)
	fmt.Printf("O consumo recomendado de calorias é 15 calorias por libra.\n")

	// calculating the number of calories
	calories := helpers.NumberOfCalories(number)
	fmt.Printf("A quantidade recomendada de calorias para este peso é %0.fcal\n", calories)

	// amount of protein
	protein := weight
	fmt.Printf("O consumo recomndado de proteína é 1 grama pro libra, portanto seu consumo recomendado é de %.0fg\n", protein)

	// amount of carbs
	carbs := helpers.WeightToCarbs(weight)
	fmt.Printf("O consumo recomndado de carboidratos é 2 gramas pro libra, portanto seu consumo recomendado é de %.0fg\n", carbs)

	proteinCalories := helpers.CaloriesFromProtein(protein)
	carbsCalories := helpers.CaloriesFromCarbs(carbs)
	caloriesBeforeFat := proteinCalories + carbsCalories

	remainingCalories := calories - caloriesBeforeFat

	fat := helpers.CalculateFat(remainingCalories)
	fmt.Printf("O consumo recomendado de gordura é de %.0fg\n", fat)

	fatCalories := helpers.CaloriesFromFat(fat)
	totalCalories, err := helpers.TotalCalories(proteinCalories, carbsCalories, fatCalories, calories)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`
	╔══════════╤═════════════╤══════════╤═════════╗
	║ calorias │ carboidrato │ proteína │ gordura ║
	╠══════════╪═════════════╪══════════╪═════════╣
	║ %.0fcal  │ %.0fg        │ %.0fg     │ %.0fg     ║
	╚══════════╧═════════════╧══════════╧═════════╝
	
	`, totalCalories, carbs, protein, fat)

	fmt.Print("Pressione 'Enter' para terminar o programa...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
