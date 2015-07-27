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
      success: function(coordinate) {
        if (this.state.coordinate.x === 0 && this.state.coordinate.y === 0 && $.isEmptyObject(this.state.polyline)) {
          var polyline  = L.polyline([]).addTo(map);
          polyline.addLatLng(L.latLng(coordinate.x,coordinate.y));
          this.setState({polyline: polyline });
        }
        if (!_.isEqual(this.state.coordinate, coordinate)) {
          this.setState({coordinate: coordinate});
          if (!$.isEmptyObject(this.state.polyline)) {
            this.state.polyline.addLatLng(L.latLng(this.state.coordinate.x,this.state.coordinate.y));
          }
          map.setView([coordinate.x,coordinate.y], 3);
        }
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },
  getInitialState: function() {
    return {coordinate: {x: 0, y:0}, polyline: {}};
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
