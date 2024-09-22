package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	part1()
}

type packet struct {
	sender, recipient *computer
	x, y              int
}

type computer struct {
	address     int
	packetMutex *sync.Mutex
	packetQueue []packet
	command     []int
	cursor      int
	base        int
}

var computers [50]*computer = [50]*computer{}
var NAT *computer = &computer{address: 255, packetMutex: &sync.Mutex{}}

const input = "3,62,1001,62,11,10,109,2215,105,1,0,1846,1780,1120,1959,571,2182,1083,631,1485,1306,1157,2153,1654,1592,1454,664,1815,951,2124,794,695,2023,1378,1248,1518,984,1887,2058,823,1019,891,1277,1685,2089,1341,1922,1186,765,1751,1413,602,922,858,1623,1217,1559,1050,1990,734,1720,0,0,0,0,0,0,0,0,0,0,0,0,3,64,1008,64,-1,62,1006,62,88,1006,61,170,1106,0,73,3,65,20101,0,64,1,21001,66,0,2,21101,105,0,0,1105,1,436,1201,1,-1,64,1007,64,0,62,1005,62,73,7,64,67,62,1006,62,73,1002,64,2,133,1,133,68,133,101,0,0,62,1001,133,1,140,8,0,65,63,2,63,62,62,1005,62,73,1002,64,2,161,1,161,68,161,1101,1,0,0,1001,161,1,169,101,0,65,0,1101,0,1,61,1101,0,0,63,7,63,67,62,1006,62,203,1002,63,2,194,1,68,194,194,1006,0,73,1001,63,1,63,1105,1,178,21102,1,210,0,105,1,69,2101,0,1,70,1102,0,1,63,7,63,71,62,1006,62,250,1002,63,2,234,1,72,234,234,4,0,101,1,234,240,4,0,4,70,1001,63,1,63,1105,1,218,1105,1,73,109,4,21101,0,0,-3,21102,0,1,-2,20207,-2,67,-1,1206,-1,293,1202,-2,2,283,101,1,283,283,1,68,283,283,22001,0,-3,-3,21201,-2,1,-2,1106,0,263,22102,1,-3,-3,109,-4,2106,0,0,109,4,21101,0,1,-3,21101,0,0,-2,20207,-2,67,-1,1206,-1,342,1202,-2,2,332,101,1,332,332,1,68,332,332,22002,0,-3,-3,21201,-2,1,-2,1106,0,312,21202,-3,1,-3,109,-4,2106,0,0,109,1,101,1,68,358,21001,0,0,1,101,3,68,366,21001,0,0,2,21102,1,376,0,1106,0,436,22101,0,1,0,109,-1,2106,0,0,1,2,4,8,16,32,64,128,256,512,1024,2048,4096,8192,16384,32768,65536,131072,262144,524288,1048576,2097152,4194304,8388608,16777216,33554432,67108864,134217728,268435456,536870912,1073741824,2147483648,4294967296,8589934592,17179869184,34359738368,68719476736,137438953472,274877906944,549755813888,1099511627776,2199023255552,4398046511104,8796093022208,17592186044416,35184372088832,70368744177664,140737488355328,281474976710656,562949953421312,1125899906842624,109,8,21202,-6,10,-5,22207,-7,-5,-5,1205,-5,521,21102,1,0,-4,21102,1,0,-3,21101,0,51,-2,21201,-2,-1,-2,1201,-2,385,471,20102,1,0,-1,21202,-3,2,-3,22207,-7,-1,-5,1205,-5,496,21201,-3,1,-3,22102,-1,-1,-5,22201,-7,-5,-7,22207,-3,-6,-5,1205,-5,515,22102,-1,-6,-5,22201,-3,-5,-3,22201,-1,-4,-4,1205,-2,461,1106,0,547,21102,-1,1,-4,21202,-6,-1,-6,21207,-7,0,-5,1205,-5,547,22201,-7,-6,-7,21201,-4,1,-4,1106,0,529,21201,-4,0,-7,109,-8,2106,0,0,109,1,101,1,68,564,20102,1,0,0,109,-1,2106,0,0,1102,97571,1,66,1101,1,0,67,1101,598,0,68,1101,556,0,69,1101,0,1,71,1101,0,600,72,1105,1,73,1,-1605712,28,112719,1102,1,48679,66,1102,1,1,67,1102,1,629,68,1102,556,1,69,1101,0,0,71,1102,631,1,72,1106,0,73,1,1782,1102,1,7877,66,1102,1,2,67,1101,658,0,68,1101,302,0,69,1102,1,1,71,1102,1,662,72,1105,1,73,0,0,0,0,26,81559,1101,0,22453,66,1101,1,0,67,1102,1,691,68,1101,0,556,69,1102,1,1,71,1102,1,693,72,1105,1,73,1,125,35,19959,1101,0,21493,66,1101,0,5,67,1102,1,722,68,1101,0,302,69,1102,1,1,71,1101,0,732,72,1106,0,73,0,0,0,0,0,0,0,0,0,0,2,56893,1102,77191,1,66,1101,1,0,67,1102,1,761,68,1101,556,0,69,1101,1,0,71,1101,763,0,72,1105,1,73,1,1037327,8,179734,1101,57977,0,66,1101,0,1,67,1101,792,0,68,1102,556,1,69,1101,0,0,71,1101,0,794,72,1105,1,73,1,1430,1102,57173,1,66,1101,0,1,67,1101,0,821,68,1102,1,556,69,1101,0,0,71,1101,0,823,72,1106,0,73,1,1128,1101,0,37573,66,1102,1,3,67,1102,1,850,68,1102,253,1,69,1102,1,1,71,1101,856,0,72,1105,1,73,0,0,0,0,0,0,17,209158,1101,0,17569,66,1102,1,1,67,1101,0,885,68,1102,556,1,69,1101,2,0,71,1101,887,0,72,1106,0,73,1,8329,20,85972,34,54146,1101,84653,0,66,1101,1,0,67,1101,0,918,68,1102,1,556,69,1102,1,1,71,1101,0,920,72,1105,1,73,1,3167,20,64479,1101,0,88657,66,1102,1,1,67,1101,0,949,68,1101,556,0,69,1102,0,1,71,1102,1,951,72,1105,1,73,1,1419,1101,0,104579,66,1101,2,0,67,1101,0,978,68,1101,0,302,69,1101,0,1,71,1102,982,1,72,1105,1,73,0,0,0,0,6,76316,1102,33797,1,66,1101,0,1,67,1102,1011,1,68,1101,556,0,69,1102,1,3,71,1101,0,1013,72,1105,1,73,1,5,35,13306,35,26612,39,113194,1102,75781,1,66,1101,0,1,67,1102,1,1046,68,1102,1,556,69,1101,0,1,71,1101,1048,0,72,1106,0,73,1,1777,20,42986,1101,32183,0,66,1102,1,1,67,1102,1,1077,68,1102,556,1,69,1102,1,2,71,1102,1,1079,72,1105,1,73,1,10,35,6653,39,169791,1101,19079,0,66,1101,0,4,67,1102,1,1110,68,1102,253,1,69,1101,0,1,71,1102,1118,1,72,1105,1,73,0,0,0,0,0,0,0,0,7,7877,1102,1,56893,66,1101,0,4,67,1102,1147,1,68,1101,253,0,69,1102,1,1,71,1101,0,1155,72,1105,1,73,0,0,0,0,0,0,0,0,45,73189,1102,104651,1,66,1101,1,0,67,1102,1,1184,68,1101,556,0,69,1101,0,0,71,1101,0,1186,72,1105,1,73,1,1351,1101,0,43661,66,1101,1,0,67,1102,1213,1,68,1101,0,556,69,1102,1,1,71,1101,1215,0,72,1105,1,73,1,763676,28,75146,1101,16963,0,66,1102,1,1,67,1102,1,1244,68,1102,1,556,69,1102,1,1,71,1102,1246,1,72,1106,0,73,1,11,26,163118,1101,461,0,66,1101,1,0,67,1102,1275,1,68,1101,556,0,69,1101,0,0,71,1102,1,1277,72,1106,0,73,1,1263,1102,101483,1,66,1101,1,0,67,1102,1304,1,68,1102,556,1,69,1102,0,1,71,1102,1306,1,72,1106,0,73,1,1607,1101,0,14251,66,1101,0,3,67,1102,1333,1,68,1101,302,0,69,1102,1,1,71,1102,1,1339,72,1105,1,73,0,0,0,0,0,0,2,170679,1101,0,27073,66,1101,0,4,67,1102,1,1368,68,1102,1,302,69,1101,1,0,71,1101,0,1376,72,1106,0,73,0,0,0,0,0,0,0,0,1,154743,1101,0,15583,66,1101,0,1,67,1102,1,1405,68,1102,556,1,69,1101,0,3,71,1101,0,1407,72,1106,0,73,1,1,33,14998,21,302361,34,81219,1101,0,56597,66,1101,0,6,67,1102,1440,1,68,1102,1,302,69,1102,1,1,71,1101,1452,0,72,1106,0,73,0,0,0,0,0,0,0,0,0,0,0,0,45,146378,1102,6991,1,66,1101,0,1,67,1102,1,1481,68,1102,1,556,69,1101,0,1,71,1102,1,1483,72,1106,0,73,1,902,21,201574,1102,1,89867,66,1101,2,0,67,1101,1512,0,68,1101,0,302,69,1101,1,0,71,1102,1516,1,72,1106,0,73,0,0,0,0,5,10798,1102,1,61553,66,1102,1,1,67,1102,1,1545,68,1102,556,1,69,1101,0,6,71,1101,0,1547,72,1106,0,73,1,2,20,21493,17,104579,7,15754,26,244677,39,282985,39,339582,1101,73189,0,66,1102,2,1,67,1101,0,1586,68,1101,0,351,69,1102,1,1,71,1101,0,1590,72,1105,1,73,0,0,0,0,255,60251,1101,0,91463,66,1102,1,1,67,1101,1619,0,68,1101,556,0,69,1102,1,1,71,1101,0,1621,72,1105,1,73,1,6467743,28,37573,1102,1,34369,66,1101,0,1,67,1101,1650,0,68,1102,556,1,69,1102,1,1,71,1101,1652,0,72,1105,1,73,1,8933,21,100787,1102,1,77081,66,1101,0,1,67,1101,1681,0,68,1102,1,556,69,1102,1,1,71,1101,1683,0,72,1105,1,73,1,2310,33,7499,1101,0,44699,66,1101,0,1,67,1102,1712,1,68,1101,556,0,69,1101,3,0,71,1101,1714,0,72,1105,1,73,1,3,8,89867,5,5399,34,108292,1101,97387,0,66,1102,1,1,67,1102,1747,1,68,1102,556,1,69,1102,1,1,71,1102,1749,1,72,1105,1,73,1,4,20,107465,1101,0,13463,66,1101,1,0,67,1101,1778,0,68,1101,0,556,69,1102,0,1,71,1102,1,1780,72,1106,0,73,1,1250,1102,51581,1,66,1102,1,3,67,1102,1807,1,68,1101,302,0,69,1101,0,1,71,1102,1,1813,72,1105,1,73,0,0,0,0,0,0,2,113786,1102,1,73679,66,1102,1,1,67,1102,1842,1,68,1102,556,1,69,1101,1,0,71,1102,1,1844,72,1105,1,73,1,160,39,56597,1101,0,60251,66,1101,0,1,67,1101,1873,0,68,1102,556,1,69,1102,6,1,71,1101,1875,0,72,1106,0,73,1,22677,47,94099,1,51581,1,103162,9,14251,9,28502,9,42753,1101,0,81559,66,1101,0,3,67,1102,1914,1,68,1101,0,302,69,1101,1,0,71,1102,1920,1,72,1105,1,73,0,0,0,0,0,0,47,188198,1101,0,6653,66,1101,4,0,67,1102,1949,1,68,1101,302,0,69,1101,0,1,71,1101,0,1957,72,1106,0,73,0,0,0,0,0,0,0,0,39,226388,1101,63659,0,66,1101,0,1,67,1101,0,1986,68,1102,556,1,69,1102,1,1,71,1102,1,1988,72,1105,1,73,1,-2,34,27073,1101,94099,0,66,1101,2,0,67,1101,0,2017,68,1101,0,302,69,1102,1,1,71,1102,2021,1,72,1106,0,73,0,0,0,0,2,227572,1101,100787,0,66,1101,0,3,67,1101,2050,0,68,1102,1,302,69,1101,1,0,71,1101,0,2056,72,1105,1,73,0,0,0,0,0,0,6,19079,1101,0,47599,66,1102,1,1,67,1101,0,2085,68,1102,556,1,69,1102,1,1,71,1101,2087,0,72,1105,1,73,1,2677,33,22497,1102,1,7499,66,1101,0,3,67,1101,2116,0,68,1102,1,302,69,1101,0,1,71,1101,0,2122,72,1105,1,73,0,0,0,0,0,0,6,38158,1101,0,101359,66,1102,1,1,67,1101,2151,0,68,1102,1,556,69,1102,1,0,71,1102,1,2153,72,1105,1,73,1,1752,1101,48091,0,66,1101,1,0,67,1102,1,2180,68,1101,0,556,69,1101,0,0,71,1101,0,2182,72,1106,0,73,1,1463,1102,5399,1,66,1101,0,2,67,1102,1,2209,68,1101,0,302,69,1102,1,1,71,1101,0,2213,72,1106,0,73,0,0,0,0,6,57237"

