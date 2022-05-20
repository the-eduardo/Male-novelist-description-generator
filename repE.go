package main

import ("fmt"
      "math/rand"
      "time")

func repE() string {

  rand.Seed(time.Now().UnixNano())
  rep1 := rand.Intn(26)+6
  fmt.Println(rep1)

 if rep1 == 1	{return "ravish"
} else if rep1 == 2 {
  return "climb"
} else if rep1 == 3 {
  return "invade"
} else if rep1 == 4 {
  return "grope"
} else if rep1 == 6 {
  return "marry"
} else if rep1 == 7 {
  return "raw dog it with"
} else if rep1 == 5 {
  return "caress"
} else if rep1 == 8 {
  return "proposition"
} else if rep1 == 9 {
  return "correct"
} else if rep1 == 10 {
  return "emotionally manipulate"
} else if rep1 == 11 {
  return "spar with"
} else if rep1 == 12 {
  return "compliment"
} else if rep1 == 13 {
  return "hire"
} else if rep1 == 14 {
  return "booty call"
} else if rep1 == 15 {
  return "mansulain to"
} else if rep1 == 16 {
  return "Insult"
} else if rep1 == 17 {
  return "dale"
} else if rep1 == 18 {
  return "teabag"
} else if rep1 == 19 {
  return "ignore"
} else if rep1 == 20 {
  return "fondle"
} else if rep1 == 21 {
  return "worship"
} else if rep1 == 22 {
  return "examine"
} else if rep1 == 23 {
  return "manhandle"
} else if rep1 == 24 {
  return "touch"
} else if rep1 == 25 {
  return "admire"
} else if rep1 == 26 {
    return "demand things from"
  } else if rep1 == 27 {
    return "ravish"
  } else if rep1 == 28 {
    return "climb"
  } else if rep1 == 29 {
    return "invade"
  } else if rep1 == 30 {
    return "grope"
  } else if rep1 == 31 {
    return "marry"
  } else if rep1 == 32 {
  return "raw dog it with"
  }  else if rep1 == 0 {return "BUG"}
    return fmt.Sprintf("teste")}
