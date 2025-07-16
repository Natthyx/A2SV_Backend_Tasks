package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter number of subjects: ")
    subjectCountStr, _ := reader.ReadString('\n')
    subjectCountStr = strings.TrimSpace(subjectCountStr)
    subjectCount, err := strconv.Atoi(subjectCountStr)
    if err != nil || subjectCount <= 0 {
        fmt.Println("Invalid number of subjects.")
        return
    }

    grades := make(map[string]float64)

    for i := 0; i < subjectCount; i++ {
        fmt.Printf("Enter name of subject #%d: ", i+1)
        subject, _ := reader.ReadString('\n')
        subject = strings.TrimSpace(subject)

        fmt.Printf("Enter grade for %s (0-100): ", subject)
        gradeStr, _ := reader.ReadString('\n')
        gradeStr = strings.TrimSpace(gradeStr)
        grade, err := strconv.ParseFloat(gradeStr, 64)
        if err != nil || grade < 0 || grade > 100 {
            fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
            i-- 
            continue
        }

        grades[subject] = grade
    }

    average := calculateAverage(grades)

    fmt.Printf("\nHello, %s!\n", name)
    fmt.Println("Here are your grades:")
    for subject, grade := range grades {
        fmt.Printf("- %s: %.2f\n", subject, grade)
    }
    fmt.Printf("Your average grade is: %.2f\n", average)
}


func calculateAverage(grades map[string]float64) float64 {
    var total float64 = 0
    for _, grade := range grades {
        total += grade
    }
    return total / float64(len(grades))
}