func handlePacketSendThread(packetChannel chan packet, NATchannel chan packet) {
	for p := range packetChannel {
		// fmt.Println(p.x, p.y)
		if p.recipient == NAT {
			NATchannel <- p
		} else {
			p.recipient.packetMutex.Lock()
			p.recipient.packetQueue = append(p.recipient.packetQueue, p)
			p.recipient.packetMutex.Unlock()
		}
	}
}

func allIdle(packetChannel chan packet) bool {
	for _, c := range computers {
		if len(c.packetQueue) > 0 || len(packetChannel) > 0 {
			println("false")
			return false
		}
	}
	println("true")
	return true
}
func handleIdle(idleChannel chan bool, packetChannel chan packet) {
	for {
		if allIdle(packetChannel) {
			timer := time.NewTimer(50 * time.Millisecond)
			idleChannel <- true
			<-timer.C
		}
	}
}

func (c *computer) NAT(NATChannel chan packet, idleChannel chan bool) {
	go func() {
		prevSent := -1
		for range idleChannel {
			computers[0].packetMutex.Lock()
			if len(c.packetQueue) == 0 {
				computers[0].packetMutex.Unlock()
				continue
			}
			// fmt.Println(c.packetQueue[0])
			curSent := c.packetQueue[0].y
			// c.packetMutex.Lock()
			// c.packetQueue = make([]packet, 0)
			// c.packetMutex.Unlock()
			fmt.Println(curSent, prevSent)
			if curSent == prevSent {
				panic(curSent)
			}
			prevSent = curSent
			computers[0].packetQueue = append(computers[0].packetQueue, c.packetQueue[0])
			computers[0].packetMutex.Unlock()
		}
	}()
	for p := range NATChannel {
		p.recipient = computers[0]
		c.packetMutex.Lock()
		c.packetQueue = []packet{p}
		c.packetMutex.Unlock()
	}
}

