package common

const (
	AddCommand         = "add"
	ViewCommand        = "view"
	EditCommand        = "edit"
	DeleteCommand      = "delete"
	ExportCommand      = "export"
	SetCategoryCommand = "set-category"
)

type FlagName string

const (
	MessFlag     FlagName = "message"
	CategoryFlag FlagName = "category"
)

type ShortFlag string

const (
	MessShortFlag     ShortFlag = "m"
	CategoryShortFlag ShortFlag = "c"
)
