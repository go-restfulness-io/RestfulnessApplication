package restfulnessapplication

import (
	"github.com/go-restfulness-io/restfulnessapplication/service"
	"net/http"
	"os"
)

func StartApplication() {
	println("Starting Restfulness Application ...")
	for atom, atomHandler := range service.GetAtomRegistry() {
		println("Perform Mapping")
		atom.SetStage(service.MAPPING)
		println("Call Handler")
		atomMapper := *service.NewAtomMapper(atom, atomHandler)
		atomHandler.DoMapping(&atomMapper)
		//atomHandler.AtomHandler(&atom, new(Service.Request))
	}

	err := http.ListenAndServe(":8080", &HttpHandler{})
	if err != nil {
		println("cannot start http server")
		os.Exit(100)
	}

}
