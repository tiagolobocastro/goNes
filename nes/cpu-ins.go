package gones

/*
will be useful for debugging
func getMode(mode uint8) string {
	switch mode {
	case ModeInvalid: return "ModeInvalid"
	case ModeZeroPage: return "ModeZeroPage"
	case ModeIndexedZeroPageX: return "ModeIndexedZeroPageX"
	case ModeIndexedZeroPageY: return "ModeIndexedZeroPageY"
	case ModeAbsolute: return "ModeAbsolute"
	case ModeIndexedAbsoluteX: return "ModeIndexedAbsoluteX"
	case ModeIndexedAbsoluteY: return "ModeIndexedAbsoluteY"
	case ModeIndirect: return "ModeIndirect"
	case ModeImplied: return "ModeImplied"
	case ModeAccumulator: return "ModeAccumulator"
	case ModeImmediate: return "ModeImmediate"
	case ModeRelative: return "ModeRelative"
	case ModeIndexedIndirectX: return "ModeIndexedIndirectX"
	case ModeIndirectIndexedY: return "ModeIndirectIndexedY"
	}
	return "ModeInvalid"
}
*/

func (c *Cpu) setupIns() {
	// illegal operations have a len of 0 for now
	// created using table form http://www.oxyron.de/html/opcodes02.html

	c.addIns2("BRK", 0x00, 1, 7, 0, ModeImplied, c.brk)
	c.addIns2("ORA", 0x01, 2, 6, 0, ModeIndexedIndirectX, c.ora)
	c.addIns("KIL", 0x02, 0, 2, 0, ModeImplied)
	c.addIns("SLO", 0x03, 0, 8, 0, ModeIndexedIndirectX)
	c.addIns2("NOP", 0x04, 2, 3, 0, ModeZeroPage, c.nop)
	c.addIns2("ORA", 0x05, 2, 3, 0, ModeZeroPage, c.ora)
	c.addIns2("ASL", 0x06, 2, 5, 0, ModeZeroPage, c.asl)
	c.addIns("SLO", 0x07, 0, 5, 0, ModeZeroPage)
	c.addIns2("PHP", 0x08, 1, 3, 0, ModeImplied, c.php)
	c.addIns2("ORA", 0x09, 2, 2, 0, ModeImmediate, c.ora)
	c.addIns2("ASL", 0x0a, 1, 2, 0, ModeAccumulator, c.asl)
	c.addIns("ANC", 0x0b, 0, 2, 0, ModeImmediate)
	c.addIns2("NOP", 0x0c, 3, 4, 0, ModeAbsolute, c.nop)
	c.addIns2("ORA", 0x0d, 3, 4, 0, ModeAbsolute, c.ora)
	c.addIns2("ASL", 0x0e, 3, 6, 0, ModeAbsolute, c.asl)
	c.addIns("SLO", 0x0f, 0, 6, 0, ModeAbsolute)
	c.addIns2("BPL", 0x10, 2, 2, 1, ModeRelative, c.bpl)
	c.addIns2("ORA", 0x11, 2, 5, 1, ModeIndirectIndexedY, c.ora)
	c.addIns("KIL", 0x12, 0, 2, 0, ModeImplied)
	c.addIns("SLO", 0x13, 0, 8, 0, ModeIndirectIndexedY)
	c.addIns2("NOP", 0x14, 2, 4, 0, ModeIndexedZeroPageX, c.nop)
	c.addIns2("ORA", 0x15, 2, 4, 0, ModeIndexedZeroPageX, c.ora)
	c.addIns2("ASL", 0x16, 2, 6, 0, ModeIndexedZeroPageX, c.asl)
	c.addIns("SLO", 0x17, 0, 6, 0, ModeIndexedZeroPageX)
	c.addIns2("CLC", 0x18, 1, 2, 0, ModeImplied, c.clc)
	c.addIns2("ORA", 0x19, 3, 4, 1, ModeIndexedAbsoluteY, c.ora)
	c.addIns2("NOP", 0x1a, 1, 2, 0, ModeImplied, c.nop)
	c.addIns("SLO", 0x1b, 0, 7, 0, ModeIndexedAbsoluteY)
	c.addIns2("NOP", 0x1c, 3, 4, 1, ModeIndexedAbsoluteX, c.nop)
	c.addIns2("ORA", 0x1d, 3, 4, 1, ModeIndexedAbsoluteX, c.ora)
	c.addIns2("ASL", 0x1e, 3, 7, 0, ModeIndexedAbsoluteX, c.asl)
	c.addIns("SLO", 0x1f, 0, 7, 0, ModeIndexedAbsoluteX)
	c.addIns2("JSR", 0x20, 3, 6, 0, ModeAbsolute, c.jsr)
	c.addIns2("AND", 0x21, 2, 6, 0, ModeIndexedIndirectX, c.and)
	c.addIns("KIL", 0x22, 0, 2, 0, ModeImplied)
	c.addIns("RLA", 0x23, 0, 8, 0, ModeIndexedIndirectX)
	c.addIns2("BIT", 0x24, 2, 3, 0, ModeZeroPage, c.bit)
	c.addIns2("AND", 0x25, 2, 3, 0, ModeZeroPage, c.and)
	c.addIns2("ROL", 0x26, 2, 5, 0, ModeZeroPage, c.rol)
	c.addIns("RLA", 0x27, 0, 5, 0, ModeZeroPage)
	c.addIns2("PLP", 0x28, 1, 4, 0, ModeImplied, c.plp)
	c.addIns2("AND", 0x29, 2, 2, 0, ModeImmediate, c.and)
	c.addIns2("ROL", 0x2a, 1, 2, 0, ModeAccumulator, c.rol)
	c.addIns("ANC", 0x2b, 0, 2, 0, ModeImmediate)
	c.addIns2("BIT", 0x2c, 3, 4, 0, ModeAbsolute, c.bit)
	c.addIns2("AND", 0x2d, 3, 4, 0, ModeAbsolute, c.and)
	c.addIns2("ROL", 0x2e, 3, 6, 0, ModeAbsolute, c.rol)
	c.addIns("RLA", 0x2f, 0, 6, 0, ModeAbsolute)
	c.addIns2("BMI", 0x30, 2, 2, 1, ModeRelative, c.bmi)
	c.addIns2("AND", 0x31, 2, 5, 1, ModeIndirectIndexedY, c.and)
	c.addIns("KIL", 0x32, 0, 2, 0, ModeImplied)
	c.addIns("RLA", 0x33, 0, 8, 0, ModeIndirectIndexedY)
	c.addIns2("NOP", 0x34, 2, 4, 0, ModeIndexedZeroPageX, c.nop)
	c.addIns2("AND", 0x35, 2, 4, 0, ModeIndexedZeroPageX, c.and)
	c.addIns2("ROL", 0x36, 2, 6, 0, ModeIndexedZeroPageX, c.rol)
	c.addIns("RLA", 0x37, 0, 6, 0, ModeIndexedZeroPageX)
	c.addIns2("SEC", 0x38, 1, 2, 0, ModeImplied, c.sec)
	c.addIns2("AND", 0x39, 3, 4, 1, ModeIndexedAbsoluteY, c.and)
	c.addIns2("NOP", 0x3a, 1, 2, 0, ModeImplied, c.nop)
	c.addIns("RLA", 0x3b, 0, 7, 0, ModeIndexedAbsoluteY)
	c.addIns2("NOP", 0x3c, 3, 4, 1, ModeIndexedAbsoluteX, c.nop)
	c.addIns2("AND", 0x3d, 3, 4, 1, ModeIndexedAbsoluteX, c.and)
	c.addIns2("ROL", 0x3e, 3, 7, 0, ModeIndexedAbsoluteX, c.rol)
	c.addIns("RLA", 0x3f, 0, 7, 0, ModeIndexedAbsoluteX)
	c.addIns2("RTI", 0x40, 1, 6, 0, ModeImplied, c.rti)
	c.addIns2("EOR", 0x41, 2, 6, 0, ModeIndexedIndirectX, c.eor)
	c.addIns("KIL", 0x42, 0, 2, 0, ModeImplied)
	c.addIns("SRE", 0x43, 0, 8, 0, ModeIndexedIndirectX)
	c.addIns2("NOP", 0x44, 2, 3, 0, ModeZeroPage, c.nop)
	c.addIns2("EOR", 0x45, 2, 3, 0, ModeZeroPage, c.eor)
	c.addIns2("LSR", 0x46, 2, 5, 0, ModeZeroPage, c.lsr)
	c.addIns("SRE", 0x47, 0, 5, 0, ModeZeroPage)
	c.addIns2("PHA", 0x48, 1, 3, 0, ModeImplied, c.pha)
	c.addIns2("EOR", 0x49, 2, 2, 0, ModeImmediate, c.eor)
	c.addIns2("LSR", 0x4a, 1, 2, 0, ModeAccumulator, c.lsr)
	c.addIns("ALR", 0x4b, 0, 2, 0, ModeImmediate)
	c.addIns2("JMP", 0x4c, 3, 3, 0, ModeAbsolute, c.jmp)
	c.addIns2("EOR", 0x4d, 3, 4, 0, ModeAbsolute, c.eor)
	c.addIns2("LSR", 0x4e, 3, 6, 0, ModeAbsolute, c.lsr)
	c.addIns("SRE", 0x4f, 0, 6, 0, ModeAbsolute)
	c.addIns2("BVC", 0x50, 2, 2, 1, ModeRelative, c.bvc)
	c.addIns2("EOR", 0x51, 2, 5, 1, ModeIndirectIndexedY, c.eor)
	c.addIns("KIL", 0x52, 0, 2, 0, ModeImplied)
	c.addIns("SRE", 0x53, 0, 8, 0, ModeIndirectIndexedY)
	c.addIns2("NOP", 0x54, 2, 4, 0, ModeIndexedZeroPageX, c.nop)
	c.addIns2("EOR", 0x55, 2, 4, 0, ModeIndexedZeroPageX, c.eor)
	c.addIns2("LSR", 0x56, 2, 6, 0, ModeIndexedZeroPageX, c.lsr)
	c.addIns("SRE", 0x57, 0, 6, 0, ModeIndexedZeroPageX)
	c.addIns2("CLI", 0x58, 1, 2, 0, ModeImplied, c.cli)
	c.addIns2("EOR", 0x59, 3, 4, 1, ModeIndexedAbsoluteY, c.eor)
	c.addIns2("NOP", 0x5a, 1, 2, 0, ModeImplied, c.nop)
	c.addIns("SRE", 0x5b, 0, 7, 0, ModeIndexedAbsoluteY)
	c.addIns2("NOP", 0x5c, 3, 4, 1, ModeIndexedAbsoluteX, c.nop)
	c.addIns2("EOR", 0x5d, 3, 4, 1, ModeIndexedAbsoluteX, c.eor)
	c.addIns2("LSR", 0x5e, 3, 7, 0, ModeIndexedAbsoluteX, c.lsr)
	c.addIns("SRE", 0x5f, 0, 7, 0, ModeIndexedAbsoluteX)
	c.addIns2("RTS", 0x60, 1, 6, 0, ModeImplied, c.rts)
	c.addIns2("ADC", 0x61, 2, 6, 0, ModeIndexedIndirectX, c.adc)
	c.addIns("KIL", 0x62, 0, 2, 0, ModeImplied)
	c.addIns("RRA", 0x63, 0, 8, 0, ModeIndexedIndirectX)
	c.addIns2("NOP", 0x64, 2, 3, 0, ModeZeroPage, c.nop)
	c.addIns2("ADC", 0x65, 2, 3, 0, ModeZeroPage, c.adc)
	c.addIns2("ROR", 0x66, 2, 5, 0, ModeZeroPage, c.ror)
	c.addIns("RRA", 0x67, 0, 5, 0, ModeZeroPage)
	c.addIns2("PLA", 0x68, 1, 4, 0, ModeImplied, c.pla)
	c.addIns2("ADC", 0x69, 2, 2, 0, ModeImmediate, c.adc)
	c.addIns2("ROR", 0x6a, 1, 2, 0, ModeAccumulator, c.ror)
	c.addIns("ARR", 0x6b, 0, 2, 0, ModeImmediate)
	c.addIns2("JMP", 0x6c, 3, 5, 0, ModeIndirect, c.jmp)
	c.addIns2("ADC", 0x6d, 3, 4, 0, ModeAbsolute, c.adc)
	c.addIns2("ROR", 0x6e, 3, 6, 0, ModeAbsolute, c.ror)
	c.addIns("RRA", 0x6f, 0, 6, 0, ModeAbsolute)
	c.addIns2("BVS", 0x70, 2, 2, 1, ModeRelative, c.bvs)
	c.addIns2("ADC", 0x71, 2, 5, 1, ModeIndirectIndexedY, c.adc)
	c.addIns("KIL", 0x72, 0, 2, 0, ModeImplied)
	c.addIns("RRA", 0x73, 0, 8, 0, ModeIndirectIndexedY)
	c.addIns2("NOP", 0x74, 2, 4, 0, ModeIndexedZeroPageX, c.nop)
	c.addIns2("ADC", 0x75, 2, 4, 0, ModeIndexedZeroPageX, c.adc)
	c.addIns2("ROR", 0x76, 2, 6, 0, ModeIndexedZeroPageX, c.ror)
	c.addIns("RRA", 0x77, 0, 6, 0, ModeIndexedZeroPageX)
	c.addIns2("SEI", 0x78, 1, 2, 0, ModeImplied, c.sei)
	c.addIns2("ADC", 0x79, 3, 4, 1, ModeIndexedAbsoluteY, c.adc)
	c.addIns2("NOP", 0x7a, 1, 2, 0, ModeImplied, c.nop)
	c.addIns("RRA", 0x7b, 0, 7, 0, ModeIndexedAbsoluteY)
	c.addIns2("NOP", 0x7c, 3, 4, 1, ModeIndexedAbsoluteX, c.nop)
	c.addIns2("ADC", 0x7d, 3, 4, 1, ModeIndexedAbsoluteX, c.adc)
	c.addIns2("ROR", 0x7e, 3, 7, 0, ModeIndexedAbsoluteX, c.ror)
	c.addIns("RRA", 0x7f, 0, 7, 0, ModeIndexedAbsoluteX)
	c.addIns2("NOP", 0x80, 2, 2, 0, ModeImmediate, c.nop)
	c.addIns2("STA", 0x81, 2, 6, 0, ModeIndexedIndirectX, c.sta)
	c.addIns2("NOP", 0x82, 0, 2, 0, ModeImmediate, c.nop)
	c.addIns("SAX", 0x83, 0, 6, 0, ModeIndexedIndirectX)
	c.addIns2("STY", 0x84, 2, 3, 0, ModeZeroPage, c.sty)
	c.addIns2("STA", 0x85, 2, 3, 0, ModeZeroPage, c.sta)
	c.addIns2("STX", 0x86, 2, 3, 0, ModeZeroPage, c.stx)
	c.addIns("SAX", 0x87, 0, 3, 0, ModeZeroPage)
	c.addIns2("DEY", 0x88, 1, 2, 0, ModeImplied, c.dey)
	c.addIns2("NOP", 0x89, 0, 2, 0, ModeImmediate, c.nop)
	c.addIns2("TXA", 0x8a, 1, 2, 0, ModeImplied, c.txa)
	c.addIns("XAA", 0x8b, 0, 2, 0, ModeImmediate)
	c.addIns2("STY", 0x8c, 3, 4, 0, ModeAbsolute, c.sty)
	c.addIns2("STA", 0x8d, 3, 4, 0, ModeAbsolute, c.sta)
	c.addIns2("STX", 0x8e, 3, 4, 0, ModeAbsolute, c.stx)
	c.addIns("SAX", 0x8f, 0, 4, 0, ModeAbsolute)
	c.addIns2("BCC", 0x90, 2, 2, 1, ModeRelative, c.bcc)
	c.addIns2("STA", 0x91, 2, 6, 0, ModeIndirectIndexedY, c.sta)
	c.addIns("KIL", 0x92, 0, 2, 0, ModeImplied)
	c.addIns("AHX", 0x93, 0, 6, 0, ModeIndirectIndexedY)
	c.addIns2("STY", 0x94, 2, 4, 0, ModeIndexedZeroPageX, c.sty)
	c.addIns2("STA", 0x95, 2, 4, 0, ModeIndexedZeroPageX, c.sta)
	c.addIns2("STX", 0x96, 2, 4, 0, ModeIndexedZeroPageY, c.stx)
	c.addIns("SAX", 0x97, 0, 4, 0, ModeIndexedZeroPageY)
	c.addIns2("TYA", 0x98, 1, 2, 0, ModeImplied, c.tya)
	c.addIns2("STA", 0x99, 3, 5, 0, ModeIndexedAbsoluteY, c.sta)
	c.addIns2("TXS", 0x9a, 1, 2, 0, ModeImplied, c.txs)
	c.addIns("TAS", 0x9b, 0, 5, 0, ModeIndexedAbsoluteY)
	c.addIns("SHY", 0x9c, 0, 5, 0, ModeIndexedAbsoluteX)
	c.addIns2("STA", 0x9d, 3, 5, 0, ModeIndexedAbsoluteX, c.sta)
	c.addIns("SHX", 0x9e, 0, 5, 0, ModeIndexedAbsoluteY)
	c.addIns("AHX", 0x9f, 0, 5, 0, ModeIndexedAbsoluteY)
	c.addIns2("LDY", 0xa0, 2, 2, 0, ModeImmediate, c.ldy)
	c.addIns2("LDA", 0xa1, 2, 6, 0, ModeIndexedIndirectX, c.lda)
	c.addIns2("LDX", 0xa2, 2, 2, 0, ModeImmediate, c.ldx)
	c.addIns("LAX", 0xa3, 0, 6, 0, ModeIndexedIndirectX)
	c.addIns2("LDY", 0xa4, 2, 3, 0, ModeZeroPage, c.ldy)
	c.addIns2("LDA", 0xa5, 2, 3, 0, ModeZeroPage, c.lda)
	c.addIns2("LDX", 0xa6, 2, 3, 0, ModeZeroPage, c.ldx)
	c.addIns("LAX", 0xa7, 0, 3, 0, ModeZeroPage)
	c.addIns2("TAY", 0xa8, 1, 2, 0, ModeImplied, c.tay)
	c.addIns2("LDA", 0xa9, 2, 2, 0, ModeImmediate, c.lda)
	c.addIns2("TAX", 0xaa, 1, 2, 0, ModeImplied, c.tax)
	c.addIns("LAX", 0xab, 0, 2, 0, ModeImmediate)
	c.addIns2("LDY", 0xac, 3, 4, 0, ModeAbsolute, c.ldy)
	c.addIns2("LDA", 0xad, 3, 4, 0, ModeAbsolute, c.lda)
	c.addIns2("LDX", 0xae, 3, 4, 0, ModeAbsolute, c.ldx)
	c.addIns("LAX", 0xaf, 0, 4, 0, ModeAbsolute)
	c.addIns2("BCS", 0xb0, 2, 2, 1, ModeRelative, c.bcs)
	c.addIns2("LDA", 0xb1, 2, 5, 1, ModeIndirectIndexedY, c.lda)
	c.addIns("KIL", 0xb2, 0, 2, 0, ModeImplied)
	c.addIns("LAX", 0xb3, 0, 5, 1, ModeIndirectIndexedY)
	c.addIns2("LDY", 0xb4, 2, 4, 0, ModeIndexedZeroPageX, c.ldy)
	c.addIns2("LDA", 0xb5, 2, 4, 0, ModeIndexedZeroPageX, c.lda)
	c.addIns2("LDX", 0xb6, 2, 4, 0, ModeIndexedZeroPageY, c.ldx)
	c.addIns("LAX", 0xb7, 0, 4, 0, ModeIndexedZeroPageY)
	c.addIns2("CLV", 0xb8, 1, 2, 0, ModeImplied, c.clv)
	c.addIns2("LDA", 0xb9, 3, 4, 1, ModeIndexedAbsoluteY, c.lda)
	c.addIns2("TSX", 0xba, 1, 2, 0, ModeImplied, c.tsx)
	c.addIns("LAS", 0xbb, 0, 4, 1, ModeIndexedAbsoluteY)
	c.addIns2("LDY", 0xbc, 3, 4, 1, ModeIndexedAbsoluteX, c.ldy)
	c.addIns2("LDA", 0xbd, 3, 4, 1, ModeIndexedAbsoluteX, c.lda)
	c.addIns2("LDX", 0xbe, 3, 4, 1, ModeIndexedAbsoluteY, c.ldx)
	c.addIns("LAX", 0xbf, 0, 4, 1, ModeIndexedAbsoluteY)
	c.addIns2("CPY", 0xc0, 2, 2, 0, ModeImmediate, c.cpy)
	c.addIns2("CMP", 0xc1, 2, 6, 0, ModeIndexedIndirectX, c.cmp)
	c.addIns2("NOP", 0xc2, 0, 2, 0, ModeImmediate, c.nop)
	c.addIns("DCP", 0xc3, 0, 8, 0, ModeIndexedIndirectX)
	c.addIns2("CPY", 0xc4, 2, 3, 0, ModeZeroPage, c.cpy)
	c.addIns2("CMP", 0xc5, 2, 3, 0, ModeZeroPage, c.cmp)
	c.addIns2("DEC", 0xc6, 2, 5, 0, ModeZeroPage, c.dec)
	c.addIns("DCP", 0xc7, 0, 5, 0, ModeZeroPage)
	c.addIns2("INY", 0xc8, 1, 2, 0, ModeImplied, c.iny)
	c.addIns2("CMP", 0xc9, 2, 2, 0, ModeImmediate, c.cmp)
	c.addIns2("DEX", 0xca, 1, 2, 0, ModeImplied, c.dex)
	c.addIns("AXS", 0xcb, 0, 2, 0, ModeImmediate)
	c.addIns2("CPY", 0xcc, 3, 4, 0, ModeAbsolute, c.cpy)
	c.addIns2("CMP", 0xcd, 3, 4, 0, ModeAbsolute, c.cmp)
	c.addIns2("DEC", 0xce, 3, 6, 0, ModeAbsolute, c.dec)
	c.addIns("DCP", 0xcf, 0, 6, 0, ModeAbsolute)
	c.addIns2("BNE", 0xd0, 2, 2, 1, ModeRelative, c.bne)
	c.addIns2("CMP", 0xd1, 2, 5, 1, ModeIndirectIndexedY, c.cmp)
	c.addIns("KIL", 0xd2, 0, 2, 0, ModeImplied)
	c.addIns("DCP", 0xd3, 0, 8, 0, ModeIndirectIndexedY)
	c.addIns2("NOP", 0xd4, 2, 4, 0, ModeIndexedZeroPageX, c.nop)
	c.addIns2("CMP", 0xd5, 2, 4, 0, ModeIndexedZeroPageX, c.cmp)
	c.addIns2("DEC", 0xd6, 2, 6, 0, ModeIndexedZeroPageX, c.dec)
	c.addIns("DCP", 0xd7, 0, 6, 0, ModeIndexedZeroPageX)
	c.addIns2("CLD", 0xd8, 1, 2, 0, ModeImplied, c.cld)
	c.addIns2("CMP", 0xd9, 3, 4, 1, ModeIndexedAbsoluteY, c.cmp)
	c.addIns2("NOP", 0xda, 1, 2, 0, ModeImplied, c.nop)
	c.addIns("DCP", 0xdb, 0, 7, 0, ModeIndexedAbsoluteY)
	c.addIns2("NOP", 0xdc, 3, 4, 1, ModeIndexedAbsoluteX, c.nop)
	c.addIns2("CMP", 0xdd, 3, 4, 1, ModeIndexedAbsoluteX, c.cmp)
	c.addIns2("DEC", 0xde, 3, 7, 0, ModeIndexedAbsoluteX, c.dec)
	c.addIns("DCP", 0xdf, 0, 7, 0, ModeIndexedAbsoluteX)
	c.addIns2("CPX", 0xe0, 2, 2, 0, ModeImmediate, c.cpx)
	c.addIns2("SBC", 0xe1, 2, 6, 0, ModeIndexedIndirectX, c.sbc)
	c.addIns2("NOP", 0xe2, 0, 2, 0, ModeImmediate, c.nop)
	c.addIns("ISB", 0xe3, 0, 8, 0, ModeIndexedIndirectX)
	c.addIns2("CPX", 0xe4, 2, 3, 0, ModeZeroPage, c.cpx)
	c.addIns2("SBC", 0xe5, 2, 3, 0, ModeZeroPage, c.sbc)
	c.addIns2("INC", 0xe6, 2, 5, 0, ModeZeroPage, c.inc)
	c.addIns("ISB", 0xe7, 0, 5, 0, ModeZeroPage)
	c.addIns2("INX", 0xe8, 1, 2, 0, ModeImplied, c.inx)
	c.addIns2("SBC", 0xe9, 2, 2, 0, ModeImmediate, c.sbc)
	c.addIns2("NOP", 0xea, 1, 2, 0, ModeImplied, c.nop)
	c.addIns2("SBC", 0xeb, 0, 2, 0, ModeImmediate, c.sbc)
	c.addIns2("CPX", 0xec, 3, 4, 0, ModeAbsolute, c.cpx)
	c.addIns2("SBC", 0xed, 3, 4, 0, ModeAbsolute, c.sbc)
	c.addIns2("INC", 0xee, 3, 6, 0, ModeAbsolute, c.inc)
	c.addIns("ISB", 0xef, 0, 6, 0, ModeAbsolute)
	c.addIns2("BEQ", 0xf0, 2, 2, 1, ModeRelative, c.beq)
	c.addIns2("SBC", 0xf1, 2, 5, 1, ModeIndirectIndexedY, c.sbc)
	c.addIns("KIL", 0xf2, 0, 2, 0, ModeImplied)
	c.addIns("ISB", 0xf3, 0, 8, 0, ModeIndirectIndexedY)
	c.addIns2("NOP", 0xf4, 2, 4, 0, ModeIndexedZeroPageX, c.nop)
	c.addIns2("SBC", 0xf5, 2, 4, 0, ModeIndexedZeroPageX, c.sbc)
	c.addIns2("INC", 0xf6, 2, 6, 0, ModeIndexedZeroPageX, c.inc)
	c.addIns("ISB", 0xf7, 0, 6, 0, ModeIndexedZeroPageX)
	c.addIns2("SED", 0xf8, 1, 2, 0, ModeImplied, c.sed)
	c.addIns2("SBC", 0xf9, 3, 4, 1, ModeIndexedAbsoluteY, c.sbc)
	c.addIns2("NOP", 0xfa, 1, 2, 0, ModeImplied, c.nop)
	c.addIns("ISB", 0xfb, 0, 7, 0, ModeIndexedAbsoluteY)
	c.addIns2("NOP", 0xfc, 3, 4, 1, ModeIndexedAbsoluteX, c.nop)
	c.addIns2("SBC", 0xfd, 3, 4, 1, ModeIndexedAbsoluteX, c.sbc)
	c.addIns2("INC", 0xfe, 3, 7, 0, ModeIndexedAbsoluteX, c.inc)
	c.addIns("ISB", 0xff, 0, 7, 0, ModeIndexedAbsoluteX)
}

func (c *Cpu) unhandled() {
	message := "oops... unhandled instruction!\n"
	c.Logf(message)
	panic(message)
}

func (c *Cpu) addIns(opName string, opCode uint8, opLength uint8, opCycles uint8, opPageCycles uint8, addrMode uint8) {
	c.ins[opCode] = Instruction{opLength, opCycles, opPageCycles, addrMode,
		opCode, opName, c.unhandled, false}
}
func (c *Cpu) addIns2(opName string, opCode uint8, opLength uint8, opCycles uint8, opPageCycles uint8, addrMode uint8, f func()) {
	c.ins[opCode] = Instruction{opLength, opCycles, opPageCycles, addrMode,
		opCode, opName, f, true}
}
