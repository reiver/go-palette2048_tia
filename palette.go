package palette2048_tia

import (
	"github.com/reiver/go-palette2048"
)

// Palette provides the TIA palette.
var Palette [palette2048.ByteSize]uint8 = [palette2048.ByteSize]uint8{
	0x00,0x00,0x00, 255, // 0,0
	0x40,0x40,0x40, 255, // 0,1
	0x6C,0x6C,0x6C, 255, // 0,2
	0x90,0x90,0x90, 255, // 0,3
	0xB0,0xB0,0xB0, 255, // 0,4
	0xC8,0xC8,0xC8, 255, // 0,5
	0xDC,0xDC,0xDC, 255, // 0,6
	0xEC,0xEC,0xEC, 255, // 0,7

	0x44,0x44,0x00, 255, // 1,0
	0x64,0x64,0x10, 255, // 1,1
	0x84,0x84,0x24, 255, // 1,2
	0xA0,0xA0,0x34, 255, // 1,3
	0xB8,0xB8,0x40, 255, // 1,4
	0xD0,0xD0,0x50, 255, // 1,5
	0xE8,0xE8,0x5C, 255, // 1,6
	0xFC,0xFC,0x68, 255, // 1,7

	0x70,0x28,0x00, 255, // 2,0
	0x84,0x44,0x14, 255, // 2,1
	0x98,0x5C,0x28, 255, // 2,2
	0xAC,0x78,0x3C, 255, // 2,3
	0xBC,0x8C,0x4C, 255, // 2,4
	0xCC,0xA0,0x5C, 255, // 2,5
	0xDC,0xB4,0x68, 255, // 2,6
	0xEC,0xC8,0x78, 255, // 2,7

	0x84,0x18,0x00, 255, // 3,0
	0x98,0x34,0x18, 255, // 3,1
	0xAC,0x50,0x30, 255, // 3,2
	0xC0,0x68,0x48, 255, // 3,3
	0xD0,0x80,0x5C, 255, // 3,4
	0xE0,0x94,0x70, 255, // 3,5
	0xEC,0xA8,0x80, 255, // 3,6
	0xFC,0xBC,0x94, 255, // 3,7

	0x88,0x00,0x00, 255, // 4,0
	0x9C,0x20,0x20, 255, // 4,1
	0xB0,0x3C,0x3C, 255, // 4,2
	0xC0,0x58,0x58, 255, // 4,3
	0xD0,0x70,0x70, 255, // 4,4
	0xE0,0x88,0x88, 255, // 4,5
	0xEC,0xA0,0xA0, 255, // 4,6
	0xFC,0xB4,0xB4, 255, // 4,7

	0x78,0x00,0x5C, 255, // 5,0
	0x8C,0x20,0x74, 255, // 5,1
	0xA0,0x3C,0x88, 255, // 5,2
	0xB0,0x58,0x9C, 255, // 5,3
	0xC0,0x70,0xB0, 255, // 5,4
	0xD0,0x84,0xC0, 255, // 5,5
	0xDC,0x9C,0xD0, 255, // 5,6
	0xEC,0xB0,0xE0, 255, // 5,7

	0x48,0x00,0x78, 255, // 6,0
	0x60,0x20,0x90, 255, // 6,1
	0x78,0x3C,0xA4, 255, // 6,2
	0x8C,0x58,0xB8, 255, // 6,3
	0xA0,0x70,0xCC, 255, // 6,4
	0xB4,0x84,0xDC, 255, // 6,5
	0xC4,0x9C,0xEC, 255, // 6,6
	0xD4,0xB0,0xFC, 255, // 6,7

	0x14,0x00,0x84, 255, // 7,0
	0x30,0x20,0x98, 255, // 7,1
	0x4C,0x3C,0xAC, 255, // 7,2
	0x68,0x58,0xC0, 255, // 7,3
	0x7C,0x70,0xD0, 255, // 7,4
	0x94,0x88,0xE0, 255, // 7,5
	0xA8,0xA0,0xEC, 255, // 7,6
	0xBC,0xB4,0xFC, 255, // 7,7

	0x00,0x00,0x88, 255, // 8,0
	0x1C,0x20,0x9C, 255, // 8,1
	0x38,0x40,0xB0, 255, // 8,2
	0x50,0x5C,0xC0, 255, // 8,3
	0x68,0x74,0xD0, 255, // 8,4
	0x7C,0x8C,0xE0, 255, // 8,5
	0x90,0xA4,0xEC, 255, // 8,6
	0xA4,0xB8,0xFC, 255, // 8,7

	0x00,0x18,0x7C, 255, // 9,0
	0x1C,0x38,0x90, 255, // 9,1
	0x38,0x54,0xA8, 255, // 9,2
	0x50,0x70,0xBC, 255, // 9,3
	0x68,0x88,0xCC, 255, // 9,4
	0x7C,0x9C,0xDC, 255, // 9,5
	0x90,0xB4,0xEC, 255, // 9,6
	0xA4,0xC8,0xFC, 255, // 9,7

	0x00,0x2C,0x5C, 255, // a,0
	0x1C,0x4C,0x78, 255, // a,1
	0x38,0x68,0x90, 255, // a,2
	0x50,0x84,0xAC, 255, // a,3
	0x68,0x9C,0xC0, 255, // a,4
	0x7C,0xB4,0xD4, 255, // a,5
	0x90,0xCC,0xE8, 255, // a,6
	0xA4,0xE0,0xFC, 255, // a,7

	0x00,0x3C,0x2C, 255, // b,0
	0x1C,0x5C,0x48, 255, // b,1
	0x38,0x7C,0x64, 255, // b,2
	0x50,0x9C,0x80, 255, // b,3
	0x68,0xB4,0x94, 255, // b,4
	0x7C,0xD0,0xAC, 255, // b,5
	0x90,0xE4,0xC0, 255, // b,6
	0xA4,0xFC,0xD4, 255, // b,7

	0x00,0x3C,0x00, 255, // c,0
	0x20,0x5C,0x20, 255, // c,1
	0x40,0x7C,0x40, 255, // c,2
	0x5C,0x9C,0x5C, 255, // c,3
	0x74,0xB4,0x74, 255, // c,4
	0x8C,0xD0,0x8C, 255, // c,5
	0xA4,0xE4,0xA4, 255, // c,6
	0xB8,0xFC,0xB8, 255, // c,7

	0x14,0x38,0x00, 255, // d,0
	0x34,0x5C,0x1C, 255, // d,1
	0x50,0x7C,0x38, 255, // d,2
	0x6C,0x98,0x50, 255, // d,3
	0x84,0xB4,0x68, 255, // d,4
	0x9C,0xCC,0x7C, 255, // d,5
	0xB4,0xE4,0x90, 255, // d,6
	0xC8,0xFC,0xA4, 255, // d,7

	0x2C,0x30,0x00, 255, // e,0
	0x4C,0x50,0x1C, 255, // e,1
	0x68,0x70,0x34, 255, // e,2
	0x84,0x8C,0x4C, 255, // e,3
	0x9C,0xA8,0x64, 255, // e,4
	0xB4,0xC0,0x78, 255, // e,5
	0xCC,0xD4,0x88, 255, // e,6
	0xE0,0xEC,0x9C, 255, // e,7

	0x44,0x28,0x00, 255, // f,0
	0x64,0x48,0x18, 255, // f,1
	0x84,0x68,0x30, 255, // f,2
	0xA0,0x84,0x44, 255, // f,3
	0xB8,0x9C,0x58, 255, // f,4
	0xD0,0xB4,0x6C, 255, // f,5
	0xE8,0xCC,0x7C, 255, // f,6
	0xFC,0xE0,0x8C, 255, // f,7
}
