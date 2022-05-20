package main
import ("fmt"
      "math/rand"
      "time")

func repC() string {

  rand.Seed(time.Now().UnixNano())
  rep1 := rand.Intn(26)+1
  fmt.Println(rep1)


  if rep1 == 1 {
  return "kitten"
} else if rep1 == 2 {
  return "mountain"
} else if rep1 == 3 {
  return "pillow"
} else if rep1 == 4 {
  return "ice cream cone"
} else if rep1 == 5 {
  return "tulip"
} else if rep1 == 6 {
  return "lake"
} else if rep1 == 7 {
  return "fortress"
} else if rep1 == 8 {
  return "lemon"
} else if rep1 == 9 {
  return "popsicle"
} else if rep1 == 10 {
  return "librarian"
} else if rep1 == 11 {
  return "python"
} else if rep1 == 12 {
  return "pony"
} else if rep1 == 13 {
  return "melon"
} else if rep1 == 14 {
  return "bedsheet"
} else if rep1 == 15 {
  return "muffin"
} else if rep1 == 16 {
  return "bunny rabbit"
} else if rep1 == 17 {
  return "fish"
} else if rep1 == 18 {
  return "princess"
} else if rep1 == 19 {
  return "ghost"
} else if rep1 == 20 {
  return "waterfall"
} else if rep1 == 21 {
  return "mango"
} else if rep1 == 22 {
  return "harpy"
} else if rep1 == 23 {
  return "ship"
} else if rep1 == 24 {
  return "nun"
} else if rep1 == 25 {
  return "berry"
} else if rep1 == 26 {
  return "car"
  }  else if rep1 == 0 {return "BUG"}
    return fmt.Sprintf("teste")}
