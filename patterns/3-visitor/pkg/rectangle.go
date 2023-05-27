package pkg

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}
