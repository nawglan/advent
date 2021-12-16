package main

import (
	"fmt"
	"strings"
)

const DEBUG = false

func hex2bin(hex string) string {
	hex2bits := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	binary := ""

	for _, c := range strings.Split(hex, "") {
		binary += hex2bits[c]
	}

	return binary
}

func bin2dec(bin string) int {
	mul := 1
	sum := 0
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			sum += mul
		}
		mul *= 2
	}

	return sum
}

func parsePacketVersion(binary string) (int, string) {
	version := bin2dec(binary[0:3])
	remainder := binary[3:]
	if DEBUG {
		fmt.Printf("Version: %d, remainder: %s\n", version, remainder)
	}
	return version, remainder
}

func parsePacketTypeId(binary string) (int, string) {
	id := bin2dec(binary[0:3])
	remainder := binary[3:]
	if DEBUG {
		fmt.Printf("Type ID: %d, remainder: %s\n", id, remainder)
	}
	return id, remainder
}

func parseLiteralValue(binary string) (int, string) {
	literal := ""
	n := 0
	for n < len(binary) {
		segment := binary[n : n+5]
		literal = literal + segment[1:]
		n += 5
		if segment[0] == '0' {
			break
		}
	}
	remainder := binary[n:]
	value := bin2dec(literal)
	if DEBUG {
		fmt.Printf("Literal Value: %d, remainder: %s\n", value, remainder)
	}
	return value, remainder
}

func parseLengthTypeId(binary string) (int, string) {
	lengthTypeId := bin2dec(binary[0:1])
	remainder := binary[1:]
	if DEBUG {
		fmt.Printf("Length Type ID: %d, remainder: %s\n", lengthTypeId, remainder)
	}
	return lengthTypeId, remainder
}

func parseLength(binary string) (int, string) {
	length := bin2dec(binary[0:15])
	remainder := binary[15:]
	if DEBUG {
		fmt.Printf("Length Count: %d, remainder: %s\n", length, remainder)
	}
	return length, remainder
}

func parseSubPacketCount(binary string) (int, string) {
	count := bin2dec(binary[0:11])
	remainder := binary[11:]
	if DEBUG {
		fmt.Printf("Sub Packet Count: %d, remainder: %s\n", count, remainder)
	}
	return count, remainder
}

type Packet struct {
	version    int
	typeId     int
	literal    int
	subPackets []Packet
}

func parsePacket(binary string) (Packet, string) {
	remainder := ""

	packet := Packet{}
	packet.version, remainder = parsePacketVersion(binary)
	packet.typeId, remainder = parsePacketTypeId(remainder)

	switch packet.typeId {
	case 4: // literal value
		packet.literal, remainder = parseLiteralValue(binary[6:])
	default: // operator packet
		lengthTypeId := 0
		lengthTypeId, remainder = parseLengthTypeId(remainder)
		if lengthTypeId == 0 { // 15 bits
			totalLength := 0
			totalLength, remainder = parseLength(remainder)
			for totalLength > 0 {
				subPacket, remains := parsePacket(remainder)
				packet.subPackets = append(packet.subPackets, subPacket)
				totalLength -= (len(remainder) - len(remains))
				remainder = remains
			}
		} else { // 11 bits
			numSubPackets := 0
			numSubPackets, remainder = parseSubPacketCount(remainder)
			for numSubPackets > 0 {
				subPacket, remains := parsePacket(remainder)
				packet.subPackets = append(packet.subPackets, subPacket)
				remainder = remains
				numSubPackets--
			}
		}
	}

	return packet, remainder
}

func day16(puzzle_data []string) {
	for _, dataVal := range puzzle_data {
		binary := hex2bin(dataVal)

		packet, _ := parsePacket(binary)

		queue := []Packet{packet}
		sum := 0
		for len(queue) > 0 {
			curr := queue[0]
			sum += curr.version
			queue = queue[1:]
			for _, sp := range curr.subPackets {
				queue = append(queue, sp)
			}
		}

		fmt.Printf("Day 16 (part 1): The sum of the packet versions is %d\n", sum)
	}
}
