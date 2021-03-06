package main

import (
    "fmt"
    "os"
    //"io/ioutil"
    "net/http"
    "log"
    "strings"
    "time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
    //fmt.Println(req)
    fmt.Println(req.URL)
    fmt.Println("hello:"+req.RequestURI)

}


func Read(p []byte) (n int, err error) {
  fmt.Println(p)
  return len(p),nil
}

func mylog(filename string, message string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(message); err != nil {
		log.Println(err)
	}
}

func HelloServerRoot(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    t:=time.Now()
    var mytime string
    //fmt.Print(t.Format(time.UnixDate))
    mytime=t.Format("20060102030405")
    fmt.Print(mytime)
    fmt.Print(",")
    fmt.Print(req.RemoteAddr)
    fmt.Print(",")
    fmt.Println(req.RequestURI)
    getparameters:=strings.Split(req.RequestURI,"/")
    //fmt.Println(getparameters)
    if(len(getparameters)==3){
    	project_name:=getparameters[1]
    	fw_version:=getparameters[2]
    	url:="https://iotfw.ninux.org/"
	switch {
	//case project_name=="https_server" && fw_version=="0.4":
    	//	fmt.Fprintf(w, "1,%sfw/iperf0.1.bin,",url)
	case project_name=="https_server" && fw_version=="CasettaCaplsule2.4_0.5":
    	 	fmt.Fprintf(w, "1,%sfw/https_server_CasettaCaplsule2.4_0.6.bin,",url)
	case project_name=="esp-idf_wifi_mqtt_bme280" && fw_version=="0.1":
    	 	fmt.Fprintf(w, "1,%sfw/https_server0.4.bin,",url)
	case project_name=="https_server" && fw_version=="0.4":
    	 	fmt.Fprintf(w, "1,%sfw/%s0.5.bin,",url,project_name)
	default:
    		fmt.Fprintf(w, "0,%sfw/%s.bin,",url,project_name)
    	}
    }
    logstr:=fmt.Sprintf("%s,%s,%s\n",mytime,req.RemoteAddr,req.RequestURI)
    fmt.Println(logstr)
    mylog("logs/requests", logstr)

}

func main() {
    http.HandleFunc("/", HelloServerRoot)
    http.HandleFunc("/hello/", HelloServer)
    http.Handle("/fw/", http.FileServer(http.Dir("./")))
    //err := http.ListenAndServeTLS(":443", "ca_cert.pem", "ca_key.pem", nil)
    //err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil) // WORKS
    err := http.ListenAndServeTLS(":443", "ca_cert.pem", "ca_key_encrypted.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
