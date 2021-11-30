/* RZFeeser | Alta3 Research
   CHALLEGE 01 - Create a struct named Astro */

package main

import "fmt"
// declare the struct and populate the variables as string, int and bool
type astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

func main() {
// create a slice for each astronaut and populate with variable data
    ast1 := astro{"Megan McArthur", 35, "ISS", false}
    ast2 := astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := astro{"Raja Chari", 39, "SpaceX Crew-3", true}
// print the output of each slice
    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)

}

