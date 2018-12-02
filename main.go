// lengconv конвертирует единицы длины в сантиметры, футы, метры, миллиметры
package lengconv

import "fmt"

type Millimeter float64
type Centimeter float64
type Meter float64
type Foot float64
type Inch float64

const (
	FootC  Centimeter = 30.48
	FootM  Foot       = 3.281
	FootMM Millimeter = 304.8
	InchC  Centimeter = 2.54
	InchMm Millimeter = 25.4
	InchM  Meter      = 0.0254
	InchF  Foot       = 0.0833333
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

func (m Millimeter) String() string {
	return fmt.Sprintf("%g mm", m)
}

func (i Inch) String() string {
	return fmt.Sprintf("%g in", i)
}
