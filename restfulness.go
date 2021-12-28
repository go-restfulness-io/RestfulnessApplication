package restfulnessapplication

import (
	"net/http"
	"os"
)

func StartApplication() {
	println("Starting Restfulness Application ...")
	for atom, atomHandler := range Service.GetAtomRegistry() {
		println("Perform Mapping")
		atom.SetStage(Service.MAPPING)
		println("Call Handler")
		atomMapper := *Service.NewAtomMapper(atom, atomHandler)
		atomHandler.DoMapping(&atomMapper)
		//atomHandler.AtomHandler(&atom, new(Service.Request))
	}

	err := http.ListenAndServe(":8080", &HttpHandler{})
	if err != nil {
		println("cannot start http server")
		os.Exit(100)
	}

}
