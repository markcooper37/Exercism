(ns robot-name)

(def letters ["A" "B" "C" "D" "E" "F" "G" "H" "I" "J" "K" "L" "M" "N" "O" "P" "Q" "R" "S" "T" "U" "V" "W" "X" "Y" "Z"])

(def used-names (atom #{}))

(defn generate-name []
  (let [name (str (rand-nth letters) (rand-nth letters) (rand-int 10) (rand-int 10) (rand-int 10))]
    (if (get @used-names name)
      (recur)
      (do
        (swap! used-names conj name)
        name))))

(defn robot []
  (atom {:name (generate-name)})
)

(defn robot-name [robot]
  (get @robot :name)
)

(defn reset-name [robot]
  (swap! robot #(assoc % :name (generate-name)))
)
