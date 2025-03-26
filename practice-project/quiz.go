package main
import "fmt"


func main(){
	var name string
	var ans string
	var score int
	fmt.Println("Hello!!")
	fmt.Print("May I know your name? : ")
	fmt.Scan(&name)
	fmt.Println("Hello!", name, "Are you ready to start the Quiz......!!!")
	fmt.Print("Enter (Y) for yes and (N) for no: ")
	fmt.Scan(&ans)

	if(ans == "Y" || ans == "y"){
		fmt.Println("Let's begin the quiz!")
		fmt.Println("Q1: 1 + 1 = ? ")
		var userAns int
		fmt.Print("Write Answer:")
		fmt.Scan(&userAns)
		if(userAns == 2){
			score++
            fmt.Println("Correct Answer!")
        }else {
            fmt.Println("Wrong Answer!")
        }

		fmt.Println("Q2: 2 * 3 =? ")
        fmt.Print("Write Answer:")
        fmt.Scan(&userAns)
        if(userAns == 6){
            score++
            fmt.Println("Correct Answer!")
        } else {
            fmt.Println("Wrong Answer!")
        }

		fmt.Println("Your Score : ", score)
	}else {
		fmt.Println("Okay, Quiz is over!")
        return
	}



}