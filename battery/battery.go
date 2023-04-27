package battery

import (
    "fmt"
    "io/ioutil"
    "strconv"
)

func Percentage() int{
    // Read the battery percentage from the sysfs filesystem
    content, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
    if err != nil {
        panic(err)
    }

    // Convert the battery percentage from a string to an integer
    percentage, err := strconv.Atoi(string(content[:len(content)-1]))
    if err != nil {
        panic(err)
    }
    // Print the battery percentage
    fmt.Printf("Battery percentage: %d%%\n", percentage)
    return percentage
}
