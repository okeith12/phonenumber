package dbReader

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/lib/pq"
	data "github.com/okeith12/quizGame/phonenumbernorm/pkg/data"
)

//OpenDatabase opens the database passed into the name
func OpenDatabase(db_name string) (*sql.DB,error){
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = ENTERPASSWORDHERE
	)
	 var dbname  =  db_name

	psqlconnect := fmt.Sprintf("host = %s port = %d user= %s password = %d dbname = %s sslmode=disable",host,port,user,password,dbname)
	database,err := sql.Open("postgres",psqlconnect)
	if err != nil{
		fmt.Println(err)
	}

	err = database.Ping()
	if err !=nil{
		fmt.Println("Database not found",err)
		return database,err
	}
	return database,nil
}
//ReadAllDataFromTable is taking a database table name and opening it and if its not there then it is created and returning the all te data from the rows.....but for this assignment, it is only opening the Phone Number Database and one parameter
func ReadAllDataFromTable(database *sql.DB, tableName *string, dataExist bool) []string{
	
	protectedDBname := pq.QuoteIdentifier(*tableName)
	openTable :=fmt.Sprintf("CREATE TABLE  IF NOT EXISTS public.%s(Number text, ID int)",protectedDBname)
	_,err := database.Exec(openTable)
	if err != nil{
		fmt.Println("Error creating table",err)
	}

	if !dataExist{
		insertData(&protectedDBname,database)
	}
	var dataAll []string
	rows, err := database.Query(`SELECT "number" FROM "phonenumbers" `)
	if err != nil{
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next(){
		var num string
		rows.Scan(&num)
		dataAll= append(dataAll,num)
	}
	return dataAll
}
//insertData takes a file and read it into the Database that was declared in the parameter
func insertData(dbName *string, db *sql.DB){
	data,_:= os.ReadFile("phone_numbers")
		pn := strings.Split(string(data),"\n")
	
		for id,d := range pn{
			dbINSERTstmt := fmt.Sprintf("INSERT into %s(Number,ID) values($1,$2)",*dbName)
			_,e := db.Exec(dbINSERTstmt,d,id)
			if e != nil{
				fmt.Println("Error inserting data",e)
			}
		}
}
//UpdateData takes the formatted data and input it into the database
func UpdateData(datta []string,tableName *string, db *sql.DB){
	data.RemoveSpecialCharacters(datta)
	protectedDBname:=pq.QuoteIdentifier(*tableName)
	for index,d := range datta{
		UPDATEstmt :=fmt.Sprintf("UPDATE %s SET Number = $2 WHERE ID = $1",protectedDBname)
		_,err := db.Exec(UPDATEstmt,index,d)
		if err !=nil{
			fmt.Println(err)
		}
	}
	DeleteDuplicates(datta,tableName,db)
	
}
//DELETE Duplicates find the duplicates in the Data and delete the ROW with the corresponding ID
func DeleteDuplicates(data []string,tableName *string, db *sql.DB){
	tempMap := make(map[string]bool)
	tempSlice := []string{}
	for index,item := range data{
		if _, val := tempMap[item]; !val{
				tempMap[item] = true
				
				tempSlice = append(tempSlice, item)
		}else{			
			fmt.Println("Deleting Duplicate Number........")
			DELETEstmt :=fmt.Sprintf("DELETE FROM %s WHERE id = $1",*tableName)
			_,e := db.Exec(DELETEstmt,index)
			if e != nil{
				fmt.Println(e)
			}
		
		}
	}


}
