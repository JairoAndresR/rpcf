package managers

import (
	"fmt"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"

	"github.com/JairoAndresR/collector"
)

func init() {
	err := core.Injector.Provide(newAuthorCollector)
	core.CheckInjection(err, "newAuthorCollector")
}

type authorCollector struct {
	parser                 collector.Parser
	authorWriter           ports.AuthorsWriter
	authorDefinitionReader ports.AuthorsDefinitionReader
	linker                 ports.AuthorsLinker
}

func newAuthorCollector(
	adr ports.AuthorsDefinitionReader,
	aw ports.AuthorsWriter,
	linker ports.AuthorsLinker,
) ports.AuthorsCollector {

	parser := collector.NewParser()
	return &authorCollector{
		parser:                 parser,
		authorWriter:           aw,
		authorDefinitionReader: adr,
		linker:                 linker,
	}
}

func (c *authorCollector) Process(content string) ([]*products.Author, []error) {
	ps, errs := c.Parse(content)

	if len(errs) > 0 {
		return nil, errs
	}

	errors := make([]error, 0)
	authors := make([]*products.Author, 0)

	for _, es := range ps {
		a, err := c.authorWriter.Create(es)
		if err != nil {
			errors = append(errors, err)
		}
		authors = append(authors, a)
	}

	return authors, errors
}

func (c *authorCollector) Parse(content string) ([]*products.Author, []error) {

	payload, err := products.NewContentPayload(content)

	if err != nil {
		return []*products.Author{}, []error{err}
	}

	// Get the author definition
	ad, err := c.authorDefinitionReader.GetAuthorDefinition()
	if err != nil {
		return []*products.Author{}, []error{err}
	}

	errors := make([]error, 0)

	// From the raw definition it will try to create a yml valid
	// and parse to  create definition structure with target fields to be parsed
	// based on regular expressions
	definition, err := collector.NewDefinition(ad.Definition)

	if err != nil {
		errors = append(errors, err)
	}

	// It will try to parse the content based on the target fields and regex expressions
	results, err := c.parser.Parse(definition, payload.Content)

	if err != nil {
		errors = append(errors, err)
	}

	// The result is a []map[field_name]=field value which will be try to generate a
	// []*products.Author
	authors, es := products.NewAuthorsFromResults(results)
	for _, author := range authors {
		fmt.Println("author id", author.ID)
	}
	authors = c.linker.Link(authors, payload.GrupLACCode)
	if len(es) > 0 {
		errors = append(errors, es...)
	}

	return authors, errors
}
