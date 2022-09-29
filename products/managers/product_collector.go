package managers

import (
	"github.com/JairoAndresR/collector"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
	"fmt"
)

func init() {
	err := core.Injector.Provide(newProductCollector)
	core.CheckInjection(err, "newProductCollector")
}

type productCollector struct {
	parser                  collector.Parser
	productDefinitionReader ports.ProductDefinitionReader
	productGenericWriter    ports.GenericProductWriter
	authorsReader           ports.AuthorsReader
	linker                  ports.ProductAuthorLinker
}

func newProductCollector(
	pdr ports.ProductDefinitionReader,
	pw ports.GenericProductWriter,
	ar ports.AuthorsReader,
	linker ports.ProductAuthorLinker,
) ports.ProductCollector {

	parser := collector.NewParser()
	return &productCollector{
		parser:                  parser,
		productDefinitionReader: pdr,
		productGenericWriter:    pw,
		authorsReader:           ar,
		linker:                  linker,
	}
}

func (c *productCollector) Process(content string) []error {
	ps, errs := c.Parse(content)

	if len(errs) > 0 {
		return errs
	}
	errors := make([]error, 0)
	// save in database the products
	_, es := c.productGenericWriter.WriteGenerics(ps)
	if len(es) > 0 {
		errors = append(errors, es...)
	}
	return errors
}

func (c *productCollector) Parse(content string) ([]*products.ProductResult, []error) {

	results := make([]*products.ProductResult, 0)
	payload, err := products.NewContentPayload(content)

	if err != nil {
		return results, []error{err}
	}

	definitions, err := c.productDefinitionReader.GetAll()
	if err != nil {
		return results, []error{err}
	}

	authors, err := c.authorsReader.GetAll()
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
		fmt.Println("parser result")
		for _, mapf := range result {
			for key, value := range mapf {
				fmt.Println(key, ": ", value)
			}
			
		}
		if err != nil {
			errors = append(errors, err)
			continue
		}

		products := products.NewProductResults(result, pd.Name, payload.GrupLACCode, payload.GrupLACName)
		products = c.linker.Link(authors, products)
		results = append(results, products...)
	}
	return results, errors
}
