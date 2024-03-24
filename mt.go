package mt

import (
	"bufio"
	"os"
)

type cur_move struct {
	new_state int  // next_state
	new_sym   byte // sym from alphabet
	shift     int  // -1, 0, 1 (L, S, R)
}

type cur_state struct {
	state int
	sym   byte
}

type MT struct {
	state int
	pos   int
	moves map[cur_state]cur_move
}

func (mt *MT) Move(tape []byte) bool {
	next_state := mt.moves[cur_state{mt.state, tape[mt.pos]}]
	tape[mt.pos] = next_state.new_sym
	mt.state = next_state.new_state
	mt.pos += next_state.shift
	if mt.state == -1 && mt.pos != 0 {
		panic("Конечное состояние достигнуто не в начале ленты")
	}
	return mt.state != -1
}

func (mt *MT) add_move(state int, sym byte, new_state int, new_sym byte, str_shift byte) {
	convert := map[byte]int{'L': -1, 'S': 0, 'R': 1, 'l': -1, 's': 0, 'r': 1}
	shift := convert[str_shift]
	mt.moves[cur_state{state, sym}] = cur_move{new_state, new_sym, shift}
}

func (mt *MT) Read_file_moves(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cur_string := scanner.Text()
		state := int(cur_string[0]) - '0'
		sym := cur_string[1]
		new_state := int(cur_string[3]) - '0'
		new_sym := cur_string[4]
		str_shift := cur_string[5]
		mt.add_move(state, sym, new_state, new_sym, str_shift)
	}
}

func Create_mt() *MT {
	mt := new(MT)
	mt.state = 0
	mt.pos = 0
	mt.moves = map[cur_state]cur_move{}
	return mt
}
