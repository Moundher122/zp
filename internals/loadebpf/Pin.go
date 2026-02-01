package loadebpf

type Pinner interface {
	Pin(string) error
}

func Pin(m Pinner, path string) error {
	return m.Pin(path)
}
