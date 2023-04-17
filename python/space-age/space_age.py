class SpaceAge:
    def __init__(self, seconds):
        self.years = seconds / 31557600

    def on_mercury(age):
        return round(age.years / 0.2408467, 2)
    
    def on_venus(age):
        return round(age.years / 0.61519726, 2)
    
    def on_earth(age):
        return round(age.years, 2)

    def on_mars(age):
        return round(age.years / 1.8808158, 2)

    def on_jupiter(age):
        return round(age.years / 11.862615, 2)

    def on_saturn(age):
        return round(age.years / 29.447498, 2)

    def on_uranus(age):
        return round(age.years / 84.016846, 2)

    def on_neptune(age):
        return round(age.years / 164.79132, 2)
