(ns top-five-things.handler
  (:require
   [ring.adapter.jetty :as jetty])
  (:require [clojure.java.io :as io]
            [liberator.core :refer [resource defresource]]
            [liberator.dev :refer [wrap-trace]]
            [ring.middleware.params :refer [wrap-params]]
            [ring.util.response :refer [resource-response]]
            [compojure.core :refer [context ANY routes defroutes]]
            [compojure.handler :refer [api]]
            [environ.core :refer [env]])
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

;; Ignore favicon.ico requests
;;
;; (wrap-ignore-favicon-request [handler]
;;
(defn wrap-ignore-favicon-request [handler]
  (fn [request]
    (if (= (:uri request) "/favicon.ico")
      {:status 404}
      (handler request))))

(defn assemble-routes []
  (->
   (routes
    (ANY "/" [] index-resource)
    (ANY "/static/*" [] static-resource)
    (ANY ["/lists/:id" :id #".*"] [id] (list-resource id)))))

(def app
  (-> (assemble-routes)
      (wrap-trace :header :ui)
      wrap-ignore-favicon-request
      api
      wrap-params))

(defn start [options]
  (jetty/run-jetty #'app (assoc options :join? false)))

(defn -main [& [port]]
  (let [port (Integer. (or port (env :port) 5000))]
    (start {:port port})))

