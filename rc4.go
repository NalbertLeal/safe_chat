package main

import (
	"fmt"
	"math"
)
func swap(val1, val2 uint8){
	temp := val1
	val1 = val2
	val2 = temp
}
func main() {
	var K[keylen]uint8
	keylen := 256
	encrypt(K, PlainText, keylen)
}
func encrypt(K []uint8, plainText string, keylen int) {
	
	PT:=[]byte(plainText)
	K:=[]byte(key)
	var S [256]uint8
	var Taux [256]uint8

	//Inicializando S e T (state e temp)
	for i:=0; i<=255; i++ {
		S[i] = i
		Taux[i] = K[math.Mod(i, keylen)]
	}
	
	//Permuntação inicial em S
	for i:=0, j:=0; i<=255; i++ {
		j = math.Mod( (j + S[i] + Taux[i]), 256)
		swap(S[i], S[j])
	}

	for i:=0, j:=0; len(PT); {
		i = math.Mod((i + 1), 256)
		j = math.Mod((j + S[i]), 256)

		swap(S[i], S[j])

		t := math.Mod((S[i]+S[j]), 256)
		k = S[t]

		PT[i] = k^PT[i]
	}
}