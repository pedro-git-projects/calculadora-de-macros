package out

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

type figureStruct struct {
	phrase   string
	fontName string
	strict   bool
}

const (
	newLines     string = "\n\n\n\n\n\n"
	instructions string = "Insira o seu peso alvo em quilos delimitando eventuais casas decimais com um ponto."
	example      string = "Por exemplo, \"88.50\""
	arrow        string = "-> "
)

var (
	f1 = figureStruct{"Calculadora", "", true}
	f2 = figureStruct{"de", "", true}
	f3 = figureStruct{"Macros", "", true}
)

var title1 = figure.NewFigure(f1.phrase, f1.fontName, f1.strict)
var title2 = figure.NewFigure(f2.phrase, f2.fontName, f2.strict)
var title3 = figure.NewFigure(f3.phrase, f3.fontName, f3.strict)

func Greet() {
	title1.Print()
	title2.Print()
	title3.Print()
	fmt.Print(newLines)
	fmt.Println(instructions)
	fmt.Println(example)
	fmt.Print(arrow)
}

func Results(weightInLbs float32, totalCalories float32, protein float32, carbs float32, fat float32) {
	fmt.Printf("O seu peso em libras é %2.f.\n", weightInLbs)
	fmt.Println("O consumo recomendado de calorias é 15 calorias por libra.")
	fmt.Printf("A quantidade recomendada de calorias para este peso é %0.fcal.\n", totalCalories)
	fmt.Printf("O consumo recomendado de proteína é de 1 grama por libra, portanto seu consumo indicado é de %0.fg.\n", protein)
	fmt.Printf("O consumo recomendado de carboidratos é de 2 gramas por libra, portanto seu consumo indicado é de %0.fg.\n", carbs)
	fmt.Printf("O consumo recomendado de gordura é obtido através da conversão pelas calorias restantes, portanto seu consumo indicado é de %0.fg.\n", fat)
	fmt.Printf(`
	╔══════════╤═════════════╤══════════╤═════════╗
	║ calorias │ carboidrato │ proteína │ gordura ║
	╠══════════╪═════════════╪══════════╪═════════╣
	║ %.0fcal  │ %.0fg        │ %.0fg     │ %.0fg    ║
	╚══════════╧═════════════╧══════════╧═════════╝
	
	`, totalCalories, carbs, protein, fat)

	fmt.Print("Pressione 'Enter' para terminar o programa...")
}