func (c *computer) computeThread(packetChannel chan packet) {
	command := c.command
	length := len(command)
	base := c.base
	curPacketToSend := packet{c, nil, -1, -1}
	receivingPacket := false
	firstInput := true
	for i := 0; i < length; {
		// println(i, command[i])
		numStr := strconv.Itoa(command[i])
		for len(numStr) < 5 {
			numStr = "0" + numStr
		}
		opcode, _ := strconv.Atoi(numStr[3:])
		paramModes := []byte{numStr[2], numStr[1], numStr[0]}
		// fmt.Println(string(paramModes[0]), string(paramModes[1]), string(paramModes[2]))
		var value1, value2 int
		if opcode == 99 {
			fmt.Println(opcode)
			break
		}

		if paramModes[0] == '1' {
			value1 = command[i+1]
		} else if paramModes[0] == '2' {
			value1 = command[base+command[i+1]]
		} else { //if opcode != 3 && opcode != 9 {

			value1 = command[command[i+1]]
		}
		if paramModes[1] == '1' {
			value2 = command[i+2]
		} else if paramModes[1] == '2' {
			value2 = command[base+command[i+2]]
		} else if opcode != 3 && opcode != 9 {
			value2 = command[command[i+2]]
		}
		// fmt.Println(opcode)
		switch opcode {
		case 1: //addition
			// println(numStr, command[i+3], string(paramModes[2]))
			// println(command[command[i+3]])
			if paramModes[2] == '2' {
				command[base+command[i+3]] = value1 + value2
			} else {
				command[command[i+3]] = value1 + value2
			}
			i += 4
		case 2: //multiplication
			if paramModes[2] == '2' {
				command[base+command[i+3]] = value1 * value2
			} else {
				command[command[i+3]] = value1 * value2
			}
			i += 4
		case 3: //input
			// if firstInput == -1 {
			// 	command[command[i+1]] = secondInput
			// } else {
			// 	command[command[i+1]] = firstInput
			// 	firstInput = -1
			// }
			value := -1
			c.packetMutex.Lock()
			if len(c.packetQueue) != 0 && !firstInput {

				if receivingPacket {
					value = c.packetQueue[0].y
					receivingPacket = false
					c.packetQueue = c.packetQueue[1:]
				} else {
					receivingPacket = true
					value = c.packetQueue[0].x
				}
			}
			if firstInput {
				value = c.address
				firstInput = false
			}
			c.packetMutex.Unlock()

			if paramModes[0] == '2' {
				command[base+command[i+1]] = value
			} else {
				command[command[i+1]] = value
			}

			i += 2
		case 4: //output

			if curPacketToSend.recipient == nil {
				if value1 == 255 {
					curPacketToSend.recipient = NAT
				} else {
					curPacketToSend.recipient = computers[value1]
				}
			} else if curPacketToSend.x == -1 {
				curPacketToSend.x = value1
			} else {
				curPacketToSend.y = value1
				packetChannel <- curPacketToSend
				curPacketToSend = packet{c, nil, -1, -1}
			}

			// if rob.x > 100 || rob.y > 100 ||
			// 	rob.x < -100 || rob.y < -100 {
			// 	println("break")
			// 	return
			// }
			// bre := false
			i += 2
		case 5:
			if value1 != 0 {
				i = value2
			} else {
				i += 3
			}
		case 6:
			if value1 == 0 {
				i = value2
			} else {
				i += 3
			}
		case 7:
			if value1 < value2 {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 1
				} else {
					command[command[i+3]] = 1
				}
			} else {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 0
				} else {
					command[command[i+3]] = 0
				}
			}
			i += 4
		case 8:
			if value1 == value2 {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 1
				} else {
					command[command[i+3]] = 1
				}
			} else {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 0
				} else {
					command[command[i+3]] = 0
				}
			}
			i += 4
		case 9:
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(command[i+1]))
			// fmt.Printf("%d,%d\n", value1, command[i+1])
			fmt.Printf("%d\n", base)
			base += value1
			// base += command[i+1]
			// fmt.Printf("%d\n", base)
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(value1))
			i += 2
		case 99:
			break
		}

	}
}

func part1() {
	commands := strings.Split(input, ",")
	command := make([]int, len(commands))

	for j, s := range commands {
		command[j], _ = strconv.Atoi(s)
	}
	for i := 0; i < 50; i++ {
		cmd := make([]int, len(command))
		copy(cmd, command)
		cmd = append(cmd, make([]int, 1000)...)
		computers[i] = &computer{i, &sync.Mutex{}, make([]packet, 0), cmd, 0, 0}
	}
	packetChannel, NATChannel, idleChannel := make(chan packet), make(chan packet), make(chan bool)
	for i := range computers {
		go computers[i].computeThread(packetChannel)
	}
	go NAT.NAT(NATChannel, idleChannel)
	go handleIdle(idleChannel, packetChannel)
	handlePacketSendThread(packetChannel, NATChannel)
}
