;; (defproject top-five-things "1.0.0-SNAPSHOT"
;;   :description "Web server and API for www.top-five-things.com"
;;   :url "https://www.top-five-things.com"
;;   :license {:name "The MIT License (MIT)"
;;             :url "https://opensource.org/licenses/MIT"}
;;   :dependencies [[org.clojure/clojure "1.10.1"]
;;                  [compojure "1.6.1"]
;;                  [ring/ring-defaults "0.3.2"]
;;                  [environ "1.1.0"]
;;                  [liberator "0.15.3"]
;;                  [hiccup "1.0.5"]]
;;   :min-lein-version "2.0.0"
;;   :plugins [[environ/environ.lein "0.3.1"]
;;             [lein-ring "0.12.5"]]
;;   :ring {:handler top-five-things/app}
;;   :hooks [environ.leiningen.hooks]
;;   :uberjar-name "top-five-things.jar"
;;   :profiles
;;   {:dev {:dependencies [[javax.servlet/servlet-api "2.5"]
;;                         [ring/ring-mock "0.3.2"]]}})

(defproject top-five "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :min-lein-version "2.0.0"
  :dependencies [[org.clojure/clojure "1.10.0"]
                 [compojure "1.6.1"]
                 [ring/ring-defaults "0.3.2"]
                 [ring/ring-jetty-adapter "1.7.1"]
                 [environ "1.1.0"]
                 [liberator "0.15.3"]
                 [hiccup "1.0.5"]
                 [com.taoensso/carmine "2.19.1"]]
  :plugins [[lein-ring "0.12.5"]
            [environ/environ.lein "0.3.1"]]
  :hooks [environ.leiningen.hooks]
  :ring {:handler top-five-things.handler/app}
  :uberjar-name "top-five-things.jar"
  :profiles {:dev {:dependencies [[javax.servlet/servlet-api "2.5"]
                                  [ring/ring-mock "0.3.2"]]}})
