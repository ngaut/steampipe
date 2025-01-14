package parse

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/turbot/steampipe/steampipeconfig/modconfig"
)

var ConfigBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "connection",
			LabelNames: []string{"name"},
		},
		{
			Type:       "options",
			LabelNames: []string{"type"},
		},
	},
}

var ConnectionBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name:     "plugin",
			Required: true,
		},
		{
			Name: "type",
		},
		{
			Name: "connections",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "options",
			LabelNames: []string{"type"},
		},
	},
}

// ModBlockSchema :: top level schema for all mod resources
var ModBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       string(modconfig.BlockTypeMod),
			LabelNames: []string{"name"},
		},
		{
			Type:       modconfig.BlockTypeVariable,
			LabelNames: []string{"name"},
		},
		{
			Type:       modconfig.BlockTypeQuery,
			LabelNames: []string{"name"},
		},
		{
			Type:       modconfig.BlockTypeControl,
			LabelNames: []string{"name"},
		},
		{
			Type:       modconfig.BlockTypeBenchmark,
			LabelNames: []string{"name"},
		},
		{
			Type:       modconfig.BlockTypeReport,
			LabelNames: []string{"name"},
		},
		{
			Type:       modconfig.BlockTypePanel,
			LabelNames: []string{"name"},
		},
		{
			Type: modconfig.BlockTypeLocals,
		},
	},
}

var PanelBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "title"},
		{Name: "text"},
		{Name: "type"},
		{Name: "width"},
		{Name: "height"},
		{Name: "source"},
		{Name: "sql"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "panel",
			LabelNames: []string{"name"},
		},
		{
			Type:       "report",
			LabelNames: []string{"type"},
		},
	},
}

var ReportBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "title"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "panel",
			LabelNames: []string{"name"},
		},
		{
			Type:       "report",
			LabelNames: []string{"type"},
		},
	},
}

var ControlBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "description"},
		{Name: "documentation"},
		{Name: "search_path"},
		{Name: "search_path_prefix"},
		{Name: "severity"},
		{Name: "sql"},
		{Name: "query"},
		{Name: "tags"},
		{Name: "title"},
		{Name: "args"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "param",
			LabelNames: []string{"name"},
		},
	},
}

var QueryBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "description"},
		{Name: "documentation"},
		{Name: "search_path"},
		{Name: "search_path_prefix"},
		{Name: "sql"},
		{Name: "tags"},
		{Name: "title"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "param",
			LabelNames: []string{"name"},
		},
	},
}
var ParamDefBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "description"},
		{Name: "default"},
	},
}
