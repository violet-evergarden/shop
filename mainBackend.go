package main

import (
	_ "github.com/app"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/models"

	//greeter "github.com/protoBuf"
	"github.com/routers"
)


func main() {
	//service := micro.NewService(
	//	micro.Name("greeter"),
	//	micro.Version("latest"),
	//	)
	//service.Init()
	////var gree Greeter
	////gree := Greeter{}
	//greeter.RegisterGreeterHandler(service.Server(),new(Greeter))
	//if err := service.Run(); err != nil {
	//	//log.Fatal(err)
	//	fmt.Println(err)
	//}

	//go RegistryByETCD()
	//RegistryByETCD()
	//listen()

	//carss := &cars{}
	//carss.Name = "123123"
	//fmt.Println(carss.Name)
	//carss.run()
	//
	////fmt.Println(carss.run())
	//
	//return

	//var str1 string
	//str1 = "123__3)_3_12"
	//
	////str := "i am str"
	//fmt.Println(str1[0:3])
	//
	//
	//str := "123_a23"
	//buf := bytes.Buffer{}
	//var lastChar rune = 0
	//for _,v := range str {
	//	if lastChar == '_' && v == '_' {
	//		buf.WriteRune('_')
	//		continue
	//	}
	//	if v == '_' {
	//
	//	}
	//	if lastChar == '_' {
	//		if v >= 97 && v <= 122 {
	//			buf.WriteRune(v-32)
	//		}
	//		continue
	//	}
	//	lastChar = v
	//
	//	buf.WriteRune(v)
	//}
	//fmt.Printf("i am string:%s\n",buf.String())
	//
	//conditions := &msql.Conditions{}
	//conditions.Push("add_time",">",0)
	//
	//user := &User{}

	//fields := reflect2(user)
	//aa := reflect34(user)
	//fmt.Println("i am aa")
	//fmt.Println(aa)
	//fmt.Println("i am aa")
	////fields[1] = "123123"
	////fields[2] = "123123"
	//*fields[0].(*int) = 123
	//
	//fmt.Printf("i am user.Id = %d",user.Id)
	//fmt.Println(fields[0])
	//m := &msql.Msql{}
	//m.Select([]string{"id"}).From("shop").Where(conditions).Query(user)
	//
	//fmt.Printf("i am user.Id = %d",user.Id)

	//return

	//dsn := "root:Forever634312.@tcp(192.168.0.200:3306)/server"
	//db,err := sql.Open("mysql",dsn)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//db.SetMaxOpenConns(50)
	//db.SetMaxIdleConns(5)
	//defer db.Close()
	//
	//
	//sli := []int{1,2,3,4,5}
	//sli = sli[2:5]
	//
	//fmt.Println(sli)
	//
	////res,err := db.Query("select `id` from `admin`")
	//res,err := db.Query("insert into `admin` (`username`,`password`) values ('xiaozhong01','')")
	//if err != nil{
	//	var abc = err.Error()[6:10]
	//	if abc == "1062" {
	//		fmt.Println("i am 1062")
	//	}
	//	//fmt.Println(err.Error()[6:10])
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	////fmt.Println(res.Err().Error())
	//fmt.Println(res.Columns())

	//for res.Next() {
	//	var id int
	//	fmt.Println(res.Scan(&id))
	//	fmt.Println(id)
	//}
	//fmt.Println(res.Columns())

	r := routers.InitRouterForBackend()
	r.Run("0.0.0.0:8082")
}
