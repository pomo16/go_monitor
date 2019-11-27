package main

func Init() {

}

func main() {
	Init()
	if err := InstanceRoutine().Run(":8080"); err != nil {
		panic(err)
	}
}
