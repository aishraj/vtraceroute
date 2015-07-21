var React = require('react');
// Assuming Mapbox/Leaflet is already exposed as `L`

var MapboxMap = React.createClass({
  componentDidMount: function(argument) {
    var props = this.props;

    var mapId = props.mapId || props.src || "mapbox.streets";

    var options = {};
    var ownProps = ['mapId', 'onMapCreated'];
    for (var k in props) {
      if (props.hasOwnProperty(k) && ownProps.indexOf(k) === -1) {
        options[k] = props[k];
      }
    }

    var map = L.mapbox.map(this.getDOMNode(), mapId, options);

    if (this.props.onMapCreated) {
      this.props.onMapCreated(map, L);
    }
  },

  render: function() {
    var mapStyle = {
      width: '100%',
      height: '100%'
    };

    return (
      <div style={mapStyle}></div>
    );
  }
});

module.exports = MapboxMap;
