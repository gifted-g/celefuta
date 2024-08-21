package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "time"
)

func main() {
    fmt.Println("\nData Extractor\n")
    fmt.Println("Reading Biodata from file 'CCCSP FUTA Final-Year Brethren Biodata.csv'... (Ensure the Google form document is saved in the same folder as this program with exactly this name: CCCSP FUTA Final-Year Brethren Biodata.csv)\n")

    // Read the CSV file
    source, err := ioutil.ReadFile("CCCSP FUTA Final-Year Brethren Biodata.csv")
    if err != nil {
        fmt.Printf("FAILED: %v\n", err)
        time.Sleep(15 * time.Second)
        return
    }

    // Create Personal directory if it doesn't exist
    if _, err := os.Stat("Personal"); os.IsNotExist(err) {
        os.Mkdir("Personal", 0755)
    }

    members := strings.Split(string(source), "\"\n\"")
    title := strings.Split(members[0], "\",\"")

    for m := 1; m < len(members); m++ {
        member := members[m]
        biodata := strings.Split(member, "\",\"")
        name := strings.Title(strings.TrimSpace(biodata[2]) + " " + strings.TrimSpace(biodata[3]))
        filePath := "Personal/" + name + ".txt"

        // Create and write to the output file
        output, err := os.Create(filePath)
        if err != nil {
            fmt.Printf("FAILED: %v\n", err)
            continue
        }
        writer := bufio.NewWriter(output)

        writer.WriteString("Name:\n" + name + "\n\n\n")
        writer.WriteString("Email Address:\n" + strings.TrimSpace(biodata[1]) + "\n\n\n")

        for n := 4; n < len(biodata)-1; n++ {
            entry := strings.TrimSpace(biodata[n])
            if len(entry) > 0 {
                writer.WriteString(title[n] + ":\n" + entry + "\n\n\n")
            }
        }

        writer.Flush()
        output.Close()
    }

    fmt.Println("SUCCESS: All data extracted successfully and saved in the folder 'Personal/'.")
    time.Sleep(5 * time.Second)
}
