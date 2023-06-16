(ns cars-assemble)

(defn production-rate
  "Returns the assembly line's production rate per hour,
   taking into account its success rate"
  [speed]
  (def cars (* speed 221.0))
  (cond (< speed 5) cars 
        (< speed 9) (* cars 0.9)
        (< speed 10) (* cars 0.8)
        :else (* cars 0.77)
    )
  )

(defn working-items
  "Calculates how many working cars are produced per minute"
  [speed]
  (int (/ (production-rate speed) 60))
  )
