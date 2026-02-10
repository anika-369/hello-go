package tasks

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	
)

//load tasks====================
func LoadTasks() []string{
	file, err := os.Open("tasks.txt")
	if err != nil {
		return[]string{}
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks,scanner.Text())
	}
	return tasks
}

////save tasks function===============
func SaveTasks(tasks []string) {
	data := strings.Join(tasks,"\n")
	os.WriteFile("tasks.txt", []byte(data), 0644)
}

//=========addTask==================
func AddTask(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	file, _ := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	file.WriteString(task + "\n")
	fmt.Println("Task saved")
}

//===========deleteTask=============
func DeleteTask(reader *bufio.Reader, tasks *[]string){
	if len(*tasks) == 0 {
		fmt.Println("No tasks to delete")
		return
	}
	for i, task := range *tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
	fmt.Print("Enter task number to delete: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil || num < 1 ||num > len (*tasks) {
		fmt.Println("Invalid number")
		return
	}
	index := num - 1 
	*tasks = append((*tasks)[:index],(*tasks)[index+1:]...)

	SaveTasks(*tasks)
	fmt.Println("Task deleted!")
}

//============ListTask=================
func ListTasks() {
	file, err :=  os.Open("tasks.txt")
	if err != nil {
		fmt.Println("no tasks found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1

	for scanner.Scan() {
		fmt.Println(count, scanner.Text())
		count++
	}
}


