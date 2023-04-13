package generator

//GenConstants 代码生成常量
var GenConstants = genConstants{
	UTF8:      "utf-8", //编码
	TplCrud:   "crud",  //单表 (增删改查)
	TplTree:   "tree",  //树表 (增删改查)
	QueryLike: "LIKE",  //模糊查询
	QueryEq:   "=",     //相等查询
	Require:   1,       //需要的
}

//GoConstants Go相关常量
var GoConstants = goConstants{
	TypeString: "string",    //字符串类型
	TypeFloat:  "float64",   //浮点型
	TypeInt:    "int",       //整型
	TypeDate:   "time.Time", //时间类型
}

//SqlConstants 数据库相关常量
var SqlConstants = sqlConstants{
	//数据库字符串类型
	ColumnTypeStr: []string{"char", "varchar", "nvarchar", "varchar2"},
	//数据库文本类型
	ColumnTypeText: []string{"tinytext", "text", "mediumtext", "longtext"},
	//数据库时间类型
	ColumnTypeTime: []string{"datetime", "time", "date", "timestamp"},
	//数据库数字类型
	ColumnTypeNumber: []string{"tinyint", "smallint", "mediumint", "int", "integer", "bit", "bigint",
		"float", "double", "decimal"},
	//时间日期字段名
	ColumnTimeName: []string{"create_time", "update_time", "delete_time", "start_time", "end_time"},
	//页面不需要插入字段
	ColumnNameNotAdd: []string{"id", "is_delete", "create_time", "update_time", "delete_time"},
	//页面不需要编辑字段
	ColumnNameNotEdit: []string{"is_delete", "create_time", "update_time", "delete_time"},
	//页面不需要列表字段
	ColumnNameNotList: []string{"id", "intro", "content", "is_delete", "delete_time"},
	//页面不需要查询字段
	ColumnNameNotQuery: []string{"is_delete", "create_time", "update_time", "delete_time"},
}

//HtmlConstants HTML相关常量
var HtmlConstants = htmlConstants{
	HtmlInput:       "input",       //文本框
	HtmlTextarea:    "textarea",    //文本域
	HtmlSelect:      "select",      //下拉框
	HtmlRadio:       "radio",       //单选框
	HtmlDatetime:    "datetime",    //日期控件
	HtmlImageUpload: "imageUpload", //图片上传控件
	HtmlFileUpload:  "fileUpload",  //文件上传控件
	HtmlEditor:      "editor",      //富文本控件
}

type genConstants struct {
	UTF8      string
	TplCrud   string
	TplTree   string
	QueryLike string
	QueryEq   string
	Require   uint8
}

type goConstants struct {
	TypeString string
	TypeFloat  string
	TypeInt    string
	TypeDate   string
}

type sqlConstants struct {
	ColumnTypeStr      []string
	ColumnTypeText     []string
	ColumnTypeTime     []string
	ColumnTypeNumber   []string
	ColumnTimeName     []string
	ColumnNameNotAdd   []string
	ColumnNameNotEdit  []string
	ColumnNameNotList  []string
	ColumnNameNotQuery []string
}

type htmlConstants struct {
	HtmlInput       string
	HtmlTextarea    string
	HtmlSelect      string
	HtmlRadio       string
	HtmlDatetime    string
	HtmlImageUpload string
	HtmlFileUpload  string
	HtmlEditor      string
}
