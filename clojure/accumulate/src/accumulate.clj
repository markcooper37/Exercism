(ns accumulate)

(defn accumulate [function args]
      (if (empty? args)
        []
          (let [head (first args) tail (rest args)]
            (cons (function head) (accumulate function tail))))
)
