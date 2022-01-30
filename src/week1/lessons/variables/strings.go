package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// пустая строка по-умолчанию
	var str string

	// со спец символами
	var hello string = "Привет\n\t"

	// без спец символов
	var world string = `Мир\n\t`

	// UTF-8 из коробки
	var stringHelloWorld = "Привет, мир!"

	// одинарные кавычки для байт
	var rawBinary byte = '\x27'

	// rune (uint32) для UTF-8 символов
	var someChinese rune = '课'

	helloWorld := "Привет мир"
	// конкатенация строк
	andGoodMorning := helloWorld + " и доброе утро!"

	// строки неизменяемы
	//helloWorld[0] = 72

	// получение длины строки
	bytelen := len(helloWorld)                    // 19 байт
	symbols := utf8.RuneCountInString(helloWorld) // 10 рун

	// получение подстроки, в байтах, не в символах!
	hello = helloWorld[:12] // Привет, 0-11 байты
	H := helloWorld[0]      // byte, 208, не "П"

	// конвертация в слайс байт и обратно
	byteString := []byte(helloWorld)
	helloWorldNew := string(byteString)

	fmt.Println(str, hello, world, rawBinary, someChinese, stringHelloWorld, helloWorld, andGoodMorning)
	fmt.Println(bytelen, symbols, hello, H)
	fmt.Println(byteString, helloWorldNew)
}
