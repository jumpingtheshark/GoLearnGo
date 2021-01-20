package main

type T struct{
X string
Y string
Z string

}

func (t T) Foo(){
	
}

func main(){

	t:=T{"X", "Y", "Z"}
	print (t.X)
	
}
