package main
import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"go-todo/internal/tasks"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	taskList := tasks.LoadTasks()

	for {
		fmt.Println("\nOptions: add | list | delete | exit")
		fmt.Print("Enter command: ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		
		switch command {
		case "add":
			tasks.AddTask(reader)
			taskList = tasks.LoadTasks()
		case "list":
			tasks.ListTasks()
		case "delete":
			tasks.DeleteTask(reader, &taskList)
		case "exit":
			fmt.Println("bye")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
} 
