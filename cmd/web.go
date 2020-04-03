/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"flag"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"

	product "github.com/captainkovalsky/shop/services"
)

type View struct {
	Current         string
	CurrentCategory string
	Categories      product.CategoriesMap
	Products        product.ViewData
}

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run Website",
	Long:  `Run website for processing orders`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("starting ...")
		var wait time.Duration
		flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
		flag.Parse()

		r := mux.NewRouter()
		// Add your routes as needed
		r.HandleFunc("/", productsHandler)
		r.HandleFunc("/", productsHandler).Queries("category", "{category}")
		r.HandleFunc("/bucket", bucketHandler)

		srv := &http.Server{
			Addr: "0.0.0.0:80",
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      r, // Pass our instance of gorilla/mux in.
		}

		// Run our server in a goroutine so that it doesn't block.
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
			}
			log.Println("server started")
		}()

		c := make(chan os.Signal, 1)
		// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
		// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
		signal.Notify(c, os.Interrupt)

		// Block until we receive our signal.
		<-c

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()
		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		shutdownErr := srv.Shutdown(ctx)

		if shutdownErr != nil {
			log.Error("failed to graceful shutdown the applicayion")
			os.Exit(0)
		}
		// Optionally, you could run srv.Shutdown in a goroutine and block on
		// <-ctx.Done() if your application should wait for other services
		// to finalize based on context cancellation.
		log.Println("shutting down")
		os.Exit(0)
	},
}

func bucketHandler(w http.ResponseWriter, r *http.Request) {
	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"templates/layout.html",
		"templates/bucket.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, View{
		Current:    "bucket",
		Categories: product.Categories,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	if category == "" {
		category = "backery"
	}

	log.Println("fetch items for ", category)
	fetcher := product.GoogleSheetFetcher{}

	items := fetcher.Fetch(category)

	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"templates/layout.html",
		"templates/index.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, View{
		Current:         "list",
		CurrentCategory: product.GetTitle(category),
		Categories:      product.Categories,
		Products:        items,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
