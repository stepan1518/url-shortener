package services

type LinkService interface {
	GetLink(short_name string) string
	AddLink(link string)
}
