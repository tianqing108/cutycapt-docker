// Create by Yale 2018/12/4 15:20
package main

import (
	"bytes"
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
func Run(cmdPath string,args ...string) (string,string,error) {

	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd:=exec.Command(cmdPath,args...)
	cmd.Stdout = &out
	cmd.Stderr = &outErr

	err := cmd.Start()
	if err!=nil {
		return "","",err
	}
	err = cmd.Wait()
	if err!=nil {
		return "","",err
	}
	return out.String(),outErr.String(),nil
}
func GetCutyParms(params []string,r *http.Request, key string, defaultVal string) []string {
	p:=fmt.Sprintf("--%s=%s",key,GetQuery(r,key,defaultVal))
	params = append(params,p)
	return params
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

	params = append(params,"--server-args=\"-screen 0, 1920x1080x24\"")

	params = append(params,"CutyCapt")

	params = GetCutyParms(params,r,"url","")
	params = GetCutyParms(params,r,"min-width","720")
	params = GetCutyParms(params,r,"min-height","200")
	params = GetCutyParms(params,r,"javascript","on")
	params = GetCutyParms(params,r,"delay","3000")
	params = GetCutyParms(params,r,"max-wait","20000")
	params = GetCutyParms(params,r,"out-format","png")

	of:=GetQuery(r,"out-format","png")

	os.Mkdir("tmp",0666)
	outName:=fmt.Sprintf("tmp/%d.%s",time.Now().UnixNano(),of)

	params = append(params,"--out="+outName)

	log.Println(strings.Join(params," "))
	_,_,err:=Run("xvfb-run",params...)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
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