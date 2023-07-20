(ns strain)

(defn retain [pred coll]
      (if (empty? coll)
        []
          (let [head (first coll) tail (rest coll)]
            (if (pred head) (cons head (retain pred tail)) (retain pred tail))))
)

(defn discard [pred coll]
      (if (empty? coll)
        []
          (let [head (first coll) tail (rest coll)]
            (if (pred head) (discard pred tail) (cons head (discard pred tail)))))
)
