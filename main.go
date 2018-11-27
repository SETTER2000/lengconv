// Преобразования единиц длины
// Сантиметры, Футы, Метры
package lengconv

import "fmt"

type Centimeter float64
type Meter float64
type Foot float64

const (
	FootC Centimeter = 30.48
)

func (c Centimeter) String() string {
	return fmt.Sprintf("%g cm", c)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g ft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}
