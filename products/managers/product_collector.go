package managers

import (
	"github.com/udistritali3plus/collector"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductCollector)
	core.CheckInjection(err, "newProductCollector")
}

type productCollector struct {
	parser                  collector.Parser
	productDefinitionReader ports.ProductDefinitionReader
	productGenericWriter    ports.GenericProductWriter
}

func newProductCollector(pdr ports.ProductDefinitionReader, pw ports.GenericProductWriter) ports.ProductCollector {
	parser := collector.NewParser()
	return &productCollector{
		parser:                  parser,
		productDefinitionReader: pdr,
		productGenericWriter:    pw,
	}
}

func (c *productCollector) Process(content string) []error {
	ps, errs := c.Parse(content)

	if len(errs) > 0 {
		return errs
	}

	errors := make([]error, 0)

	// save each
	for _, p := range ps {
		_, es := c.productGenericWriter.WriteGenerics(p)

		if len(es) > 0 {
			errors = append(errors, es...)
		}
	}

	return errors
}

func (c *productCollector) Parse(content string) ([]*products.ParsedProducts, []error) {

	results := make([]*products.ParsedProducts, 0)
	payload, err := products.NewContentPayload(content)

	if err != nil {
		return results, []error{err}
	}

	definitions, err := c.productDefinitionReader.GetAll()

	if err != nil {
		return results, []error{err}
	}

	errors := make([]error, 0)
	// for each product definition configured in the database
	// it will try to parse the content and generates a map(result)
	// with the possible research products for each definition.
	for _, pd := range definitions {
		// It will try to generate a yaml representation that contains the fields
		// and the regular expression to extracts the fields of each product
		definition, err := collector.NewDefinition(pd.Definition)

		if err != nil {
			errors = append(errors, err)
			continue
		}

		// it will try to create a map (result) that contains the parsed product
		// information from the content and the definition
		result, err := c.parser.Parse(definition, payload.Content)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		pp := products.NewParsedProducts(result, pd.Name, payload.GrupLACCode, payload.GrupLACName)
		results = append(results, pp)
	}
	return results, errors
}
