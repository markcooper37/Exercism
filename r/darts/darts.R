score <- function(x, y) {
  distance <- sqrt(x^2 + y^2)
  if (distance <= 1) {
    return(10)
  } else if (distance <= 5) {
    return(5)
  } else if (distance <= 10) {
    return(1)
  }
  return(0)
}
