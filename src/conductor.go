package conductor

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const Version = "0.0.1"
const VersionName = "Lifeboat"

type Config struct {
	Port           int64  `json:"port"`
	StringPort     string `json:"stringport"`
	ViewsDirectory string `json:"viewsDirectory"`
}

type ConductorConfig struct {
	Port     int64          `json:"port"`
	Database DatabaseConfig `json:"database"`
}
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var Conf Config
var ConductorConf ConductorConfig
var Templates *template.Template

func Init() error {
	Conf.StringPort = ":" + strconv.FormatInt(Conf.Port, 10)
	return nil
}

func RefreshTemplates() {
	Templates = template.Must(template.ParseGlob(Conf.ViewsDirectory))
}

func ReadConfig() {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {

	}
	json.Unmarshal(data, &ConductorConf)
}

func StartServer() error {
	fmt.Println("🤔  Conductor: Version", Version, VersionName)
	fmt.Println("Running now on port", Conf.Port)

	ReadConfig()

	RefreshTemplates()
	r := mux.NewRouter()

	err := Connect("")
	if err != nil {
		panic(err)
	}

	assets := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets").Handler(assets)

	r.HandleFunc("/", utilPingHandler)
	r.HandleFunc("/members", membersHandler)
	r.HandleFunc("/register", authRegisterHandler).Methods("GET")
	r.HandleFunc("/register", authRegisterProcessHandler).Methods("POST")
	r.HandleFunc("/login", authLoginHandler).Methods("GET")
	r.HandleFunc("/login", authLoginProcessHandler).Methods("POST")
	r.HandleFunc("/logout", authLogoutHandler)
	r.HandleFunc("/recover", authRecoverHandler).Methods("GET")
	r.HandleFunc("/recover", authRecoverProcessHandler).Methods("POST")
	r.HandleFunc("/forums", forumsHandler)
	r.HandleFunc("/forums/post/{forum_guid}", NewTopicHandler).Methods("GET")
	r.HandleFunc("/forums/post/{forum_guid}", NewTopicProcessHandler).Methods("POST")
	r.HandleFunc("/forums/{forum_guid}", forumHandler)
	r.HandleFunc("/forums/{forum_key}/{topic_key}", TopicHandler).Methods("GET")
	r.HandleFunc("/forums/{forum_key}/{topic_key}", NewPostProcessHandler).Methods("POST")
	r.HandleFunc("/", utilPingHandler)
	r.HandleFunc("/forum/{form_guid}", utilPingHandler)
	r.HandleFunc("/topic/{topic_guid}", utilPingHandler)
	r.HandleFunc("/topic/{topic_guid}/{page}", utilPingHandler)

	r.HandleFunc("/ping.html", utilPingHandler)

	http.Handle("/", r)
	http.ListenAndServe(Conf.StringPort, r)
	return nil
}
