(ns top-five-things.index
  (:require [liberator.core :refer [resource defresource]])
  (:use [hiccup.page :only [html5]]))

(def index-body
  {:_links
   {:self {:href "/"}
    :lists {:href "/lists"}}})

(defresource index-resource
  :available-media-types ["text/html" "application/json"]
  :handle-ok
  (fn [ctx]
    (case (get-in ctx [:representation :media-type])
      "text/html" (html5 [:body [:p {:style "color: green"} "OK!"]])
      "application/json" index-body)))

