package des

const (
  tableS0 = {{1, 0, 3, 2}, {3, 2, 1, 0}, {0, 2, 1, 3}, {3, 1, 3, 2}}
  tableS1 = {{1, 1, 2, 3}, {2, 0, 1, 3}, {3, 0, 1, 0}, {2, 1, 0, 3}}
)

type Sdes struct {
  k10 uint16
  k1p8 uint16
  k2p8 uint16
}

func sdesGenerateK10(value uint16) uint16 {
  var tempValue uint16
  tempValue = value

  tempValue = tempValue << 6
  tempValue = tempValue >> 6

  var temp1K10 uint16
  var temp2K10 uint16
  temp1K10 = 0
  temp2K10 = 0

  temp1K10 = tempValue >> 7
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 6
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 5
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 7
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 8
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 8
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 3
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 9
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 6
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 10
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue << 15
  temp1K10 = temp1K10 >> 11
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 9
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 12
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 1
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 13
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 2
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 14
  temp2K10 = temp1K10 | temp2K10

  temp1K10 = 0
  temp1K10 = tempValue >> 4
  temp1K10 = temp1K10 << 15
  temp1K10 = temp1K10 >> 15
  temp2K10 = temp1K10 | temp2K10

  return temp2K10
}

func New() *Sdes {
  temp := &Sdes {
    k10: sdesGenerateK10(),
  }

  temp.sdesK1K2()

  return temp
}

func (self *Sdes) Encode(message string) string {
  var messageBytes []byte
  messageBytes = []byte(message)

  var finalMessageBytes []byte
  finalMessageBytes = []byte(message)

  var block uint16

  var blockPart1 uint16
  var blockPart2 uint16

  var OriginalBlockPart2 uint16

  for index := 0; index < len(messageBytes); index++ {
    block = self.ip( uint16(messageBytes[index]) )

    // get part 1 and part 2 from IP result

    blockPart1 = block >> 4
    blockPart1 = blockPart1 << 12
    blockPart1 = blockPart1 >> 8

    blockPart2 = block << 12
    blockPart2 = blockPart2 >> 12

    OriginalBlockPart2 = blockPart2 // will be used in the end

    // Expansion of the blockPart2
    blockPart2 = self.ep(blockPart2)

    // xor
    blockPart2 = self.xorK2p8(blockPart2)

    // divide in 2 parts the blockPart2

    var blockPart21 uint16
    var blockPart22 uint16
    blockPart21 = 0
    blockPart22 = 0

    blockPart21 = blockPart2 >> 4
    blockPart21 = blockPart21 << 12
    blockPart21 = blockPart21 >> 8

    blockPart22 = blockPart2 << 12
    blockPart22 = blockPart22 >> 12

    // 8 bits to 4 using functions S0 and S1

    blockPart2 = self.s0s1(blockPart21, blockPart22)

    var temp uint16
    var temp1 uint16
    temp = 0
    temp1 = 0

    temp = blockPart2 >> 2
    temp = temp << 15
    temp = temp >> 12
    temp1 = temp1 | temp

    temp = blockPart2 << 15
    temp = temp >> 13
    temp1 = temp1 | temp

    temp = blockPart2 >> 1
    temp = temp << 15
    temp = temp >> 14
    temp1 = temp1 | temp

    temp = blockPart2 >> 3
    temp = temp << 15
    temp = temp >> 15
    temp1 = temp1 | temp

    // second xor (blockPart1 and blockPart2)

    block = (((blockPart1 >> 4) ^ blockPart2) << 4) | OriginalBlockPart2

    // repeat code above but now using the self.k2p8 --------------------------------------------------

    // get part 1 and part 2 from IP result

    blockPart1 = block >> 4
    blockPart1 = blockPart1 << 12
    blockPart1 = blockPart1 >> 8

    blockPart2 = block << 12
    blockPart2 = blockPart2 >> 12

    OriginalBlockPart2 = blockPart2 // will be used in the end

    // Expansion of the blockPart2
    blockPart2 = self.ep(blockPart2)

    // xor
    blockPart2 = self.xorK2p8(blockPart2)

    // divide in 2 parts the blockPart2

    var blockPart21 uint16
    var blockPart22 uint16
    blockPart21 = 0
    blockPart22 = 0

    blockPart21 = blockPart2 >> 4
    blockPart21 = blockPart21 << 12
    blockPart21 = blockPart21 >> 8

    blockPart22 = blockPart2 << 12
    blockPart22 = blockPart22 >> 12

    // 8 bits to 4 using functions S0 and S1

    blockPart2 = self.s0s1(blockPart21, blockPart22)

    var temp uint16
    var temp1 uint16
    temp = 0
    temp1 = 0

    temp = blockPart2 >> 2
    temp = temp << 15
    temp = temp >> 12
    temp1 = temp1 | temp

    temp = blockPart2 << 15
    temp = temp >> 13
    temp1 = temp1 | temp

    temp = blockPart2 >> 1
    temp = temp << 15
    temp = temp >> 14
    temp1 = temp1 | temp

    temp = blockPart2 >> 3
    temp = temp << 15
    temp = temp >> 15
    temp1 = temp1 | temp

    // second xor (blockPart1 and blockPart2)

    finalMessageBytes[index] = self.ip1( (((blockPart1 >> 4) ^ blockPart2) << 4) | OriginalBlockPart2 )
  }

  return string(messageBytes)
}

