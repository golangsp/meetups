//show code1 OMIT
func (b *board) BoardRepresentation() string {
    var buffer = &bytes.Buffer{}
    for _, l := range b.squares() {
        buffer.WriteString(l.current.representation[0:1])
    }
    return buffer.String()
}
//end show code1 OMIT

//show code2 OMIT
func (b *board) BoardRepresentation() string {
    var buffer = &bytes.Buffer{}
    for _, l := range b.squares() {
        l.addTo(buffer)
    }
    return buffer.String()
}
//end show code2 OMIT