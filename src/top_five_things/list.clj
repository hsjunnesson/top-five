(ns top-five-things.list
  (:require [liberator.core :refer [defresource]])
  (:use [top-five-things.util]))

(defresource list-resource [id]
  :allowed-methods [:get]
  :known-content-type? #(check-content-type % ["application/json"])
  :exists? (fn [_] {::list [1, 2, 3, 4, 5]})
  :available-media-types ["application/json"]
  :handle-ok ::list)

