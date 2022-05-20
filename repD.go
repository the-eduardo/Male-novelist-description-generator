
package main
import ("fmt"
      "math/rand"
      "time")

func repD() string {

  rand.Seed(time.Now().UnixNano())
  rep1 := rand.Intn(26)+1
  fmt.Println(rep1)


if rep1 ==1 {
  return "longed"
} else if rep1 ==3{
  return "lusted"
} else if rep1 ==4{
  return "wished"
} else if rep1 ==5{
  return "wanted"
} else if rep1 ==6{
  return "resolved"
} else if rep1 ==7{
  return "dared"
} else if rep1 ==8{
  return "detested"
} else if rep1 ==9{
  return "trembled"
} else if rep1 ==10{
  return "thirsted"
} else if rep1 ==11{
  return "expected"
} else if rep1 ==12{
  return "planned"
} else if rep1 ==13{
  return "deigned"
} else if rep1 ==14{
  return "proposed"
} else if rep1 ==15{
  return "shuddered"
} else if rep1 ==16{
  return "refused"
} else if rep1 ==17{
  return "hated"
} else if rep1 ==18{
  return "scorned"
} else if rep1 ==19{
  return "dreaded"
} else if rep1 == 2 {
  return "did not care"
} else if rep1 == 20 {
  return "ached"
} else if rep1 == 21 {
  return "intended"
} else if rep1 == 22 {
  return "hungered"
} else if rep1 == 23 {
  return "W  craved"
} else if rep1 == 24 {
  return "hankered"
} else if rep1 == 25 {
  return "needed"
} else if rep1 == 26 {
  return "pined"
  }  else if rep1 == 0 {return "BUG"}
    return fmt.Sprintf("teste")}
