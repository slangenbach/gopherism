package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce

}

func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	var scaledRecipe []float64
	factor := float64(portions) / 2.0

	for i := 0; i < len(amounts); i++ {
		scaledRecipe = append(scaledRecipe, amounts[i]*factor)
	}
	return scaledRecipe
}
