package Utility

import(
	// "strconv"
)

// func GetNextLexString(str string) string{
// 	var temp string
// 	n:=len(str)
// 	c:=str[n-1]
// 	x:=Atoi(c)
// 	x=x+1
// 	rem:=x%10
// 	quo:=x/10
// 	// str[n-1]=strconv.Iota(rem)
// 	temp=temp+Iota(rem)
// 	// temp=append(temp,strconv.Iota(rem))
// 	for i:=n-2;i>=0;i=i-1{
// 		c:=str[i]
// 		x:=Atoi(c)
// 		x=x+quo
// 		rem:=x%10
// 		quo=x/10
// 		if i!=0{
// 			temp=temp+Iota(rem)

// 			// temp=append(temp,strconv.Iota(rem))
// 		}else{
// 			rem=Atoi(rem)
// 			rem=reverse(rem)
// 			temp=temp+Iota(rem)
// 			// temp=append(temp,strconv.Iota(rem))
// 		}
// 	}
// 	return reverse(temp)}

// func reverse(s string) (result string) {
//   for _,v := range s {
//     result = string(v) + result
//   }
//   return 
// }

// func Atoi(c char) int{
// 	if c=="0"{
// 		return 0
// 	}
// 	if c=="1"{
// 		return 1
// 	}
// 	if c=="2"{
// 		return 2
// 	}
// 	if c=="3"{
// 		return 3
// 	}
// 	if c=="4"{
// 		return 4
// 	}
// 	if c=="5"{
// 		return 5
// 	}
// 	if c=="6"{
// 		return 6
// 	}
// 	if c=="7"{
// 		return 7
// 	}
// 	if c=="8"{
// 		return 8
// 	}
// 	if c=="9"{
// 		return 9
// 	}

// 	return -1

// }

// func Iota(a int) string{
// 	if a==0{
// 		return "0"
// 	}
// 	if a==1{
// 		return "1"
// 	}
// 	if a==2{
// 		return "2"
// 	}
// 	if a==3{
// 		return "3"
// 	}
// 	if a==4{
// 		return "4"
// 	}
// 	if a==5{
// 		return "5"
// 	}
// 	if a==6{
// 		return "6"
// 	}
// 	if a==7{
// 		return "7"
// 	}
// 	if a==8{
// 		return "8"
// 	}
// 	if a==9{
// 		return "9"
// 	}
// 	return ""

// }