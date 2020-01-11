package main

type pageInterface interface {
	GetPage() Page
}

type Page struct {
	HideFooter bool
	HideNav    bool
}

func (p *Page) GetPage() *Page {
	return p
}
