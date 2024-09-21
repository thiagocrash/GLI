package main

import (
    "fmt"
    "os/exec"
    "strings"
    "runtime"
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

func get_cpu() string {
  cpu := ""
  index := 4
  for {
    line := run_cmd(fmt.Sprintf("cat /proc/cpuinfo | grep 'model name' | head -1 | awk '{print $%d}'", index))
    if len(strings.TrimSpace(line)) == 0 || index > 7 {
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

func print_data() {
  reset := "\033[0m"
  yellow := "\033[33m"
  blue := "\033[34m" 
  fmt.Println(`
         _nnnn_                              `+blue+"GLI"+"   "+run_cmd("date | awk '{print $5}'")+reset+`
         dGGGGMMb     ,"""""""""""""".       `+blue+"----------------"+reset+`
       @p~qp~~qMb    |`+yellow+` Linux Rules! `+reset+`|        `+blue+"Kernel: "+reset+run_cmd("uname -r")+`
       M|@||@) M|   _;..............'        `+blue+"Distro: "+reset+run_cmd("hostnamectl | grep 'Operating System' | sed 's/^.*: //'")+`
       @,----.JM| -'                         `+blue+"Go Version: "+reset+runtime.Version()+`
       JS^\__/  qKL                          `+blue+"PC Name: "+reset+run_cmd("hostname")+`
     dZP        qKRb                         `+blue+"Architecture: "+reset+run_cmd("hostnamectl | grep 'Architecture' | sed 's/^.*: //'")+`
    dZP          qKKb                        `+blue+"Used RAM: "+reset+run_cmd("free -h | grep Mem | awk '{print $3}'")+` 
   fZP            SMMb                       `+blue+"Total RAM: "+reset+run_cmd("free -h | grep Mem | awk '{print $2}'")+`
   HZM            MMMM                       `+blue+"CPU: "+reset+get_cpu()+`
   FqM            MMMM                       `+blue+"GPU: "+reset+get_gpu()+`
 __| ".        |\dS"qML
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
  print_data()
}
