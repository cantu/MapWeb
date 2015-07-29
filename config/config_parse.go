package config_parse

import(
	"fmt"
	"os"
	"encoding/json"
)



func ParsePostgreConf( file_path *string){
	file, err := os.Open( file_path )
	if err != nil {
		fmt.Println( err)
		return
	}
	decoder := json.Decoder(file)
	config := Config{}
	err := decoder.Decode( &config)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println( config.com_config.ProjectName)
}