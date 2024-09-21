package main

import (
    "fmt"
    "os"
    "os/exec"
    "bufio"
    "strings"
    "runtime"
)
var (
  name string
  version string
)
func run_cmd(command string) string {
  cmd := exec.Command("bash", "-c", command)
  output , err := cmd.Output()

  if err != nil {
    fmt.Printf("Command executed with a error! (%s)\n",command)
    return ""
  }
  
  return strings.TrimSpace(string(output)) 
}
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

func get_cpu() string {
  cpu := ""
  index := 4
  for {
    line := run_cmd(fmt.Sprintf("cat /proc/cpuinfo | grep 'model name' | head -1 | awk '{print $%d}'", index))
    if len(strings.TrimSpace(line)) == 0 {
      break
    } else {
      cpu = fmt.Sprintf("%s %s", cpu, line)
      index += 1
    }
  }
  return strings.TrimSpace(cpu)
}

func get_gpu() string {
  gpu := ""
  index := 5
  for {
    line := run_cmd(fmt.Sprintf("glxinfo | grep 'OpenGL renderer' | awk '{print $%d}'", index))
    if len(strings.TrimSpace(line)) == 0 {
      break
    } else {
      gpu = fmt.Sprintf("%s %s", gpu , line)
      index += 1
    }
  }
  return strings.TrimSpace(gpu)
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
  reset := "\033[0m"
  yellow := "\033[33m"
  blue := "\033[34m" 
  fmt.Println(`
         _nnnn_                              `+blue+"GLI"+reset+`
         dGGGGMMb     ,"""""""""""""".       `+blue+"----------------"+reset+`
       @p~qp~~qMb    |`+yellow+` Linux Rules! `+reset+`|        `+blue+"Kernel: "+reset+run_cmd("uname -r")+`
       M|@||@) M|   _;..............'        `+blue+"Distro: "+reset+name+`
       @,----.JM| -'                         `+blue+"Version: "+reset+version+`
       JS^\__/  qKL                          `+blue+"Go Version: "+reset+runtime.Version()+`
     dZP        qKRb                         `+blue+"PC Name: "+reset+run_cmd("hostnamectl | grep 'Static' | sed 's/^.*: //'")+`
    dZP          qKKb                        `+blue+"Architecture: "+reset+run_cmd("hostnamectl | grep 'Architecture' | sed 's/^.*: //'")+` 
   fZP            SMMb                       `+blue+"Used RAM: "+reset+run_cmd("free -h | grep Mem | awk '{print $3}'")+`
   HZM            MMMM                       `+blue+"Total RAM: "+reset+run_cmd("free -h | grep Mem | awk '{print $2}'")+`
   FqM            MMMM                       `+blue+"CPU: "+reset+get_cpu()+`
 __| ".        |\dS"qML                      `+blue+"GPU: "+reset+get_gpu()+`
 |    \.       | \' \Zq
_)      \.___.,|     .'
\____   )MMMMMM|   .'
     \-'       \--'
  `)
//
}
func main() {
  if runtime.GOOS != "linux" {
    fmt.Println("GLI is only supported on Linux")
    return
  }
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
