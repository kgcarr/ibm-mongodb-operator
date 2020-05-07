package resources

import (
  "testing"
  "fmt"
  )

//Test Making ICP MongoDB Service Function
func TestMakeICPMongoDBService(t *testing.T) {
    myService := MakeICPMongoDBService()

    if myService != nil {
      fmt.Println(myService.String())
    } else {
      t.Errorf("Did not create ICP MongoDB Service Object")
    }
}

//Test Making MongoDB Service Function
func TestMakeMongoDBService(t *testing.T) {
    myService := MakeMongoDBService()

    if myService != nil {
      fmt.Println(myService.String())
    } else {
      t.Errorf("Did not create MongoDB Service Object")
    }
}

//Test Making the Init Configmap
func TestMakeInitConfigmap(t *testing.T) {
  myConfigmap := MakeInitConfigmap()

  if myConfigmap != nil {
    fmt.Println(myConfigmap.String())
  } else {
    t.Errorf("Did not create MongoDB Init ConfigMap Object")
  }
}

//Test Making the Install Configmap
func TestMakeInstallConfigmap(t *testing.T) {
  myConfigmap := MakeInstallConfigmap()

  if myConfigmap != nil {
    fmt.Println(myConfigmap.String())
  } else {
    t.Errorf("Did not create MongoDB Install ConfigMap Object")
  }
}

//Test Making the mongod.conf Configmap
func TestMakeInstallConfigmap(t *testing.T) {
  myConfigmap := MakeMongodConfConfigmap()

  if myConfigmap != nil {
    fmt.Println(myConfigmap.String())
  } else {
    t.Errorf("Did not create MongoDB mongod.conf ConfigMap Object")
  }
}
