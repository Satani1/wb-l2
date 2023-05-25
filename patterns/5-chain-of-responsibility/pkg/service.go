package pkg

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}
