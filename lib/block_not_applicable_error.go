package markdown

type BlockNotApplicableError string

func (f BlockNotApplicableError) Error() string {
	return string(f)
}
