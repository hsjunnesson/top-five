(ns top-five-things.list
  (:require [liberator.core :refer [defresource]]
            [taoensso.carmine :as car :refer (wcar)]
            [top-five-things.db.list :as db])
  (:use [top-five-things.util]))


(defresource list-collection-resource
  :allowed-methods [:get]
  :known-content-type? #(check-content-type % ["application/json"])
  :available-media-types ["application/json"]
  :handle-ok (fn [_] [[1], [2], [3, 4]])
  )

(defresource list-resource [id]
  :allowed-methods [:get]
  :known-content-type? #(check-content-type % ["application/json"])
  :exists? (fn [_] {::list (db/get-list id)})
  :available-media-types ["application/json"]
  :handle-ok ::list)

