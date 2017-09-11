package imageHide

import (
  "os"
  "io/ioutil"
  "os/exec"
)

func WriteMessage(message []byte) error {
  err := ioutil.WriteFile("message.txt", message, 0644)
  if err != nil {
    return err
  } else {
    cmd := exec.Command("python", "pythonWriteImg.py")
    _, err := cmd.CombinedOutput()
    if err != nil {
      fmt.Prinln("Error while read the image.")
    }
    return nil
  }
}

func readImage() (string, error) {
  cmd := exec.Command("python", "pythonReadImg.py")
  _, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Prinln("Error while read the image.")
    os.Exit(1)
  }

  message, err := ioutil.ReadFile("messageOut.txt")
  if err != nil {
    return "", err
  } else {
    return string(message), nil
  }
}
