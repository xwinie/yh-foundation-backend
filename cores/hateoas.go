package cores

import (
	"encoding/json"
	"path"
)

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
	Title  string `json:"title"`
}

func (l *Link) Slash(s string) *Link {
	l.Href = path.Join(l.Href, s)
	return l
}

func (l *Link) WithRel(s string) *Link {
	l.Rel = s
	return l
}
func (l *Link) WithMethod(s string) *Link {
	l.Method = s
	return l
}
func (l *Link) WithTitle(s string) *Link {
	l.Title = s
	return l
}

type Links struct {
	links []*Link
}

func (l *Links) Add(link *Link) *Links {
	if link != nil {
		l.links = append(l.links, link)
	}
	return l
}

func (l *Links) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.links)
}

type Hateoas struct {
	Links Links `json:"_links"`
}

func (h *Hateoas) AddLinks(l Links) *Hateoas {
	h.Links = l
	return h
}

type HateoasTemplate struct {
	Links Links `json:"_link_template"`
}

func (h *HateoasTemplate) AddLinks(l Links) *HateoasTemplate {
	h.Links = l
	return h
}
func LinkTo(href string, rel string, method string, title string) *Link {
	if href != "" {
		return &Link{Rel: rel, Href: href, Method: method, Title: title}
	}
	return &Link{Rel: "self", Href: "/"}
}
