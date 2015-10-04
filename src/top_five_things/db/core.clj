(ns top-five-things.db.core
  (:require [taoensso.carmine :as car :refer (wcar)]))

(defonce ^:dynamic spec (atom {}))

(defn connect! [redis-uri]
  (swap! spec assoc :uri redis-uri))

(defmacro wcar* [& body] `(car/wcar {:pool {} :spec @spec} ~@body))

