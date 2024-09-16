package main

// "io"
// "bufio"

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)
var (
  name string
  version string
)

func get_attrib(file string , keyword string) string {
  scanner := bufio.NewScanner(strings.NewReader(file))
  ret := ""
  for scanner.Scan() {
    line := scanner.Text()
    if strings.HasPrefix(line, keyword) {
      ret = line
      break
    }
  }
  if scanner.Err() != nil {
    panic(scanner.Err())
  }
  if len(ret) > 0 {
    ret = get_between_quotes(ret)
  } 
  return ret 
}

func get_between_quotes(line string) string {
  start := strings.Index(line, "\"")
  if start == -1 {
    return "" 
  }
  end := strings.Index(line[start+1:], "\"")
  if end == -1 {
    return "" 
  }
  return line[start+1 : start+1+end] 
}
func print_data() {

  output := `
         _nnnn_                              `+"Golang Status Program"+`
         dGGGGMMb     ,"""""""""""""".       `+"Distro: "+name+`
       @p~qp~~qMb    | Linux Rules! |        `+"Version: "+version+`
       M|@||@) M|   _;..............'
       @,----.JM| -'
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    \.       | \' \Zq
_)      \.___.,|     .'
\____   )MMMMMM|   .'
     \-'       \--'
  `
  fmt.Println(output)
}
func main() {
  distro_info, err := os.ReadFile("/etc/os-release")
  if err != nil {
    panic(err)
  }
  // memory_info, err := os.ReadFile("/proc/meminfo")
  // if err != nil {
  //   panic(err)
  // }
  name = get_attrib(string(distro_info), "NAME")
  version = get_attrib(string(distro_info), "VERSION")
  print_data()
}
