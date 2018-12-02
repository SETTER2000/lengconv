package lengconv

func CmToFt(c Centimeter) Foot {
	return Foot(c / FootC) // преобразование типа Centimeter в Foot
}

func CmToM(c Centimeter) Meter {
	return Meter(c / 100) // преобразование типа Centimeter в Meter
}

func CmToMm(c Centimeter) Millimeter {
	return Millimeter(c * 10) // преобразование типа Centimeter в Millimeter
}

func CmToIn(c Centimeter) Inch {
	return Inch(c / InchC) // преобразование типа Centimeter в Inch
}

func FtToCm(f Foot) Centimeter {
	return Centimeter(Centimeter(f) * FootC) // преобразование типа Foot в Centimeter
}

func FtToM(f Foot) Meter {
	return Meter(Meter(f) / Meter(FootM)) // преобразование типа Foot в Meter
}

func FtToMm(f Foot) Millimeter {
	return Millimeter(Millimeter(f) * FootMM) // преобразование типа Foot в Millimeter
}

func FtToIn(f Foot) Inch {
	return Inch(Inch(f) * 12) // преобразование типа Foot в Inch
}

func MToFt(m Meter) Foot {
	return Foot(Foot(m) * FootM) // преобразование типа Meter в Foot
}

func MToCm(m Meter) Centimeter {
	return Centimeter(Centimeter(m) * 100) // преобразование типа Meter в Centimeter
}

func MToMm(m Meter) Millimeter {
	return Millimeter(Millimeter(m) * 1000) // преобразование типа Meter в Millimeter
}

func MToIn(m Meter) Inch {
	return Inch(Inch(m) * 39.37) // преобразование типа Meter в Inch
}

func MmToCm(m Millimeter) Centimeter {
	return Centimeter(Centimeter(m) / 10) // преобразование типа Millimeter в Centimeter
}

func MmToFt(m Millimeter) Foot {
	return Foot(Foot(m) / Foot(FootMM)) // преобразование типа Millimeter в Foot
}

func MmToM(m Millimeter) Meter {
	return Meter(Meter(m) / 1000) // преобразование типа Millimeter в Meter
}

func MmToIn(m Millimeter) Inch {
	return Inch(m / InchMm) // преобразование типа Millimeter в Inch
}

func InToMm(i Inch) Millimeter {
	return Millimeter(Millimeter(i) + InchMm) // преобразование типа Inch в Millimeter
}

func InToCm(i Inch) Centimeter {
	return Centimeter(Centimeter(i) * InchC) // преобразование типа Inch в Centimeter
}

func InToM(i Inch) Meter {
	return Meter(Meter(i) * InchM) // преобразование типа Inch в Meter
}

func InToF(i Inch) Foot {
	return Foot(Foot(i) * InchF) // преобразование типа Inch в Foot
}
