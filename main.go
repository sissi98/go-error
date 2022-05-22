package main
import (
	"fmt"
	"go-error/app/dao"
)
func main() {
	defer dao.DB.Close()
	user ,err :=dao.QueryUserById("99999")
	if err != nil{
	  fmt.Printf("query user err : %+v",err)
	  return
	}
	fmt.Println("query user : ",user)
}
