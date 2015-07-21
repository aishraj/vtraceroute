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
    return (
      <div className="container">
        <MapboxMap
          mapId="mapbox.comic"
          zoomControl={false}
          center={[59.907433, 30.299848]} zoom={17}/>
      </div>
    );
  }
});

App.start = function () {
    React.render(<App/>, document.getElementById('content'));
};

module.exports = window.App = App;
