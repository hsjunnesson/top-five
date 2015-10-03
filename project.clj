(defproject top-five-things "1.0.0-SNAPSHOT"
  :description "Web server and API for www.top-five-things.com"
  :url "https://www.top-five-things.com"
  :license {:name "The MIT License (MIT)"
            :url "https://opensource.org/licenses/MIT"}
  :dependencies [[org.clojure/clojure "1.7.0"]
                 [compojure "1.4.0"]
                 [ring/ring-jetty-adapter "1.4.0"]
                 [environ "1.0.1"]]
  :min-lein-version "2.0.0"
  :plugins [[environ/environ.lein "0.3.1"]]
  :hooks [environ.leiningen.hooks]
  :uberjar-name "top-five-things.jar"
  :profiles {:production {:env {:production true}}})
