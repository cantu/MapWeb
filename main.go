package main

import(
	"fmt"
	"RoadBuilder/grab_db"
	//"RoadBuilder/web"
	//"RoadBuilder/config_parse"

)

func main(){
	fmt.Println("start main.go")


	//web.StartWebServer()
	//config_parse.ParsePostgreConf("config.json")
	//grab_db.InitLocationHistory()
	grab_db.GetBookingID()
}