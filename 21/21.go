package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// Пользователь, который совершает действие
type User struct {
}

// Метод пользователя, который подключает прибор с EU-вилкой в розетку:
func (u *User) InsertEUPlugIntoSocket(socket Socket) {
	fmt.Println("Пользователь вставляет EU-вилку в розетку")
	socket.InsertIntoEUSocket()
}

type Socket interface {
	InsertIntoEUSocket()
}

// Дом, в котором пользователь жил до переезда
type HouseWithEUSockets struct {
}

// В доме была возможность подключить приборы в EU-розетку
func (eu *HouseWithEUSockets) InsertIntoEUSocket() {
	fmt.Println("EU-вилка вставлена в EU-розетку")
}

// Дом, в который пользователь переехал
type HouseWithUSSockets struct {
}

// В доме есть возможность подключить приборы в US-розетку
func (us *HouseWithUSSockets) InsertinIntoUSSocket() {
	fmt.Println("US-вилка вставлена в US-розетку")
}

// Адаптер (переходник) для использования EU-вилки в доме с US-розетками
type USSocketsAdapter struct {
	houseWithUSSockets HouseWithUSSockets
}

// Конструктор для USSocketsAdapter
func NewUSSocketsAdapter(houseWithUSSockets HouseWithUSSockets) *USSocketsAdapter {
	return &USSocketsAdapter{houseWithUSSockets: houseWithUSSockets}
}

// Реализация интерфейса Socket с использованием адаптера
func (usadapter *USSocketsAdapter) InsertIntoEUSocket() {
	fmt.Println("Надеваем US-переходник на EU-вилку")
	usadapter.houseWithUSSockets.InsertinIntoUSSocket()
}

func main() {
	Adam := &User{}

	fmt.Println("Пользователь решил подключить чайник с EU-вилкой, чтобы вскипятить воду")
	HouseInBerlin := &HouseWithEUSockets{}
	Adam.InsertEUPlugIntoSocket(HouseInBerlin)

	fmt.Println("Пользователь переехал из Берлина в Нью-Йорк и теперь дома только US-розетки")
	HouseInNewYork := &HouseWithUSSockets{}
	fmt.Println("Пользователь решил подключить чайник с EU-вилкой, чтобы вскипятить воду")
	EUtoUSplugAdapter := NewUSSocketsAdapter(*HouseInNewYork)
	Adam.InsertEUPlugIntoSocket(EUtoUSplugAdapter)
}
