(ns perfect-numbers)

(defn factor-sum [number]
      (reduce + (filter #(zero? (rem number %)) (range 1 number)))
)

(defn classify [number]
      (if (< number 1) 
        (throw (IllegalArgumentException. "number cannot be negative")) 
        (let [sum (factor-sum number)] 
          (cond (> sum number) :abundant
            (< sum number) :deficient
            :else :perfect)))
)

