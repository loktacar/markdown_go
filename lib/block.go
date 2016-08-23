package markdown

type Block interface {
	//ParseNext(line string) (Block, bool, error)

	Render() string

	//IsEmpty() bool

	//Append(string)
}
