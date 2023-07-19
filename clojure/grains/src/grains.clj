(ns grains)

(defn square [number]
    (reduce *' (repeat (- number 1) 2))
)

(defn total []
    (reduce + (map square (range 1 (inc 64))))
)
