// Create by Yale 2018/12/4 15:20
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func FileExist(filePath string) bool {
	_,err:=os.Stat(filePath)
	if err==nil {
		return  true
	}
	return false
}
func GetCutyParms(params []string,r *http.Request, key string, defaultVal string) []string {
	v:=GetQuery(r,key,defaultVal)
	if len(v) == 0 {
		return params
	}
	return append(params,fmt.Sprintf("--%s=%s",key,v))
}
func GetQuery(r *http.Request, key string, defaultVal string) string {
	values, ok := r.URL.Query()[key]
	if ok && len(values) > 0 && len(values[0]) > 0 {
		return values[0]
	}
	return defaultVal
}
func HandlerCutyCapt(w http.ResponseWriter,r *http.Request)  {
	url:=GetQuery(r,"url","")
	if len(url) == 0 || !strings.HasPrefix(url,"http"){
        w.WriteHeader(http.StatusBadRequest)
		return
	}
	params:=make([]string,0)

	params = append(params,"xvfb-run")
	params = append(params,"--server-args=\"-screen 0, 1920x1080x24\"")

	params = append(params,"CutyCapt")

	params = append(params,fmt.Sprintf("--url=\"%s\"",url))

	of:=GetQuery(r,"out-format","png")
	os.Mkdir("/tmp",0666)
	outName:=fmt.Sprintf("/tmp/%d.%s",time.Now().UnixNano(),of)
	params = append(params,"--out="+outName)

	params = GetCutyParms(params,r,"out-format","png")
	params = GetCutyParms(params,r,"min-width","800")
	params = GetCutyParms(params,r,"min-height","600")
	params = GetCutyParms(params,r,"max-wait","30000")
	params = GetCutyParms(params,r,"delay","0")
	params = GetCutyParms(params,r,"user-style-path","")
	params = GetCutyParms(params,r,"header","")
	params = GetCutyParms(params,r,"method","get")
	params = GetCutyParms(params,r,"body-string","")
	params = GetCutyParms(params,r,"body-base64","")
	params = GetCutyParms(params,r,"app-name","")
	params = GetCutyParms(params,r,"app-version","")
	params = GetCutyParms(params,r,"user-agent","")
	params = GetCutyParms(params,r,"javascript","on")
	params = GetCutyParms(params,r,"java","off")
	params = GetCutyParms(params,r,"plugins","off")
	params = GetCutyParms(params,r,"private-browsing","off")
	params = GetCutyParms(params,r,"auto-load-images","on")
	params = GetCutyParms(params,r,"js-can-open-windows","off")
	params = GetCutyParms(params,r,"js-can-access-clipboard","off")
	params = GetCutyParms(params,r,"print-backgrounds","off")
	params = GetCutyParms(params,r,"zoom-factor","")
	params = GetCutyParms(params,r,"zoom-text-only","off")
	params = GetCutyParms(params,r,"http-proxy","")


	cmdStr:=strings.Join(params," ")

	log.Println(cmdStr)
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w,err.Error())
		log.Println(err.Error())
		return
	}
	if FileExist(outName) {
		contents, err := ioutil.ReadFile(outName)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w,err.Error())
			log.Println(err.Error())
			return
		}
		w.Header().Set("content-type", "image/"+of)
		w.Write(contents)
	}else{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w,"convert err")
		log.Println("convert err")
	}

}
func HandlerLog(handler http.Handler)http.Handler  {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method,r.URL,r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}
func main()  {
	port:=flag.String("port","9600","http port")
	flag.Parse()


	http.HandleFunc("/cutycapt",HandlerCutyCapt)

	server := &http.Server{
		Addr:    ":"+*port,
		Handler: HandlerLog(http.DefaultServeMux),
	}

	log.Println("listen at : "+*port)
	err := server.ListenAndServe()
	if err!=nil{
		log.Println(err.Error())
	}

}