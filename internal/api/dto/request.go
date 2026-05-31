package dto

type UpdateUserRequest struct {
	UserHandle string `json:"user_handle" binding:"required,min=3,max=30"`
	UserName   string `json:"user_name" binding:"required,min=1,max=50"`
}

type CreateProjectRequest struct {
	Name        string  `json:"name" binding:"required,min=1,max=100"`
	Description *string `json:"description"`
}

type UpdateProjectRequest struct {
	Name        string  `json:"name" binding:"required,min=1,max=100"`
	Description *string `json:"description"`
}

type CreateDummyDataRequest struct {
	TableName      string   `json:"table_name" binding:"required"`
	ColumnName     []string `json:"column_name" binding:"required,min=1"`
	ColumnType     []string `json:"column_type" binding:"required,min=1"`
	ColumnValidate []string `json:"column_validate"`
}

type UpdateDummyDataRequest struct {
	TableName      string   `json:"table_name" binding:"required"`
	ColumnName     []string `json:"column_name" binding:"required,min=1"`
	ColumnType     []string `json:"column_type" binding:"required,min=1"`
	ColumnValidate []string `json:"column_validate"`
}
