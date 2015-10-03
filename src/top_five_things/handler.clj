(ns top-five-things.handler
  (:require
   [ring.adapter.jetty :as jetty])
  (:require [clojure.java.io :as io]
            [liberator.core :refer [resource defresource]]
            [liberator.dev :refer [wrap-trace]]
            [ring.middleware.params :refer [wrap-params]]
            [compojure.core :refer [context ANY routes defroutes]]
            [compojure.handler :refer [api]])
  (:use [top-five-things.index]
        [top-five-things.list]
        [top-five-things.util]
        [ring.util.mime-type :only [ext-mime-type]]))

(let [static-dir (io/file "static")]
  (defresource static-resource

    :available-media-types
    #(let [path (get-in % [:request :route-params :*])]
       (if-let [mime-type (ext-mime-type path)]
         [mime-type]
         []))

    :exists?
    #(let [path (get-in % [:request :route-params :*])]
       (let [f (io/file static-dir path)]
         [(.exists f) {::file f}]))

    :handle-ok (fn [{f ::file}] f)

    :last-modified (fn [{f ::file}] (.lastModified f))))

(defn assemble-routes []
  (->
   (routes
    (ANY "/" [] index-resource)
    (ANY "/static/*" [] static-resource)
    (ANY ["/lists/:id" :id #".*"] [id] (list-resource id)))))

(def handler
  (-> (assemble-routes)
      (wrap-trace :header :ui)
      api
      wrap-params))

(defn start [options]
  (jetty/run-jetty #'handler (assoc options :join? false)))

(defn -main
  ([port]
     (start {:port (Integer/parseInt port)}))
  ([]
     (-main "8000")))

