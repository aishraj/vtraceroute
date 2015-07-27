var React = require('react'),
    MapboxMap = require('./MapboxMap');

var App = React.createClass({
  render: function() {
    /*
     * Attributes:
     * - `mapId` (or `src` -- for semantics)
     *        -- that "id or url or tilejson" parameter of L.mapbox.map
     * - `onMapCreated(map, L)`
     *        -- a special callback allowing you to further set up your map
     * - any of the L.mapbox.map and L.Map options, see
     *   https://www.mapbox.com/mapbox.js/api/v2.1.6/l-mapbox-map/ and
     *   https://www.mapbox.com/mapbox.js/api/v2.1.6/l-map-class/
     *   for more info.
     */
     var onMapCreated = function(map, L, dataUrl) {
       var polyline = L.polyline([]).addTo(map);
       var pointsAdded = 0;
       var add = function(that) {

          // `addLatLng` takes a new latLng coordinate and puts it at the end of the
          // line. You optionally pull points from your data or generate them. Here
          // we make a sine wave with some math.
          polyline.addLatLng(
              L.latLng(
                  Math.cos(pointsAdded / 20) * 30,
                  pointsAdded));
          console.log("the dataUrl insides is " + that.dataUrl);
          // Pan the map along with where the line is being added.
          map.setView([0, pointsAdded], 3);

          // Continue to draw and pan the map by calling `add()`
          // until `pointsAdded` reaches 360.
          if (++pointsAdded < 360) window.setTimeout( function() {
            add(that)
          }, 100);
        };
        add(this);
     };




    return (
      <div className="container">
        <MapboxMap
          mapId="mapbox.streets"
          dataUrl="/api/v1/places.json"
          onMapCreated={onMapCreated}
          zoomControl={true}
          center={[0, 0]} zoom={3}/>
      </div>
    );
  }
});

App.start = function () {
    React.render(<App/>, document.getElementById('content'));
};

module.exports = window.App = App;
