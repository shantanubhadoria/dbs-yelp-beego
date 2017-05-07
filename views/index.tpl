
<!-- index.html -->
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8" />
	<title>DBS Yelp</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" integrity="sha512-dTfge/zgoMYpP7QbHy4gWMEGsbsdZeCXz7irItjcC3sPUFtf0kuFbDz/ixG7ArTxmDjLXDmezHubeNikyKGVyQ==" crossorigin="anonymous">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/marked/0.3.2/marked.min.js"></script>
	<script src="/static/js/jquery.cookie.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js" integrity="sha512-K1qjQ+NcF2TYO/eI3M6v8EiNYZfA95pQumfvcVrTHtwQVDG+aHRqLi/ETn2uB+1JqwYqVG3LIvdm9lj6imS/pQ==" crossorigin="anonymous"></script>
</head>
<body>
  <script src="https://maps.googleapis.com/maps/api/js?key=AIzaSyD_dOfx8jU7nyFPkjWrSltNIdzHV-2fVQU&libraries=places&callback=initMap"
        async defer></script>

	<div class="container">
    <nav class="row" style="background-color:#d32323;color:#ffffff">
      <div class="col-lg-12">
        <h1>DBS Yelp</h1>
      </div>
    </nav>
    <div id="form-container" style="margin-top:2em">
      <form class="form-inline" id="search-form">
        <div class="form-group">
          <label class="sr-only" for="search-input">Find</label>
          <div class="input-group">
            <div class="input-group-addon">Find</div>
            <input type="text" class="form-control" id="search-input" placeholder="Search ...">
          </div>
        </div>
        <div class="form-group">
          <label class="sr-only" for="exampleInputAmount">Amount (in dollars)</label>
          <div class="input-group">
            <div class="input-group-addon">Location</div>
            <input type="text" class="form-control" id="location-input" placeholder="Enter Location ...">
          </div>
        </div>
        <button type="submit" class="btn btn-primary">Search</button>
      </form>
    </div>
    <div class="row">
      <div id="restaurants-container" class="col-lg-12" style="margin-top:1em">
      </div>
    </div>
  </div>
  <script>
  var mylocation;
  $(document).ready(function() {
    $("#search-form").submit(submitForm);
  });

  function initMap() {
    var input = document.getElementById('location-input');


    mylocation = new google.maps.places.Autocomplete(input);

    // Bind the map's bounds (viewport) property to the autocomplete object,
    // so that the autocomplete requests use the current map bounds for the
    // bounds option in the request.

    mylocation.addListener('place_changed', function() {
      console.log();
    });
  }

  function submitForm(e) {
    e.preventDefault();
    var keyword = $('#search-input').val();
    var place = mylocation.getPlace();
    var latitude = place.geometry.location.lat();
    var longitude = place.geometry.location.lng();
    var url = "{{ .Baseurl }}" + "v1/" + "restaurant/" + keyword + "/" + latitude + "/" + longitude + "/";
    $.ajax({
      url: url,
      dataType: 'json',
      cache: false,
      success: function(data) {
        var container = $("#restaurants-container");
        container.html("");
        data.Businesses.forEach(function(business) {
          var html = "";
          html += '<div class="row" style="margin-top:1em">'
                + '    <div class="col-lg-3">'
                + '      <img src="' + business.image_url + '" style="max-height:10em"/>'
                + '    </div>'
                + '    <div class="col-lg-9">'
                + '      <h3>' + business.name + '</h3>'
                + '      <strong>rating: ' + business.rating + '</strong>'
                + '      <strong>phone: ' + business.phone + '</strong>'
                + '    </div>'
                + '</div>';
          console.log(html);
          container.append(html);
        });

      }.bind(this),
      error: function(xhr, status, err) {
        console.error(url, status, err.toString());
      }.bind(this)
    });

  }
  </script>
</body>
</html>
