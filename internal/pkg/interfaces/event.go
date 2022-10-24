package interfaces

type Event interface{}

type Type string

const (
	UserCreated    Type = "user_created"
	ProjectCreated Type = "project_created"
	SessionCreated Type = "session_created"
	JoinedSession  Type = "joined_session"
	LeftSession    Type = "left_session"
	PathCreated    Type = "path_created"
	PixelDrawn     Type = "pixel_drawn"
	PixelErased    Type = "pixel_erased"
)

type AggregateType string

const (
	Session AggregateType = "session"
	User    AggregateType = "user"
	Project AggregateType = "project"
)
