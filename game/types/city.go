package types

type City string

func (c City) Name() string {
	return string(c)
}
