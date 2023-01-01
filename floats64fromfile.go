//Пакет datafile предназначен для чтения данных из файлов.
package datafile

import (
         "bufio"
         "os"
         "strconv"
)

                                                    //старый код func GetFloats(fileName string) ([3]float64, error) {
func GetFloats(fileName string) ([]float64, error) {
                                                    //старый код var numbers [3]float64
       var numbers []float64
       var temp float64

       file, err := os.Open(fileName)
       if err != nil {
              return nil, err//старое было return numbers, err
       }
       //i := 0
       scanner :=bufio.NewScanner(file)
       for scanner.Scan() {
              temp, err = strconv.ParseFloat(scanner.Text(), 64) // numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
              if err != nil {
                     return nil, err//старое было return numbers, err
              }
              numbers = append(numbers, temp)
              //i++
       }
       err = file.Close()
       if err != nil {
              return nil, err//старое было return numbers, err
       }
       if scanner.Err() !=nil {
              return nil, scanner.Err()//старое было return numbers, scanner.Err()
       }
       return numbers, nil
}
 