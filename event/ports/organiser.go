package ports

type OrganiserPort interface {
	CreateOrganiser(eventId uint64, userid uint64) error
}
