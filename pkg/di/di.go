package di

type Builder func(c *Container) (interface{}, error)

var c *Container

type Container struct {
	definition map[string]*Definition
}

type Definition struct {
	Name      string
	Build     Builder
	completed bool
	cache     interface{}
}

func (c *Container) AddDefinition(defs ...*Definition) {
	for _, def := range defs {
		if def.Name != "" {
			c.definition[def.Name] = def
		}
	}
}

func (c *Container) Get(name string) interface{} {
	if def, ok := c.definition[name]; ok {
		return def.build(c)
	}

	return nil
}

func (def *Definition) build(c *Container) interface{} {
	if def.completed {
		return def.cache
	}

	obj, err := def.Build(c)
	if err != nil {
		return obj
	}
	def.completed = true
	def.cache = obj
	return def.cache
}

func AddDefinition(defs ...*Definition) {
	c.AddDefinition(defs...)
}

func Get(name string) interface{} {
	return c.Get(name)
}

func init() {
	c = &Container{
		definition: make(map[string]*Definition),
	}
}
