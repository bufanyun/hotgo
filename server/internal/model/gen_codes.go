package model

// GenCodesColumn 生成表字段属性
type GenCodesColumn struct {
	// 表属性
	Id           int64       `json:"id" dc:"序号"`
	Name         string      `json:"name" dc:"字段列名"`
	Dc           string      `json:"dc" dc:"字段描述"`
	DataType     string      `json:"dataType" dc:"字段类型"`
	SqlType      string      `json:"sqlType" dc:"物理类型"`
	Length       int64       `json:"length" dc:"字段长度"`
	IsAllowNull  string      `json:"isAllowNull" dc:"是否允许为空"`
	DefaultValue interface{} `json:"defaultValue" dc:"默认值"`
	Index        string      `json:"index" dc:"索引"`
	Extra        string      `json:"extra" dc:"额外选项"`
	// 自定义生成属性
	//Alias      string `json:"alias" dc:"字段别名"`
	GoName     string `json:"goName" dc:"Go属性"`
	GoType     string `json:"goType" dc:"Go类型"`
	TsName     string `json:"tsName" dc:"Ts属性"`
	TsType     string `json:"tsType" dc:"Ts类型"`
	IsList     bool   `json:"isList" dc:"列表"`
	IsExport   bool   `json:"isExport" dc:"导出"`
	IsSort     bool   `json:"isSort" dc:"排序"`
	IsQuery    bool   `json:"isQuery" dc:"查询"`
	QueryWhere string `json:"queryWhere" dc:"查询条件"`
	IsEdit     bool   `json:"isEdit" dc:"编辑"`
	Required   bool   `json:"required" dc:"必填"`
	Unique     bool   `json:"unique" dc:"唯一性"`
	FormMode   string `json:"formMode" dc:"表单组件"`
	FormRole   string `json:"formRole" dc:"表单验证"`
	DictType   int64  `json:"dictType" dc:"字典类型ID"`
}
