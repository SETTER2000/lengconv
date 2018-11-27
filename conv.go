package lengconv

func CmToFt(c Centimeter) Foot {
	return Foot(c / FootC) // преобразование типа Centimeter в Foot
}

func CmToM(c Centimeter) Meter {
	return Meter(c / 100) // преобразование типа Centimeter в Meter
}

func FtToCm(f Foot) Centimeter {
	return Centimeter(Centimeter(f) * FootC) // преобразование типа Foot в Centimeter
}

func FtToM(f Foot) Meter {
	return Meter(Meter(f) / Meter(FootM)) // преобразование типа Foot в Meter
}

func MToFt(m Meter) Foot {
	return Foot(Foot(m) * FootM) // преобразование типа Meter в Foot
}

func MToCm(m Meter) Centimeter {
	return Centimeter(Centimeter(m) * 100) // преобразование типа Meter в Centimeter
}
