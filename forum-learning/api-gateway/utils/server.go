package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	Host string
	Port string
}

type config struct {
	AppName      string
	SaltKey      []byte
	SecretKey    []byte
	State        string
	EndpointsMap map[string]Endpoint
}

// Server Struct
type Server struct {
	Config     config
	Router     *gin.Engine
	APIRouter  *gin.RouterGroup
	HTTPClient *http.Client
}

// NewServer is a constructor for Server Struct
func NewServer() (*Server, error) {

	server := Server{}

	err := server.readConfigFile()

	if err != nil {
		return nil, err
	}

	server.Config.EndpointsMap = make(map[string]Endpoint)

	err = server.readEndpointMapsConfig()

	if err != nil {
		return nil, err
	}

	gin.SetMode(server.Config.State)
	server.Router = gin.Default()
	server.APIRouter = server.Router.Group("api")
	server.HTTPClient = &http.Client{}

	return &server, nil
}

func (s *Server) readConfigFile() error {
	s.Config.AppName = "Forum_Learning-API_GATEWAY"
	s.Config.State = os.Getenv("GIN_MODE")

	s.Config.SaltKey = []byte("d\x8f\xef\x83`\xb1*\xd5[\xedu\xdb0\x8bJ\x94\xe0\xf0\xa5\xf1\x91\xc7t\xa0")
	s.Config.SecretKey = []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	return nil
}

func (s *Server) readEndpointMapsConfig() error {

	if _, err := os.Stat("./config/config.json"); os.IsNotExist(err) {
		return fmt.Errorf("Config File Not Found")
	}

	jsonFile, err := os.Open("./config/config.json")

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	var result map[string][]map[string]string

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		return err
	}

	config, ok := result["endpoints_map"]

	if !ok {
		return fmt.Errorf("Config File Format Invalid")
	}

	for i := 0; i < len(config); i++ {

		prefix := config[i]["prefix"]
		host := config[i]["host"]
		port := config[i]["port"]

		s.Config.EndpointsMap[prefix] = Endpoint{
			Host: host,
			Port: port,
		}
	}

	return nil

}
