# HttpStatusCodeAPI
Description: A simple API endpoint for response http status codes 

## How to use?
1. Clones repo and run 
   ```go
      go run main.go
   ```
   If you don't have GO, go [here](https://golang.org/) for instructions of installation<br/>
   Access the endpoint at ***localhost:8080/status/[code number]***<br/>
   Eg: localhost:8080/status/404 <br/>
   => This will return a document with **404** status<br/>
2. Opens executable file by double-clicking release file that is compatible with your system. 
3. Accesses API through already hosted link, please follow the format below:<br/>
   **https://arcane-basin-93470.herokuapp.com/status/[_code number_]**<br/>
   Eg: https://arcane-basin-93470.herokuapp.com/status/200 <br/>
   => This will return a document with **200** status<br/>

## Features
1. Add delay time for response, by simply include a delay query <br/>
Eg1: \[host]/status/503?delay=10  => response with status code 503 after 10 seconds <br/>
Eg2: \[host]/status/500?delay=10s => response with status code 500 after 10 seconds <br/>
Eg3: \[host]/status/501?delay=5000ms => response with status code 501 after 5000 miliseconds <br/>

Note: This API has all the status codes that you can get back from GET request, 2xx - 5xx 
