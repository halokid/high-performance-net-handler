package utils

func CheckErr(err error) {
  if err != nil {
    panic(err.Error())
  }
}
