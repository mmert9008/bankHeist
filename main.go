package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  
  openedVault := rand.Intn(100)

  if isHeistOn == true && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if openedVault < 70 {
    isHeistOn = false
    fmt.Println("The vault can't be opened.")
  }

  leftSafely := rand.Intn(5)

  if isHeistOn {
    switch leftSafely {
    case 0:
      isHeistOn = false
      fmt.Println("You failed to get away.")
    case 1:
      isHeistOn = false
      fmt.Println("You got caught.")
    case 2:      
      isHeistOn = false
      fmt.Println("You didn't get away.")
    case 3:      
      isHeistOn = false
      fmt.Println("You were caught.")
    default:
      fmt.Println("Start the getaway car!")
    }
  }
  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Printf("Amount stolen: $%d\n", amtStolen)
  }

  //fmt.Println(isHeistOn)
}

