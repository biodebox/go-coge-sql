package models

import (
	"fmt"
	"strings"
)

type (
	Parser struct {
		PathSep      string
		ListSep      string
		NameSep      string
		PrefixExclud string
		AllSymb      string
	}
)

func NewParser() *Parser {
	return &Parser{
		PathSep:      `.`,
		ListSep:      `|`,
		NameSep:      `:`,
		PrefixExclud: `!`,
		AllSymb:      `*`,
	}
}

func (p *Parser) Parse(expression string) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	raw := strings.Split(expression, p.PathSep)
	switch len(raw) {
	case 1:
		if err := p.parseTable(raw[0], res); err != nil {
			return nil, err
		}
	case 2:
		if err := p.parseSchema(raw[0], res); err != nil {
			return nil, err
		}
		if err := p.parseTable(raw[1], res); err != nil {
			return nil, err
		}
	case 3:
		if err := p.parseDatabase(raw[0], res); err != nil {
			return nil, err
		}
		if err := p.parseSchema(raw[1], res); err != nil {
			return nil, err
		}
		if err := p.parseTable(raw[2], res); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf(`expression is wrong: %s`, expression)
	}

	return res, nil
}

func (p *Parser) parseTable(expression string, res map[string]interface{}) error {
	expression = strings.TrimSpace(expression)
	if expression == p.AllSymb {
		return nil
	}

	if strings.HasPrefix(expression, `(`) {
		all, aliases, included, excluded, err := p.parseList(expression)
		if err != nil {
			return err
		}
		if !all && len(included) > 0 {
			res[`tables`] = included
		}
		if len(excluded) > 0 {
			res[`excluded_tables`] = excluded
		}
		if len(aliases) >= 0 {
			res[`alias_tables`] = aliases
		}
	} else {
		raw := strings.SplitN(expression, `(`, 2)
		if len(raw) > 1 {
			expression = raw[0]
			all, aliases, included, excluded, err := p.parseList(`(`+raw[1])
			if err != nil {
				return err
			}
			if !all && len(included) > 0 {
				res[`columns`] = included
			}
			if len(excluded) > 0 {
				res[`excluded_columns`] = excluded
			}
			if len(aliases) > 0 {
				res[`alias_columns`] = aliases
			}
		}
		res[`tables`] = []string{expression}
	}
	return nil
}

func (p *Parser) parseList(expression string) (bool, map[string]string, []string, []string, error) {
	all := false
	items := strings.Split(expression[1:len(expression)-1], p.ListSep)
	aliases := map[string]string{}
	var included []string
	var excluded []string
	for _, item := range items {
		if item == p.AllSymb {
			all = true
			continue
		}
		if strings.HasPrefix(item, p.PrefixExclud) {
			excluded = append(excluded, item[1:])
			continue
		}
		name, alias, err := p.parseName(item)
		if err != nil {
			return false, nil, nil, nil, err
		}
		if len(alias) > 0 {
			aliases[name] = alias
		}
		included = append(included, name)
	}
	return all, aliases, included, excluded, nil
}

func (p *Parser) parseName(item string) (string, string, error) {
	raw := strings.SplitN(item, p.NameSep, 2)
	if len(raw) == 2 {
		return raw[0], raw[1], nil
	}
	return item, ``, nil
}

func (p *Parser) parseSchema(expression string, res map[string]interface{}) error {
	expression = strings.TrimSpace(expression)
	if expression == p.AllSymb {
		return nil
	}

	if strings.HasPrefix(expression, `(`) {
		all, _, included, excluded, err := p.parseList(expression)
		if err != nil {
			return err
		}
		if !all && len(included) > 0 {
			res[`schemas`] = included
		}
		if len(excluded) > 0 {
			res[`excluded_schemas`] = excluded
		}
	} else {
		res[`schemas`] = []string{expression}
	}
	return nil
}

func (p *Parser) parseDatabase(expression string, res map[string]interface{}) error {
	expression = strings.TrimSpace(expression)
	if expression == p.AllSymb {
		return nil
	}

	if strings.HasPrefix(expression, `(`) {
		all, _, included, excluded, err := p.parseList(expression)
		if err != nil {
			return err
		}
		if !all && len(included) > 0 {
			res[`databases`] = included
		}
		if len(excluded) > 0 {
			res[`excluded_databases`] = excluded
		}
	} else {
		res[`databases`] = []string{expression}
	}
	return nil
}
