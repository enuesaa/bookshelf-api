package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

func main() {
	plug.Provide(NewProvider)
}

func NewProvider(db plug.DB, logger plug.Logger) plug.ProviderInterface {
	return &Provider{db, logger}
}

type Provider struct {
	db plug.DB
	logger plug.Logger
}

func (p *Provider) OnStartup() error {
	return p.db.Connect()
}

func (p *Provider) OnShutdown() error {
	return p.db.Close()
}

func (p *Provider) Info() (plug.Info, error) {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
		Teaboxes: []plug.Teabox{
			{
				Name: "links",
				Inputs: []plug.TeaboxInput{
					{Name: "title", Type: plug.TeaboxInputText},
					{Name: "link", Type: plug.TeaboxInputText},
				},
			},
			{
				Name: "notes",
				Inputs: []plug.TeaboxInput{
					{Name: "title", Type: plug.TeaboxInputText},
					{Name: "memo", Type: plug.TeaboxInputText},
				},
			},
		},
		Actions: []plug.Action{
			{
				Name:   "reset",
				Comment: "reset link data",
			},
			{
				Name:   "addSampleData",
				Comment: "add sample data",
			},
		},
	}
	return info, nil
}

func (p *Provider) List(ar plug.ListArgs) ([]db.Tea, error) {
	query := p.db.Use(ar.Teabox)

	return query.FindAll(db.M{}, db.M{})
}

func (p *Provider) Get(ar plug.GetArgs) (db.Tea, error) {
	query := p.db.Use(ar.Teabox)

	return query.Get(ar.TeaId)
}

func (p *Provider) validate(teabox string, data map[string]interface{}) error {
	switch teabox {
	case "links":
		if err := ValidateLinkTea(data); err != nil {
			return err
		}
	case "notes":
		if err := ValidateNoteTea(data); err != nil {
			return err
		}
	}
	return nil
}

func (p *Provider) Create(ar plug.CreateArgs) (string, error) {
	if err := p.validate(ar.Teabox, ar.Data); err != nil {
		return "", err
	}
	return p.db.Use(ar.Teabox).Create(ar.Data)
}

func (p *Provider) Update(ar plug.UpdateArgs) (string, error) {
	if err := p.validate(ar.Teabox, ar.Data); err != nil {
		return "", err
	}
	return p.db.Use(ar.Teabox).Update(ar.TeaId, ar.Data)
}

func (p *Provider) Delete(ar plug.DeleteArgs) error {
	return p.db.Use(ar.Teabox).Delete(ar.TeaId)
}

func (p *Provider) Act(ar plug.ActArgs) (string, error) {
	if ar.Action == "reset" {
		query := p.db.Use("links")

		links, err := query.FindAll(db.M{}, db.M{})
		if err != nil {
			return "err", err
		}
		for _, link := range links {
			if err := query.Delete(link.Id()); err != nil {
				return "err", err
			}
		}
		return "reset links!", nil
	}

	if ar.Action == "addSampleData" {
		query := p.db.Use("links")

		list := []map[string]interface{}{
			{
				"title": "sample1",
				"link": "https://example.com/",
			},
			{
				"title": "sample2",
				"link": "https://example.com/",
			},
		}

		for _, item := range list {
			_, err := query.Create(item)
			if err != nil {
				return "err", err
			}
		}
		return "add sample data!", nil
	}

	return "", nil
}
