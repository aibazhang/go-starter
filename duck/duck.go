package duck

// Duck
type DuckI interface {
	quack() string
}

type duck struct{}

func NewDuck() DuckI {
	return &duck{}
}

func (d *duck) quack() string {
	return "QUUAAAACCCCKKKKK!!!!!!"
}

// Self Calling Duck
type selfCallingDuck struct {
	name string
}

func NewNameDuck(name string) DuckI {
	return &selfCallingDuck{name}
}

func (d *selfCallingDuck) quack() string {
	return d.name
}
