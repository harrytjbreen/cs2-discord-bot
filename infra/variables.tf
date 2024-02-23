variable "cs2_maps" {
  // I know that prem is not a "Map" but it works for this model :)
  type = list(string)
  default = [
    "mirage",
    "inferno",
    "nuke",
    "overpass",
    "vertigo",
    "dust2",
    "ancient",
    "anubis",
    "premier",
  ]
}