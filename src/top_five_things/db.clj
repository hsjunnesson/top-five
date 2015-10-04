(ns top-five-things.db
  (:require [taoensso.carmine :as car :refer (wcar)]))

(defn init [redis-uri]
  (defmacro wcar* [& body] `(car/wcar {:pool {} :spec {:uri redis-uri}} ~@body))
  )

