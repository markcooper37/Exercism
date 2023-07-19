(ns sum-of-multiples)

(defn sum-of-multiples [multiples number]
      (->> multiples
           (map #(range % number %))
           flatten
           distinct
           (reduce +))
)
