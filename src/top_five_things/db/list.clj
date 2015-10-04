(ns top-five-things.db.list
  (:import [java.util Date])
  (:require [taoensso.carmine :as car]
            [clj-time.core :as time]
            [clj-time.format :as timef])
  (:use [top-five-things.db.core :only [wcar*]]))

(defn make-list [id title created-at items]
  {:id id :title title :created-at created-at :items items})

(defn get-list [id]
  (wcar* (car/ping)))

(defn create-list!
  "Creates a list in the connected redis.
   Returns the id of the list or throws an error."
  [title items]
  (let [id (str (wcar* (car/incr "next_list_id")))
        created-at (timef/unparse (timef/formatters :date-time) (time/now))]
    (wcar*
     (car/hmset* (str "list:" id)
                 {:id id
                  :title title
                  :created-at created-at})
     (doseq [item items] (car/rpush (str "list_items:" id) item))
     (car/lpush "lists" id))
    id))

;;(create-list! "Test title", ["one", "two", "three"])
;;(wcar* car/hmget "list:5" :id)
;;(wcar* car/lrange "lists" 0 -1)
;;(t/show-formatters)
;;(wcar* (car/incr "testfoo"))
;;(get-list 1)

