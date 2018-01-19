package space

const (
  Earth = 31557600
  Mercury = 0.2408467
  Venus = 0.61519726
  Mars = 1.8808158
  Jupiter = 11.862615
  Saturn = 29.447498
  Uranus = 84.016846
  Neptune = 164.79132

)
type Planet string
func Age (sec float64 , name Planet ) float64 {

var planet float64

  switch name {
  case "Earth": planet = Earth

  case "Mercury": planet = Mercury * Earth

  case "Venus": planet = Venus * Earth

  case "Mars": planet = Mars * Earth

  case "Jupiter": planet = Jupiter * Earth

  case "Saturn": planet = Saturn * Earth

  case "Uranus": planet = Uranus * Earth

  case "Neptune": planet = Neptune * Earth

  }

  year := sec / planet

  return year


}
