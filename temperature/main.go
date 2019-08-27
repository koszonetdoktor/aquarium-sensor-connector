package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "errors"
  "strconv"
  "time"
)

func main() {
    sensors, err := getSensors()
    if err != nil {
        fmt.Println("ERROR", err)
        return
    }
    fmt.Println("SENSORS", sensors)
    for {
      temperature, err := readTemperature(sensors[0])
      if err != nil {
          fmt.Println(err)
      }
      fmt.Println("Temp: ", temperature)
      time.Sleep(1 * time.Second)
    }
}

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

func readTemperature(sensor string) (float64, error) {
  data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + sensor + "/w1_slave")
  if err != nil {
      return 0.0, errors.New("Sensor read error")
  }
  
  raw := string(data)
  
  i := strings.LastIndex(raw, "t=")
  
  if i == -1 {
      return 0.0, errors.New("Sensor read error")
  }
  
  c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
  if err != nil {
      return 0.0, errors.New("Sensor read error")
  }
  
  return c/ 1000.0, nil
}









