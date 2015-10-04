(ns top-five-things.handler
  (:require
   [ring.adapter.jetty :as jetty])
  (:require [clojure.java.io :as io]
            [liberator.core :refer [resource defresource]]
            [liberator.dev :refer [wrap-trace]]
            [ring.middleware.params :refer [wrap-params]]
            [compojure.core :refer [context ANY routes defroutes]]
            [compojure.handler :refer [api]]
            [environ.core :refer [env]]
            [top-five-things.db :as db])
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
    (ANY ["/lists/:id" :id #".*"] [id] (list-resource id))
    (ANY "/lists" [] list-collection-resource))))

(def handler
  (-> (assemble-routes)
      (wrap-trace :header :ui)
      api
      wrap-params))

(defn start [options]
  (db/init)
  (jetty/run-jetty #'handler (assoc options :join? false)))

(defn -main [& [port]]
  (let [port (Integer. (env :port))
        redis-url (env :redis-url)]
    (if (not port) (throw (Exception. "$PORT env not set.")))
    (if (not redis-url) (throw (Exception. "$REDIS_URL env not set")))
    (db/init redis-url)
    (jetty/run-jetty #'handler {:port port :join? false})))

