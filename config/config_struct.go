package config_parse



type Config struct{
	com_config struct{
		ProjectName []string
	}

	dataBase struct{
		postgre struct{
			host		[]string
			user		[]string
			password	[]string
			dbname		[]string
			sslmode		[]string
		}
		mysql struct{
			host		[]string	//hello
			user		[]string
			password	[]string
			dbname		[]string
		}
	}

}





/*
{

"com_config":{
"ProjectName":"GrabTaxi_RoadBuilder",

},

"database": {
"postgre": {
"host": "192.168.56.101",
"user": "postgres",
"password": "123456",
"dbname": "grabtaxi",
"sslmode": "disable"
},
"mysql": {
"host": "192.168.56.101",
"user": "root",
"password": "root",
"dbname": "grabtaxi"
}
}
}

*/