(ns difference-of-squares)

(defn sum-of-squares [n]
  (/ (* (* n (+ n 1)) (+ (* 2 n) 1)) 6))

(defn square-of-sum [n]
  (/ (* (* n n) (* (+ n 1) (+ n 1))) 4))

(defn difference [n]
  (- (square-of-sum n) (sum-of-squares n)))
