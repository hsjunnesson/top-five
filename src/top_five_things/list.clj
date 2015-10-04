(ns top-five-things.list
  (:require [liberator.core :refer [defresource]]
            [taoensso.carmine :as car :refer (wcar)])
  (:use [top-five-things.util]))

(defresource list-collection-resource
  :allowed-methods [:get]
  :known-content-type? #(check-content-type % ["application/json"])
  :available-media-types ["application/json"]
  :handle-ok (fn [_] [[1], [2], [3, 4]])
  )

(defresource list-resource
  :allowed-methods [:get]
  :known-content-type? #(check-content-type % ["application/json"])
  :exists? (fn [_] {::list [1, 2, 3, 4, 5]})
  :available-media-types ["application/json"]
  :handle-ok ::list)

