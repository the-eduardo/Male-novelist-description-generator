package main

import ("fmt"
      "math/rand"
      "time")

func repB() string {

  rand.Seed(time.Now().UnixNano())
  rep1 := rand.Intn(26)+1
  fmt.Println(rep1)

   if rep1 == 1	{return "silken"
   } else if rep1 == 2 {
     return "wiry"
} else if rep1 == 3	{
  return "gleaming"
} else if rep1 == 4	{
  return "withholding"
} else if rep1 == 5	{
  return "bulbous"
} else if rep1 == 6	{
  return "brittle"
} else if rep1 == 7	{
  return "soft"
} else if rep1 == 8	{
  return "dewy"
} else if rep1 == 9	{
  return "fat"
} else if rep1 == 10	{
  return "shrill"
} else if rep1 == 11	{
  return "haughty"
} else if rep1 == 12	{
  return "expensive"
} else if rep1 == 13	{
  return "plump"
} else if rep1 == 14	{
  return "withered"
} else if rep1 == 15	{
  return "tempestuous"
} else if rep1 == 16	{
  return "voluptuous"
} else if rep1 == 17	{
  return "silken"
} else if rep1 == 18	{
  return "frigid"
} else if rep1 == 19	{
  return "luscious"
} else if rep1 == 20	{
  return "wrinkled"
} else if rep1 == 21	{
  return "rippling"
} else if rep1 == 22	{
  return "muddle-aged"
} else if rep1 == 23	{
  return "juicy"
} else if rep1 == 24	{
  return "ripe"
} else if rep1 == 25	{
  return "shiny"
} else if rep1 == 26	{
  return "creamy"
}  else if rep1 == 0 {return "BUG"}
  return fmt.Sprintf("teste")}
