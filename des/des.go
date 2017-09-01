package des

type Des struct {
	key uint16
	//cypherKey uint16
}

func New(key uint16) *Des {
	d := new(Des)
	d.key = key
	//cypherKey = encryptKey(key)//Possivel?
	return d
}

func (self *Des) Encrypt(plainText string, key uint16) string {
	cypherKey := encryptKey(key)
	var cypher string

	for _, v := range plainText {
		cypher += encryptBlock(v)
	}

	return cypher
}

func (self *Des) Decrypt(encriptedText string, key uint16) string {
	//pegar ultimos 4 bits
	var deCypher string

	for _, v := range encriptedText {
		deCypher += decryptBlock(v)
	}

	return deCypher
}

func encryptKey(key uint16) {
	p10 := [10]uint16{3, 5, 2, 7, 4, 10, 1, 9, 8, 6}
	p8 := [8]uint16{6, 3, 7, 4, 8, 5, 10, 9}

	//.....
}

func encryptBlock(block uint8) {
	end := encryptLastFourBits(block)
	begin := getFirstFour(block)
	xor := end ^ begin //XOR
	return addFirstFourLastFour(xor, getLastFour(block))
}

func decryptBlock(v uint8) {
	cypheredBegin := encryptLastFourBits(v)
	result := getFirstFour(v)
	firstFour := reverseXor(cypheredBegin, result)
	return addFirstFourLastFour(firstFour, getLastFour(v))
}

func expand4To8(block uint8) uint8 {
	var temp1 uint8
	var encryptedParts uint8

	encryptedParts = 0

	// 4ª posicao  para 1ª (em um elemento de 8 bits, a 4ª posição é a última)
	temp1 = block << 7

	encryptedParts = temp1

	// 1ª posicao para 2ª
	temp1 = block >> 3 //limpa os bits ao lado direito do bit na primeira posição
	temp1 <<= 7 //limpa os bits ao lado esquerdo do bit na primeira posição
	temp1 >>= 1 //bit isolado, põe na posição correta

	encryptedParts |= temp1

	// 2ª posicao para 3ª
	temp1 = block >> 2
	temp1 <<= 7
	temp1 >>= 2

	encryptedParts |= temp1

	// 3 posicao para 4ª
	temp1 = block >> 1
	temp1 <<= 7
	temp1 >>= 3

	encryptedParts |= temp1

	// 2 posicao para 5ª
	temp1 = block >> 2
	temp1 <<= 7
	temp1 >>= 4

	encryptedParts |= temp1

	// 3 posicao para 6ª
	temp1 = block >> 1
	temp1 <<= 7
	temp1 >>= 5

	encryptedParts |= temp1

	// 4 posicao para 7ª
	temp1 = block << 7
	temp1 >>= 6

	encryptedParts |= temp1

	// 1 posicao para 8ª
	temp1 = block >> 3
	temp1 <<= 7
	temp1 >>= 7

	encryptedParts |= temp1

	return encryptedParts
}

func xorKey(block1 uint8, block2 uint8) uint8 {
	return (block1 ^ block2)
}

func encryptLastFourBits(block uint8, key uint16) uint8 {
	getLastFour(block)

	tempBlock := expand4To8(block)
	tempBlock = xorKey(tempBlock, key)

	//expande 4 em 8 {4,1,2,3,2,3,4,1}
	//XOR com a key cifrada
	//divide em dois e faz o s0 e s1
	//une resultado de s0 e s1
	//returna resultado antes do xor com o inicio do byte
	//....
}

func addFirstFourLastFour(begin uint8, end uint8) {
	//posiciona os primeiros quatro bits no começo do byte
	//realiza o OR para unir o inicio do byte com o final do byte
	return (begin << 4) | end
}
func reverseXor(operand uint8, result uint8) {
	//...
}
func getFirstFour(v uint8) {
	beg := v >> 4 //limpa os ultimos 4 bits
	return beg
}
func getLastFour(v uint8) {
	end := v << 4  //limpa inicio até o 4º bit
	end = end >> 4 //poe os ultimos 4 bits na posição correta
	return end
}
