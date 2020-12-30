package routes

// RouteMatch is string enum containing the variations of matching that are capable for a route and url path
type RouteMatch string

const (
	FullMatch    RouteMatch = "Full"
	PartialMatch RouteMatch = "Partial" // Wild card matching
	NoMatch      RouteMatch = "None"
)
