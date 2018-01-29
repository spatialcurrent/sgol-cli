server {
  url = "http://localhost:8002/api/graph"
}

formats {

  format {
    id = "csv"
    extension = "csv"
    aliases = ["csv"]
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
    id = "hcl"
    extension = "hcl"
    aliases = ["hcl"]
  }

  format {
    id = "json"
    extension = "json"
    aliases = ["json"]
  }

  format {
    id = "jsonl"
    extension = "jsonl"
    aliases = ["jsonl"]
  }

  format {
    id = "toml"
    extension = "toml"
    aliases = ["toml"]
  }

  format {
    id = "yaml"
    extension = "yml"
    aliases = ["yml", "yaml"]
  }

}

queries {

  query {
    name = "Select_By_Vertex"
    sgol = "SELECT {{ .Vertex }} OUTPUT entities"
    required = ["Vertex"]
  }

  query {
    name = "Select_Activity_By_Alias"
    sgol = "SELECT $Activity FILTER CollectionContains(aliases, \"{{ .Name | lower }}\") OUTPUT entities"
    required = ["Name"]
  }

  query {
    name = "PointOfInterest_By_PointOfInterestType"
    sgol = "SELECT $PointOfInterestType FILTER CollectionContains(aliases, \"{{ .Name | lower }}\") NAV $PointOfInterest $HasType INPUT LIMIT {{ .Limit }} OUTPUT entities"
    required = ["Name", "Limit"]
  }

  query {
    name = "PointOfInterest_By_Activity"
    sgol = "SELECT $Activity FILTER CollectionContains(aliases, \"{{ .Name | lower }}\") NAV $PointOfInterestType $HasActivity INPUT NAV $PointOfInterest $HasType INPUT LIMIT {{ .Limit }} OUTPUT entities"
    required = ["Name", "Limit"]
  }

}
