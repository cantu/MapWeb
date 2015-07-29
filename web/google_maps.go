package web
import(
	"fmt"
	"net/http"
	"strings"
	"log"
	"text/template"
)

func sayhelloName( w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v," "))
	}
	fmt.Fprint( w, "hello from go server!")
}

func login( w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method == "GET"{
		t,_:= template.ParseFiles("login.gtpl.html")
		t.Execute(w, nil)
	}else{
		//post method
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func StartWebServer(){
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err:= http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}