package helpers

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
