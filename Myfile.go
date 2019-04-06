package main
import ("fmt")
import("github.com/EZDIR")
//import ("github.com/ehteshamz/greetings")

func ConcatSlice(sliceToConcat []byte) string{
  temp:=""
  for i := 0; i < len(sliceToConcat); i++ {
        temp+=string(sliceToConcat[i])
        if(i!=len(sliceToConcat)-1){
          temp+="-"
        }
    }
  return temp
}

func Encrypt(sliceToEncrypt[] byte,ceaserCount int) string{
   temp:=""
   for i := 0; i < len(sliceToEncrypt); i++ {
         temp+=(string(int(sliceToEncrypt[i])+2))
     }
  return temp
}

func EZGreetings(name string) string{
  return greetings.PrintGreetings(name)
}

func main(){
  value :=""
  array:=[] byte("shayan")
  value=ConcatSlice(array)
  fmt.Printf("%s\n",value)

  value=Encrypt(array,2)
  fmt.Printf("%s\n",value)

  value=EZGreetings("Shayan")
  fmt.Printf("%s\n",value)
}
