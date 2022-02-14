package weight

import "github.com/pedro-git-projects/calculadora-de-macros/pkg/macros"

const (
	carbsCaloriesPerGram   = 4
	proteinCaloriesPerGram = 4
	fatCaloriesPerGram     = 9
)

type UserData struct {
	Kg               float32
	Lbs              float32
	NumberOfCalories float32
	macros.Macros
}

func (u *UserData) KgToLbs() float32 {
	u.Lbs = u.Kg * 2.2046
	return u.Lbs
}

func (u *UserData) TotalCalories() float32 {
	u.NumberOfCalories = u.Lbs * 15
	return u.NumberOfCalories
}

func (u *UserData) SetProtein() float32 {
	u.Macros.Protein = u.Lbs
	return u.Macros.Protein
}

func (u *UserData) SetCarbs() float32 {
	u.Macros.Carbs = u.Lbs * 2
	return u.Macros.Carbs
}

func (u UserData) GetCaloriesProtein() float32 {
	proteinCalories := u.Lbs * proteinCaloriesPerGram
	return proteinCalories
}

func (u UserData) GetCaloriesCarbs() float32 {
	carbsCalories := u.Lbs * carbsCaloriesPerGram
	return carbsCalories
}

func (u UserData) RemainingiCalories(carbsCalories float32, proteinCalories float32) float32 {
	caloriesUpToThisPoint := carbsCalories + proteinCalories
	remainingiCalories := u.NumberOfCalories - caloriesUpToThisPoint
	return remainingiCalories
}

func (u *UserData) SetFat(remainingiCalories float32) float32 {
	u.Macros.Fat = remainingiCalories / fatCaloriesPerGram
	return u.Macros.Fat
}
