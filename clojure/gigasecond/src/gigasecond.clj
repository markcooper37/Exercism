(ns gigasecond)

(defn from [year month day]  
   (let [date (.plusDays (java.time.LocalDate/of year month day)(/ 1000000000 (* 3600 24)))] 
      [(.getYear date) (.getMonthValue date) (.getDayOfMonth date)]))
