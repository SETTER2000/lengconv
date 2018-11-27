package lengconv

func CmToFt(c Centimeter) Foot {
	return Foot(c / FootC) // преобразование типа Centimeter в Foot
}

func FtToCm(f Foot) Centimeter {
	return Centimeter(Centimeter(f) * FootC) // преобразование типа Foot в Centimeter
}
