

## start container

```bash

#:docker pull yale8848/cutycapt-docker:latest
#:docker run name url2image -p 9600:9600 yale8848/cutycapt-docker:latest

```

## app use

```bash

http://127.0.0.1/cutycapt?url=http://www.baidu.com&delay=3000

```
## params
  same as: http://cutycapt.sourceforge.net/
          
  url=<url>                    The URL to capture (http:...|file:...|...)     
  out-format=<f>               png|pdf|ps|svg|jpeg,default:png 
  min-width=<int>              Minimal width for the image (default: 800)   
  min-height=<int>             Minimal height for the image (default: 600)  
  max-wait=<ms>                Don't wait more than (default: 30000, inf: 0)
  delay=<ms>                   After successful load, wait (default: 0)     
  user-style-path=<path>       Location of user style sheet file, if any    
  user-style-string=<css>      User style rules specified as text           
  header=<name>:<value>        request header; repeatable; some can't be set
  method=<get|post|put>        Specifies the request method (default: get)  
  body-string=<string>         Unencoded request body (default: none)       
  body-base64=<base64>         Base64-encoded request body (default: none)  
  app-name=<name>              appName used in User-Agent; default is none  
  app-version=<version>        appVers used in User-Agent; default is none  
  user-agent=<string>          Override the User-Agent header Qt would set  
  javascript=<on|off>          JavaScript execution (default: on)           
  java=<on|off>                Java execution (default: unknown)            
  plugins=<on|off>             Plugin execution (default: unknown)          
  private-browsing=<on|off>    Private browsing (default: unknown)          
  auto-load-images=<on|off>    Automatic image loading (default: on)        
  js-can-open-windows=<on|off> Script can open windows? (default: unknown)  
  js-can-access-clipboard=<on|off> Script clipboard privs (default: unknown)
  print-backgrounds=<on|off>   Backgrounds in PDF/PS output (default: off)  
  zoom-factor=<float>          Page zoom factor (default: no zooming)       
  zoom-text-only=<on|off>      Whether to zoom only the text (default: off) 
  http-proxy=<url>             Address for HTTP proxy server (default: none)