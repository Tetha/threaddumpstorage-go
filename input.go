package input

//import "errors"
import "unicode/utf8"

type Input struct {
    content string;
    position int;
    marks []int;
}

func CreateInput(content string) (r Input) {
    r = Input{content: content, position: 0}
    r.marks = make([]int, 10)
    return
}

func (input *Input) Current() rune {
    runeValue, _ := utf8.DecodeRuneInString(input.content[input.position:])
    return runeValue
}

func (input *Input) Advance() {
    _, width := utf8.DecodeRuneInString(input.content[input.position:])
    input.position += width
}

func (input *Input) Mark() {
    input.marks = append(input.marks, input.position)
}

func (input *Input) Rollback() (err error) {
    lastPosition := -1
    lastPosition, input.marks = input.marks[len(input.marks)-1], input.marks[:len(input.marks)-1]
    input.position = lastPosition
    return
}
