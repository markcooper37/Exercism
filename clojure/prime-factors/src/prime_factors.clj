(ns prime-factors)

(defn of [number]
  (loop [number number
         candidate 2
         factors []]
    (cond
      (= number 1) factors
      (zero? (mod number candidate))
        (recur (/ number candidate) candidate (conj factors candidate))
      :else
        (recur number (inc candidate) factors))))
