(ns leap)

(defn leap-year? [year] ;; <- argslist goes here
  ;; your code goes here
  (and (zero? (mod year 4)) (or (zero? (mod year 400)) (not (zero? (mod year 100)))) )
)
