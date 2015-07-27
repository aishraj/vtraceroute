var React = require('react'),
MapboxMap = require('./MapboxMap');

var App = React.createClass({
  render: function() {
    return (
      <div className="container">
        <MapboxMap
          url="/api/v1/places.json"
          accessToken="pk.eyJ1IjoiYXJkYWhhbCIsImEiOiJseFQyTWlrIn0.zX_ANNp_k20-iC-6VmbilA"
          zoomControl={true}
          pollInterval={25}
          center={[0, 0]} zoom={4}/>
      </div>
    );
  }
});

App.start = function () {
  React.render(<App/>, document.getElementById('content'));
};

module.exports = window.App = App;
