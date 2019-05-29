package main
import (
    "runtime"
    "os"
)
func main() {
    print("系统类型：");
    println(runtime.GOOS);

    print("系统架构：");
    println(runtime.GOARCH);
    
    print("CPU 核数：");
    println(runtime.GOMAXPROCS(0));

    print("电脑名称：");
    name, err := os.Hostname()
    if err == nil {
        println(name)
    } else {
        println(err)
    }
}
