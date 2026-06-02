package main

import (
	"fmt"
	"time"
)

func checkStatus(serviceName string){
	fmt.Printf("Starting the monitoring process for %s.....\n", serviceName)
	time.Sleep(10000 * time.Millisecond)
	fmt.Printf("Service %s is Healthy\n", serviceName)
}

func main(){
	go checkStatus("AuthAPI")
	go checkStatus("Payment Gateway")
	go checkStatus("Database Cluster")

	time.Sleep(15 * time.Second)
	fmt.Println("All initial checks executed")
}