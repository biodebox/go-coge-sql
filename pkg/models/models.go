package models

import (
	"fmt"
	"io"
	"strings"
)

//go:generate mockery -all
///go:generate coge sql model rutrackerorg_setction:models(section_id:ID|!version|*)
type (
	RequestModels struct {
		Config    string
		Arguments []string
	}
	Column struct {
		Name    string
		Type    string
		Options map[string]string
	}
	Table struct {
		DataBaseName string
		SchemaName   string
		Name         string
		Columns      []*Column
	}
	Transformer interface {
		Transform() ([]*Table, error)
	}
	Hydrator interface {
		Hydrate(table *Table, writer io.Writer) error
	}
	Configurator interface {
		GetTransformer() Transformer
		GetHydrator() Hydrator
	}
	ConfiguratorFactory func(options string, params map[string]interface{}) (Configurator, error)
	Generator           struct {
		Configurators map[string]ConfiguratorFactory
	}
)

func (g *Generator) Generate(model *RequestModels) error {
	params, err := model.Params()
	if err != nil {
		return err
	}
	name, options, err := model.Configs()
	if err != nil {
		return err
	}
	configurator, err := g.GetConfigurator(name, options, params)
	if err != nil {
		return err
	}
	_ = configurator
	return nil
}

func (r *RequestModels) Params() (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (r *RequestModels) Configs() (name, option string, err error) {
	config := strings.TrimSpace(r.Config)
	if len(config) != 0 {
		name = config
		option = ``
		ss := strings.SplitAfterN(config, `:`, 2)
		if len(ss) == 2 {
			name = strings.TrimSpace(strings.Trim(ss[0], `:`))
			option = strings.TrimSpace(ss[1])
		}
	}
	return
}

func (g *Generator) GetConfigurator(name, option string, params map[string]interface{}) (Configurator, error) {
	if factory, ok := g.Configurators[name]; ok {
		return factory(option, params)
	}

	return nil, fmt.Errorf(`wrong config: %s`, name)
}
