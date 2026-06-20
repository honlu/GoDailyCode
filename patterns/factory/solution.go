package factory

type RuleConfigParser interface {
	Parse(data []byte)
}

type JSONRuleConfigParser struct{}

func (JSONRuleConfigParser) Parse(data []byte) {}

type YAMLRuleConfigParser struct{}

func (YAMLRuleConfigParser) Parse(data []byte) {}

type RuleConfigParserFactory interface {
	CreateParser() RuleConfigParser
}

type YAMLRuleConfigParserFactory struct{}

func (YAMLRuleConfigParserFactory) CreateParser() RuleConfigParser {
	return YAMLRuleConfigParser{}
}

type JSONRuleConfigParserFactory struct{}

func (JSONRuleConfigParserFactory) CreateParser() RuleConfigParser {
	return JSONRuleConfigParser{}
}

func NewRuleConfigParserFactory(kind string) RuleConfigParserFactory {
	switch kind {
	case "json":
		return JSONRuleConfigParserFactory{}
	case "yaml":
		return YAMLRuleConfigParserFactory{}
	default:
		return nil
	}
}
