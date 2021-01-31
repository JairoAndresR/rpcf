package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"

	"github.com/udistritali3plus/collector"
)

func init() {
	err := core.Injector.Provide(newAuthorCollector)
	core.CheckInjection(err, "newAuthorCollector")
}

type authorCollector struct {
	parser                 collector.Parser
	authorWriter           ports.AuthorsWriter
	authorDefinitionReader ports.AuthorsDefinitionReader
}

func newAuthorCollector(
	adr ports.AuthorsDefinitionReader,
	aw ports.AuthorsWriter,
) ports.AuthorsCollector {

	parser := collector.NewParser()
	return &authorCollector{
		parser:                 parser,
		authorWriter:           aw,
		authorDefinitionReader: adr,
	}
}

func (c *authorCollector) Process(content string) []error {
	ps, errs := c.Parse(content)

	if len(errs) > 0 {
		return errs
	}
	errors := make([]error, 0)
	// save in database the products

	for _, es := range ps {
		// que hacer con los autores y errores devueltos?
		a, err := c.authorWriter.Create(es)
	}

	if len(es) > 0 {
		errors = append(errors, es...)
	}
	return errors
}

func (c *authorCollector) Parse(content string) ([]*products.Author, []error) {

	results := make([]*products.Author, 0)
	payload, err := products.NewContentPayload(content)

	if err != nil {
		return results, []error{err}
	}

	ad, err := c.authorDefinitionReader.GetAuthorDefinition()
	if err != nil {
		return results, []error{err}
	}

	errors := make([]error, 0)
	// for each product definition configured in the database
	// it will try to parse the content and generates a map(result)
	// with the possible research products for each definition.

	// It will try to generate a yaml representation that contains the fields
	// and the regular expression to extracts the fields of each product
	definition, err := collector.NewDefinition(ad.Definition)

	if err != nil {
		errors = append(errors, err)
	}

	// it will try to create a map (result) that contains the parsed product
	// information from the content and the definition
	result, err := c.parser.Parse(definition, payload.Content)
	if err != nil {
		errors = append(errors, err)
	}

	// No me queda claro como crear los autores desde la data parseada
	// si devolver un autor o todos los autores
	authors := products.NewProductResults(result, pd.Name, payload.GrupLACCode, payload.GrupLACName)
	results = append(results, authors...)
	return authors, errors
}
