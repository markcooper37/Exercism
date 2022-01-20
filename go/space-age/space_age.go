package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    var earthYear float64 = 31557600
	switch planet {
        case "Mercury":
    		return (seconds / earthYear) / 0.2408467
        case "Venus":
    		return (seconds / earthYear) / 0.61519726
        case "Earth":
    		return (seconds / earthYear)
        case "Mars":
   			return (seconds / earthYear) / 1.8808158
        case "Jupiter":
    		return (seconds / earthYear) / 11.862615
        case "Saturn":
    		return (seconds / earthYear) / 29.447498
        case "Uranus":
    		return (seconds / earthYear) / 84.016846
        case "Neptune":
    		return (seconds / earthYear) / 164.79132
        default:
    		return 0
    }
}