func (self *Sdes) Decode(message string) string {
  var messageBytes []byte
  messageBytes = []byte(message)

  var finalMessageBytes []byte
  finalMessageBytes = []byte(message)

  for index := 0; index < len(messageBytes); index++ {

  }

  return string(messageBytes)
}

func (self *Sdes) sdesK1K2() {
  var tempK10 uint16
  tempK10 = 0

  var tempK101 uint16
  var tempK102 uint16
  tempK101 = 0
  tempK102 = 0

  var shitfPart1 uint16
  var shitfPart2 uint16
  var shiftFinal uint16
  shitfPart1 = 0
  shitfPart2 = 0

  var tempKp8 uint16
  tempKp8 = 0

  // SHIFT 1 part 1 ------------------------------------------

  // 0000001111000000

  tempK101 = self.k10 >> 6
  tempK101 = tempK101 << 12
  tempK101 = tempK101 >> 6


  // 0000000000100000

  tempK102 = self.k10 >> 5
  tempK102 = tempK102 << 15
  tempK102 = tempK102 << 10

  // 0000000000100000 | 0000001111000000 = 0000001111100000

  shitfPart1 = tempK101 | tempK102

  // SHIFT 1 part 2 ------------------------------------------

  // 0000000000011110

  tempK101 = 0
  tempK101 = self.k10 >> 1
  tempK101 = tempK101 << 12
  tempK101 = tempK101 >> 11

  // 0000000000000001

  tempK102 = 0
  tempK102 = self.k10 << 15
  tempK102 = tempK102 >> 15

  // 0000000000000001 | 0000000000011110 = 0000000000011111

  shitfPart2 = tempK101 | tempK102

  // the Key after the shift ---------------------------------

  shiftFinal = shitfPart1 | shitfPart2

  // K1 generation -------------------------------------------

  tempK10 = 0
  tempK10 = shiftFinal >> 4
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 8
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 7
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 9
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 3
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 10
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal>> 6
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 11
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 2
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 12
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 5
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 13
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal << 15
  tempK10 = tempK10 >> 14
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 1
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 15
  tempKp8 = tempKp8 | tempK10

  self.k1p8 = tempKp8

  // shift 2 part 1.1 ---------------------------------------

  tempK101 = 0
  tempK102 = 0
  shitfPart1 = 0
  shitfPart2 = 0

  // 0000001111000000

  tempK101 = shiftFinal >> 6
  tempK101 = tempK101 << 12
  tempK101 = tempK101 >> 6

  // 0000000000100000

  tempK102 = shiftFinal >> 5
  tempK102 = tempK102 << 15
  tempK102 = tempK102 >> 10

  // 0000000000100000 | 0000001111000000 = 0000001111100000

  shitfPart1 = tempK101 | tempK102

  // shift 2 part 1.2 ---------------------------------------

  tempK101 = 0
  tempK102 = 0
  shitfPart1 = 0
  shitfPart2 = 0

  // 0000001111000000

  tempK101 = shitfPart1 >> 6
  tempK101 = tempK101 << 12
  tempK101 = tempK101 >> 6

  // 0000000000100000

  tempK102 = shitfPart1 >> 5
  tempK102 = tempK102 << 15
  tempK102 = tempK102 >> 10

  // 0000000000100000 | 0000001111000000 = 0000001111100000

  shitfPart1 = tempK101 | tempK102

  // shift 2 part 2.1 ---------------------------------------

  tempK101 = 0
  tempK102 = 0
  shitfPart1 = 0
  shitfPart2 = 0

  // 0000000000011110

  tempK101 = shiftFinal >> 1
  tempK101 = tempK101 << 12
  tempK101 = tempK101 >> 11

  // 0000000000000001

  tempK102 = shiftFinal << 15
  tempK102 = tempK102 >> 15

  // 0000000000000001 | 0000000000011110 = 0000000000011111

  shitfPart2 = tempK101 | tempK102

  // shift 2 part 2.2 ---------------------------------------

  tempK101 = 0
  tempK102 = 0
  shitfPart1 = 0
  shitfPart2 = 0

  // 0000000000011110

  tempK101 = shitfPart2 >> 1
  tempK101 = tempK101 << 12
  tempK101 = tempK101 >> 11

  // 0000000000000001

  tempK102 = shitfPart2 << 15
  tempK102 = tempK102 >> 15

  // 0000000000000001 | 0000000000011110 = 0000000000011111

  shitfPart2 = tempK101 | tempK102

  // the Key after the shift ---------------------------------

  shiftFinal = shitfPart1 | shitfPart2

  // K2 generation ----------------------------------------
  tempKp8 = 0

  tempK10 = 0
  tempK10 = shiftFinal >> 4
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 8
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 7
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 9
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 3
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 10
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 6
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 11
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 2
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 12
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 5
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 13
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal << 15
  tempK10 = tempK10 >> 14
  tempKp8 = tempKp8 | tempK10

  tempK10 = 0
  tempK10 = shiftFinal >> 1
  tempK10 = tempK10 << 15
  tempK10 = tempK10 >> 15
  tempKp8 = tempKp8 | tempK10

  self.k2p8 = tempKp8
}

