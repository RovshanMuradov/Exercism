package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	rem := OvenTime - actualMinutesInOven
	return rem
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	pt := numberOfLayers * 2
	return pt
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	el := PreparationTime(numberOfLayers) + actualMinutesInOven
	return el
}
