package main

type UserID int

func main() {
	idx := 1
	var uid UserID = 42

	// даже если базовый тип одинаковый, разные типы не совместимы
	// myID := idx

	myID := UserID(idx)
	println(uid, myID)
}
