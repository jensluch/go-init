package main

const API_PATTERN = "/api/"
const API_PORT = ":8080"

func main() {
    server := Server {
        API_PATTERN,
        "",
        API_PORT,
    }
    //server.SetPattern(API_PATTERN)
    //server.SetAddr("")
    //server.SetPort(API_PORT)
    server.Start()
}
