package main

import (
	"bufio"
	"os"

	"github.com/pedro-git-projects/calculadora-de-macros/pkg/macros"
	"github.com/pedro-git-projects/calculadora-de-macros/pkg/out"
	"github.com/pedro-git-projects/calculadora-de-macros/pkg/regular"
	"github.com/pedro-git-projects/calculadora-de-macros/pkg/weight"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	out.Greet()

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	numericalInput := regular.RemoveNonNumericalInputs(input)

	userWeightStruct := weight.UserData{numericalInput, 0, 0, macros.Macros{0, 0, 0}}
	userWeightLbs := userWeightStruct.KgToLbs()
	totalCalories := userWeightStruct.TotalCalories()
	proteinIntake := userWeightStruct.SetProtein()
	carbIntake := userWeightStruct.SetCarbs()
	remaining := userWeightStruct.RemainingiCalories(userWeightStruct.GetCaloriesCarbs(), userWeightStruct.GetCaloriesProtein())
	fatIntake := userWeightStruct.SetFat(remaining)

	out.Results(userWeightLbs, totalCalories, proteinIntake, carbIntake, fatIntake)

	bufio.NewReader(os.Stdin).ReadBytes('\n')

}