// initial permutation
func (self *Sdes) ip(value uint16) uint16 {
  var tempValue uint16
  var tempPermutation uint16
  tempValue = 0
  tempPermutation = 0

  tempPermutation = 0
  tempPermutation = value >> 6
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 8
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 2
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 9
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 5
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 10
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 7
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 11
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 4
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 12
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value << 15
  tempPermutation = tempPermutation >> 13
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 3
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 14
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 1
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 15
  tempValue = tempValue | tempPermutation

  return tempValue
}


// initial permutation -1
func (self *Sdes) ip1(value uint16) uint16 {
  var tempValue uint16
  var tempPermutation uint16
  tempValue = 0
  tempPermutation = 0

  tempPermutation = 0
  tempPermutation = value >> 4
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 8
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 7
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 9
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 5
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 10
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 3
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 11
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 1
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 12
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 6
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 13
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value << 15
  tempPermutation = tempPermutation >> 14
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 2
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 15
  tempValue = tempValue | tempPermutation

  return tempValue
}

// expansion of 4 bits
func (self *Sdes) ep(value uint16) uint16 {
  var tempValue uint16
  var tempPermutation uint16
  tempValue = 0
  tempPermutation = 0

  tempPermutation = 0
  tempPermutation = value << 15
  tempPermutation = tempPermutation >> 8
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 3
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 9
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 2
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 10
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 1
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 11
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 2
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 12
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 1
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 13
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value << 15
  tempPermutation = tempPermutation >> 14
  tempValue = tempValue | tempPermutation

  tempPermutation = 0
  tempPermutation = value >> 3
  tempPermutation = tempPermutation << 15
  tempPermutation = tempPermutation >> 15
  tempValue = tempValue | tempPermutation

  return tempValue
}

func (self *Sdes) xorK1p8(value1 uint16) {
  return (value1 ^ self.k1p8)
}

func (self *Sdes) xorK2p8(value1 uint16) {
  return (value1 ^ self.k2p8)
}

func (self *Sdes) s0s1(blockPart21, blockPart22 uint16) uint16 {
  var i1 uint16
  var j1 uint16
  var i2 uint16
  var j2 uint16

  i1 = blockPart21 >> 6
  i1 = i1 << 14
  i1 = i1 >> 14

  j1 = blockPart22 >> 2
  j1 = j1 << 14
  j1 = j1 >> 14

  i2 = blockPart21 >> 4
  i2 = i2 << 14
  i2 = i2 >> 14

  j2 = blockPart22 << 14
  j2 = j2 >> 14

  var resultPart1 uint16
  resultPart1 = tableS0[j1][i1] << 2

  var resultPart2 uint16
  resultPart2 = tableS2[j2][i2]

  return (resultPart1 | resultPart2)
}
