package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func getSensors() ([]string, error) {
  data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
  if err != nil {
      return nil, err
  }
  
  fmt.Println("Data: ", string(data))
  
  sensors := strings.Split(string(data), "\n")
  //fmt.Println("sensors befoer", sensors...)
  if len(sensors) > 0 {
    sensors = sensors[:len(sensors)-1]
  }
  return sensors, nil
}

func main() {
    sensors, err := getSensors()
    if err != nil {
        fmt.Println("ERROR", err)
        return
    }
    fmt.Println("SENSORS", sensors)
}

