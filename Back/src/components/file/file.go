package file

import (
	"database/sql"
	"errors"
	"strings"
)

type File struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	DirectoryId int64  `json:"directoryId"`
	Path        string `json:"path"`
	Size        int64  `json:"size"`
	Cover       string `json:"cover"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

func InsertFile(file File, conn *sql.DB) error {
	t := `insert into file (name,type,directory_id,path,size,cover) values (?,?,?,?,?,?)`

	result, err := conn.Exec(t, file.Name, file.Type, file.DirectoryId, file.Path, file.Size, file.Cover)

	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("insert error")
	} else {
		return nil
	}
}

/*
获取文件后缀，返回文件名后缀 .png、.mp4 等，没有后缀将会返回""
传入: aaa.mp4，返回: mp4
*/
func GetFileSuffix(fileName string) string {
	// 如果有.的文件，才会有后缀名
	if strings.ContainsAny(fileName, ".") {
		paths := strings.Split(fileName, ".")
		return paths[len(paths)-1]
	}
	return ""
}