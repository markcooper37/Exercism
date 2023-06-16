(ns bird-watcher)

(def last-week [0 2 5 3 7 8 4])

(defn today [birds]
  (last birds)
  )

(defn inc-bird [birds]
  (conj (pop birds) (+ (last birds) 1))
  )

(defn day-without-birds? [birds]
  (boolean (some #{0} birds))
  )

(defn n-days-count [birds n]
  (reduce + (subvec birds 0 n))
  )

(defn busy-days [birds]
  (count (filter #(>= % 5) birds))
  )

(defn odd-week? [birds]
  (= birds [1 0 1 0 1 0 1])
  )
