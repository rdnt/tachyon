package socket2

type Hub interface {
	CreateClient() (Client, func(), error)
	Clients() ([]Client, error)
	Client(id int) (Client, error)
	DeleteClient(id int) error
}

type Hubimpl struct {
	clients   map[int]Client
	clientIdx int
}

func (h Hubimpl) CreateClient() (Client, func(), error) {
	//TODO implement me
	panic("implement me")
}

func (h Hubimpl) Clients() ([]Client, error) {
	//TODO implement me
	panic("implement me")
}

func (h Hubimpl) DeleteClient(id string) error {
	//TODO implement me
	panic("implement me")
}

func New() Hub {
	return Hubimpl{}
}

type Client interface {
	Send(b []byte) error
	Messages() (chan []byte, error)
}
