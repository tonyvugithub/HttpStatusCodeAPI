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

Note: This API has all the status codes that you can get back from GET request, 2xx - 5xx 
