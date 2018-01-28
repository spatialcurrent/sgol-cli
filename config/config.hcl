server {
  url = "http://localhost:8002/api/graph"
}

formats {

  format {
    id = "json"
    extension = "json"
    aliases = ["json"]
  }

  format {
    id = "geojson"
    extension = "geojson"
    aliases = ["geojson"]
  }

    format {
      id = "geojsonl"
      extension = "geojsonl"
      aliases = ["geojsonl"]
    }

  format {
    id = "yaml"
    extension = "yml"
    aliases = ["yml", "yaml"]
  }

}

queries {

  query {
    name = "Select_Activity_By_Alias"
    sgol = "SELECT $Activity FILTER CollectionContains(aliases, \"{{ .Name }}\") OUTPUT entities"
    required = ["Name"]
  }

  query {
    name = "PointOfInterest_By_PointOfInterestType"
    sgol = "SELECT $PointOfInterestType FILTER CollectionContains(aliases, \"{{ .Name }}\") NAV $PointOfInterest $HasType INPUT LIMIT {{ .Limit }} OUTPUT entities"
    required = ["Name", "Limit"]
  }

}
