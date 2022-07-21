/*
* Card Extended - Cards
*/ 

//Sampel Line Chart One
var ctx = document.getElementById("custom-line-chart-sample-one").getContext("2d");

var gradient = ctx.createLinearGradient(300, 0, 0, 300);
gradient.addColorStop(0, '#2962FF');
gradient.addColorStop(1, '#3949ab');

var options = {
  scaleFontColor: "#fff",
  scaleShowLabels: false,
  animation: false,
  bezierCurve: true,
  scaleStartValue: 0,
  responsive: true,
  maintainAspectRatio: true,
  scaleShowVerticalLines: false,
  scaleShowHorizontalLines: false,
  showScale: false,
};

var CustomLineChartSampleOneData = {
  scaleShowVerticalLines: "false",
  labels: ["January", "February", "March", "April", "May", "June", ],
  datasets: [{
    label: "My Second dataset",
    fillColor: gradient,
    strokeColor: "rgba(255,82,82,0.1)",
    pointColor: "#fff",
    pointStrokeColor: "#0288d1",
    pointHighlightFill: "#fff",
    pointHighlightStroke: "#0288d1",
    data: [24, 18, 20, 30, 40, 43, ]
  }]
};


//Sampel Line Chart Two


var ctx = document.getElementById("custom-line-chart-sample-two").getContext("2d");

var gradient = ctx.createLinearGradient(500, 0, 0, 200);
gradient.addColorStop(0, '#8e24aa');
gradient.addColorStop(1, '#ff6e40');

var CustomLineChartSampleTwoData = {
  scaleShowVerticalLines: "false",
  labels: ["January", "February", "March", "April", "May", "June", ],
  datasets: [{
    label: "My Second dataset",
    fillColor: gradient,
    strokeColor: "rgba(255,160,0,0.1)",
    pointColor: "#fff",
    pointStrokeColor: "rgba(255,160,0,0.9)",
    pointHighlightFill: "#fff",
    pointHighlightStroke: "rgba(255,160,0,0.1)",
    data: [24, 18, 20, 30, 40, 43, ]
  }]
};

//Sampel Line Chart Three

var ctx = document.getElementById("custom-line-chart-sample-three").getContext("2d");

var gradient = ctx.createLinearGradient(500, 0, 0, 200);
gradient.addColorStop(0, '#7b1fa2');
gradient.addColorStop(1, '#7c4dff');

var CustomLineChartSampleThreeData = {
  scaleShowVerticalLines: "false",
  labels: ["January", "February", "March", "April", "May", "June", ],
  datasets: [{
    label: "My Second dataset",
    fillColor: gradient,
    strokeColor: "rgba(67,160,71,0.1)",
    pointColor: "#fff",
    pointStrokeColor: "rgba(67,160,71,0.9)",
    pointHighlightFill: "#fff",
    pointHighlightStroke: "rgba(67,160,71,0.1)",
    data: [24, 18, 20, 30, 40, 43, ]
  }]
};


window.onload = function() {

  window.CustomLineChartSampleOne = new Chart(document.getElementById("custom-line-chart-sample-one").getContext("2d")).Line(CustomLineChartSampleOneData, options, );
  window.CustomLineChartSampleTwo = new Chart(document.getElementById("custom-line-chart-sample-two").getContext("2d")).Line(CustomLineChartSampleTwoData, options, );
  window.CustomLineChartSampleThree = new Chart(document.getElementById("custom-line-chart-sample-three").getContext("2d")).Line(CustomLineChartSampleThreeData, options, );


};