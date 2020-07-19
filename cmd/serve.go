/*
Copyright © 2020 Søren Nielsen <contact@cph.dev>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/sorennielsen/bmi/internal/bmi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var _, example, _ = bmi.CalculateWithoutStats("186", "85")

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs BMI in service mode waiting for web requests.",
	Long: `Using the 'serve' command BMI starts up as a web service.

Get BMI calculated by hitting the path /calc/<height in cm>/<weight in kg>
Example: /calc/186/85

Output: ` + example,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetUint("port")
		fmt.Printf("Starting web service on port %v\n", port)
		fmt.Printf("\thttp://localhost:%d\n", port)
		serve(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().UintP("port", "p", 8080, "Port for server to listen to.")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "BMI web service\nAdd /calc/<height>/<weight> to URL to calculate BMI.")
}

func Info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	counts := bmi.GetCounts()
	j := json.NewEncoder(w)
	j.Encode(counts)

}

func Calc(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	_, desc, err := bmi.Calculate(ps.ByName("height"), ps.ByName("weight"))
	if errors.Is(err, bmi.BMITooLow) {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to process request: %s\n", err)
		return
	}
	if errors.Is(err, bmi.BMITooHigh) {
		fmt.Printf("Error: %s\n", err)
		fmt.Println("Shutting down!")
		os.Exit(1)
		return
	}
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	fmt.Fprintf(w, desc)
}

func serve(port uint) {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/info", Info)
	router.GET("/calc/:height/:weight", Calc)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
