package harbor

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var path = "E:\\project\\k8s\\helm\\csmp\\kmc-0.1.0.tgz"

func TestClient(t *testing.T) {
	encoded := base64.StdEncoding.EncodeToString([]byte("admin" + ":" + "Harbor12345"))
	fmt.Println("encoded :=",encoded )
	fmt.Println("encoded :=",111 )
	ok,err := UploadChart(path)
	if err != nil{
		fmt.Println("err",err)
	}
	fmt.Println("ok",ok)

}

func TestRepoList(t *testing.T) {
	ListRepositories()
}
func TestGetRepository(t *testing.T) {
	ListArtifacts("csmp-backend")
}