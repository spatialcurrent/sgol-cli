# sgol-cli

# Description

**sgol-cli** is a simplified program for interacting with a SGOL server.

# Building

```
cd; go build github.com/spatialcurrent/sgol-cli/cmd/sgol
```

# Usage

You need to set 2 environment variables to use **sgol-cli**.

```
export SGOL_CONFIG_PATH=/path/to/sgol.hcl
export SGOL_AUTH_TOKEN=VALUE_OF_AUTH_TOKEN
```

**Queries**

You can use the **-q** flag to send a SGOL query to the server.

```
./sgol exec -f geojson -q "SELECT $PointOfInterest LIMIT 1 OUTPUT entities"
```

**Named Queries**

For simplicity you can use named queries and then pass variable values by using NAME=VALUE args following the flags.  See examples below.

# Examples

**Select_By_Vertex**

The following calls the named query **Select_By_Vertex** with a GeoJSON reponse format.

```shell
vagrant@vagrant:~$ ./sgol exec -f geojson -named_query Select_By_Vertex Vertex=pointofinterest_11b5cb70eb6e96270961521b71f6abe7
{"features": [{"geometry": {"coordinates": [-77.0431886, 38.9106393], "type": "Point"}, "geometry_name": "the_geom", "id": "pointofinterest_11b5cb70eb6e96270961521b71f6abe7", "properties": {"amenity": "restaurant", "cuisine": "french", "flags": [], "id": "pointofinterest_11b5cb70eb6e96270961521b71f6abe7", "name": "Cafe Dupont", "related": [], "star": false}, "type": "Feature"}], "totalFeatures": 1, "type": "FeatureCollection"}
```
**PointOfInterest_By_PointOfInterestType**

```
vagrant@vagrant:~$ ./sgol exec -f yaml -named_query PointOfInterest_By_PointOfInterestType Name=cafe Limit=1
query: SELECT $PointOfInterestType FILTER CollectionContains(aliases, "cafe") NAV $PointOfInterest $HasType INPUT LIMIT 1 OUTPUT entities
results:
- address: []
  attributes:
  - name: addr:housenumber
    value: '2459'
  - name: amenity
    value: cafe
  - name: name
    value: Tryst Coffeehouse
  - name: wheelchair
    value: 'yes'
  - name: dataset
    value: ABRA locations
  - name: source
    value: dcgis
  - name: addr:street
    value: 18th Street Northwest
  bbox_east: -77.0421524
  bbox_north: 38.9219571
  bbox_south: 38.9219571
  bbox_west: -77.0421524
  bbox_wkt: POINT (-77.0421524 38.9219571)
  flags: []
  geom_lat: 38.9219571
  geom_lon: -77.0421524
  geom_wkt: POINT (-77.0421524 38.9219571)
  id: pointofinterest_0b5065ac826c59179042079a791a110d
  related: []
success: true
```

**PointOfInterest_By_Activity**

The following calls the named_query **PointOfInterest_By_Activity** with a GeoJSON Lines response format.

```shell
vagrant@vagrant:~$ ./sgol exec -f geojsonl -named_query PointOfInterest_By_Activity Name=lunch Limit=2
{"geometry": {"coordinates": [-77.0398398, 38.9070318], "type": "Point"}, "geometry_name": "the_geom", "id": "pointofinterest_0554f718c1e3d28df407d8d174a61c83", "properties": {"addr:housenumber": "1734", "addr:street": "N Street Northwest", "amenity": "restaurant", "dataset": "ABRA locations", "flags": [], "id": "pointofinterest_0554f718c1e3d28df407d8d174a61c83", "name": "Iron Gate Inn Restaurant", "related": [], "source": "dcgis", "star": false}, "type": "Feature"}
{"geometry": {"coordinates": [-77.038713, 38.9123367], "type": "Point"}, "geometry_name": "the_geom", "id": "pointofinterest_05ab718b332651ac2147729ab1ce7421", "properties": {"amenity": "restaurant", "cuisine": "american", "flags": [], "id": "pointofinterest_05ab718b332651ac2147729ab1ce7421", "name": "Brick Lane", "phone": "+1-202-525-5309", "related": [], "star": false, "website": "http://bricklane-restaurant-dc.com/"}, "type": "Feature"}
```

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/sgol-cli/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
