//Based on https://github.com/iamale/MapboxMap
var React = require('react'),
_ = require('lodash'),
$ = require('jquery');

var MapboxMap = React.createClass({
  loadCoordinatesFromServer: function(map, L) {
    $.ajax({
      url: this.props.url,
      dataType: 'json',
      cache: false,
      success: function(coordinates) {
        if (coordinates.length > 0) {
          coordinates = coordinates.map(function(coordinate) {
            return L.latLng(coordinate.x, coordinate.y);
          });
          if ($.isEmptyObject(this.state.polyline)) {
            var polyline  = L.polyline([]).addTo(map);
            this.setState({polyline: polyline });
          }
          this.state.polyline.setLatLngs(coordinates);
          //map.setView([coordinates[coordinates.length -1].lat,coordinates[coordinates.length -1].lng], 3);
        }
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },
  getInitialState: function() {
    return {coordinates: [], polyline: {}};
  },
  componentDidMount: function(argument) {
    var props = this.props;
    var mapId = props.mapId || props.src || "mapbox.streets";
    var options = {};
    var ownProps = ['mapId','pollInterval','url'];
    for (var k in props) {
      if (props.hasOwnProperty(k) && ownProps.indexOf(k) === -1) {
        options[k] = props[k];
      }
    }
    var map = L.mapbox.map(this.getDOMNode(), mapId, options);
    var that = this;
    this.loadCoordinatesFromServer(map, L);
    setInterval(function(){
      that.loadCoordinatesFromServer(map,L);
    }, this.props.pollInterval);
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
