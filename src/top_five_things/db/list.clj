(ns top-five-things.db.list
  (:import [java.util Date])
  (:require [taoensso.carmine :as car]
            [clj-time.core :as time]
            [clj-time.format :as timef])
  (:use [top-five-things.db.core :only [wcar*]]))

(defn create-list!
  "Creates a list in the connected redis.
   Returns the id of the list or throws an error."
  [title items]
  (let [id (str (wcar* (car/incr "next_list_id")))
        created-at (timef/unparse (timef/formatters :date-time) (time/now))]
    (wcar*
     (car/hmset (str "list:" id)
                 :id id
                 :title title
                 :created-at created-at)
     (doseq [item items] (car/rpush (str "list_items:" id) item))
     (car/lpush "lists" id))
    id))

(defn get-list [id]
  "Returns a list from the connected redis."
  (let [hash-key (str "list:" id)]
    {:id id
     :title (wcar* (car/hget hash-key :title))
     :created-at (timef/parse (timef/formatters :date-time) (wcar* (car/hget hash-key :created-at)))
     :items (wcar* (car/lrange (str "list_items:" id) 0 -1))}))

