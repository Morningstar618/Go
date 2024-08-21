package main

import (
	"fmt"

	str "example.com/structs/custom_string"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your date of birth: ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("aj2000@gmail.com", "1212")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	fmt.Println("---------------------")

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	fmt.Println("---------------------")

	var name str.Str = "Ayush"
	name.Log()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
