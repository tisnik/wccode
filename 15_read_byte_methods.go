package main

type X struct {
}

type Y struct {
}

type Z struct {
}

func (x X) ReadByte() (byte, error) {
	return 0, nil
}

func (y Y) ReadByte() error {
	return nil
}

func (z Z) ReadByte() byte {
	return 0
}

func main() {
}
