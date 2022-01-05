package main

import (
	"fmt"

	dbReader "github.com/okeith12/quizGame/phonenumbernorm/pkg/database"
)

func main() {
	dataBase := "first_db"
	tableName := "Phone Numbers"
	database,err :=dbReader.OpenDatabase(dataBase)
	if err !=nil{
		fmt.Println(err)
	}
	defer database.Close()

	tableData := dbReader.ReadAllDataFromTable(database,&tableName,true)
	dbReader.UpdateData(tableData,&tableName,database)

}