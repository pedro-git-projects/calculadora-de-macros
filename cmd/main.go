package main

import (
	"bufio"
	"fmt"
	"marco-calculator/pkg/helpers"
	"os"
	"strconv"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// global variables
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
	fmt.Println("Insira o seu peso em quilos marcando eventuais casas decimais com ponto.")
	fmt.Printf("Por exemplo, \"88.50\"\n")
	fmt.Printf("-> ")

	input, err := reader.ReadString('\n')

	number, err := strconv.ParseFloat(strings.TrimSuffix(input, "\n"), 64)

	if err != nil {
		fmt.Println("Um erro ocorreu: ", err.Error())
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

	fmt.Printf("As calorias somadas até agora são %.0f\n", caloriesBeforeFat)

	remainingCalories := calories - caloriesBeforeFat
	fmt.Printf("As calorias restantes são %.0f\n", remainingCalories)

	fat := helpers.CalculateFat(remainingCalories)
	fmt.Printf("O consumo recomendado de gordura é de %.0fg\n", fat)

	fatCalories := helpers.CaloriesFromFat(fat)
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
