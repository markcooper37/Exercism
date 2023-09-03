(ns matching-brackets)

(def pairs {\] \[ \} \{ \) \(})

(defn valid?
  ([string] (valid? string ()))
  ([string stack]
   (let [[first & remaining] string
         [top & leftover] stack]
   (cond (empty? string) (empty? stack)
     (or (= first \[) (= first \{) (= first \())
     (valid? remaining (conj stack first))
     (or (= first \]) (= first \}) (= first \)))
     (cond (empty? stack) false
       (= (get pairs first) top) (valid? remaining leftover)
       :else false)
     :else (valid? remaining stack))))
)
