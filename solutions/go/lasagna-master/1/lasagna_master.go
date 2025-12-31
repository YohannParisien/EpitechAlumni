package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutesPerLayer int) int {
    if minutesPerLayer == 0 {
        minutesPerLayer = 2
    }
    return minutesPerLayer * len(layers)
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    sauceQuantity := 0.0
    noodlesQuantity := 0
    for i := 0; i < len(layers); i++ {
        if layers[i] == "sauce" {
            sauceQuantity += 0.2
        }
        if layers[i] == "noodles" {
            noodlesQuantity += 50
        }
    }
    return noodlesQuantity, sauceQuantity
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    if myList[len(myList) - 1] == "?" {
        myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
    }
} 
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    newQuantities := []float64{}
    for i := 0; i < len(quantities); i++ {
        newQuantity := (quantities[i] / 2) * float64(portions)
        newQuantities = append(newQuantities, newQuantity)
    }
    return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
