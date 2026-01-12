package personaldata

import "fmt"

// Personal holds basic user information required for activity
// and calorie calculations.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print outputs personal information to stdout in a human-readable format.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
