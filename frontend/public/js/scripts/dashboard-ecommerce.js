/*
 * Dashboard - eCommerce
 */

//Line chart with color shadow: Revenue for 2017 Chart
var thisYearCTX = document.getElementById("thisYearRevenue").getContext("2d");
var lastYearCTX = document.getElementById("lastYearRevenue").getContext("2d");

Chart.types.Line.extend({
  name: "LineAlt",
  initialize: function() {
    Chart.types.Line.prototype.initialize.apply(this, arguments);

    var ctx = this.chart.ctx;
    var originalStroke = ctx.stroke;
    ctx.stroke = function() {
      ctx.save();
      ctx.shadowColor = 'rgba(156, 46, 157,0.5)';
      ctx.shadowBlur = 20;
      ctx.shadowOffsetX = 2;
      ctx.shadowOffsetY = 20;
      originalStroke.apply(this, arguments)
      ctx.restore();
    }
  }
});

Chart.types.Line.extend({
  name: "LineAlt2",
  initialize: function() {
    Chart.types.Line.prototype.initialize.apply(this, arguments);
    var ctx = this.chart.ctx;
    var originalStroke = ctx.stroke;
    ctx.stroke = function() {
      ctx.save();
      originalStroke.apply(this, arguments)
      ctx.restore();
    }
  }
});

var thisYearData = {
  labels: ["January", "February", "March", "April", "May", "June"],
  datasets: [{
    label: "This year dataset",
    fillColor: "#9C2E9D",
    strokeColor: "#9C2E9D",
    pointColor: "transparent",
    pointStrokeColor: "transparent",
    pointHighlightFill: "#fff",
    pointHighlightStroke: "#9C2E9D",
    data: [45, 62, 15, 78, 58, 98]
  }]
};

var lastYearData = {
  labels: ["January", "February", "March", "April", "May", "June"],
  datasets: [{
    label: "Last year dataset",
    fillColor: "#E4E4E4",
    strokeColor: "#E4E4E4",
    pointColor: "transparent",
    pointStrokeColor: "transparent",
    pointHighlightFill: "#fff",
    pointHighlightStroke: "#E4E4E4",
    data: [12, 6, 35, 58, 38, 68]
  }]
};
var thisYearChart = new Chart(thisYearCTX).LineAlt(thisYearData, {
  datasetFill: false,
  scaleShowGridLines: false,

  datasetStrokeWidth: 5,
  scaleFontColor: '#9e9e9e',
  scaleGridLineColor: '#e4e4e4',
  scaleLineColor: 'transparent',
  scaleOverride: true,
  scaleSteps: 5,
  scaleStepWidth: 20,
  scaleStartValue: 0
});
var lastYearChart = new Chart(lastYearCTX).LineAlt2(lastYearData, {
  datasetFill: false,
  scaleShowVerticalLines: false,
  datasetStrokeWidth: 5,
  scaleFontColor: '#9e9e9e',
  scaleGridLineColor: '#e4e4e4',
  scaleLineColor: 'transparent',
  scaleOverride: true,
  scaleSteps: 5,
  scaleStepWidth: 20,
  scaleStartValue: 0
});

/*
 * Line chart with shadow: Monthly Earning Chart
 */

var thisYearCTX = document.getElementById("monthlyEarning").getContext("2d");

Chart.types.Line.extend({
  name: "LineSm",
  initialize: function() {
    Chart.types.Line.prototype.initialize.apply(this, arguments);
    var ctx = this.chart.ctx;
    var originalStroke = ctx.stroke;
    ctx.stroke = function() {
      ctx.save();
      ctx.shadowColor = 'rgba(158, 158, 158, 0.6)';
      ctx.shadowBlur = 14;
      ctx.shadowOffsetX = 2;
      ctx.shadowOffsetY = 16;
      originalStroke.apply(this, arguments)
      ctx.restore();
    }
  }
});

var thisYearData = {
  labels: ["1", "2", "3", "4", "5", "6", "7", "8", "8", "8", "8", "8", "8", "8"],
  datasets: [{
    label: "This year dataset",
    fillColor: "#ff6f00",
    strokeColor: "#ff6f00",
    pointColor: "transparent",
    pointStrokeColor: "transparent",
    pointHighlightFill: "#fff",
    pointHighlightStroke: "#ff6f00",
    data: [50, 40, 30, 55, 50, 90, 40, 50, 40, 70, 55, 80, 50, 55],
  }]
};

var thisYearChart = new Chart(thisYearCTX).LineSm(thisYearData, {
  bezierCurve: false,
  datasetFill: false,
  showTooltips: false,
  scaleShowVerticalLines: false,
  scaleShowGridLines: false,
  datasetStrokeWidth: 3,
  scaleFontColor: 'transparent',
  scaleGridLineColor: 'transparent',
  scaleLineColor: 'transparent',
  scaleOverride: true,
  scaleSteps: 5,
  scaleStepWidth: 20,
  scaleStartValue: 0
});



//Sampel Line Chart One


var ctx = document.getElementById("custom-line-chart-sample-one").getContext("2d");

var gradient = ctx.createLinearGradient(300, 0, 0, 300);
gradient.addColorStop(0, '#0288d1');
gradient.addColorStop(1, '#26c6da');

var options2 = {
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
gradient.addColorStop(0, '#ff6f00');
gradient.addColorStop(1, '#ffca28');

var CustomLineChartSampleThreeData = {
  scaleShowVerticalLines: "false",
  labels: ["January", "February", "March", "April", "May", "June", ],
  datasets: [{
    label: "My Second dataset",
    fillColor: gradient,
    strokeColor: "rgba(255, 202, 40,0.1)",
    pointColor: "#fff",
    pointStrokeColor: "rgba(255, 202, 40,0.9)",
    pointHighlightFill: "#fff",
    pointHighlightStroke: "rgba(255, 202, 40,0.1)",
    data: [24, 18, 20, 30, 40, 43, ]
  }]
};

$(window).on('load', function() {	
  window.CustomLineChartSampleOne = new Chart(document.getElementById("custom-line-chart-sample-one").getContext("2d")).Line(CustomLineChartSampleOneData, options2);
  window.CustomLineChartSampleTwo = new Chart(document.getElementById("custom-line-chart-sample-two").getContext("2d")).Line(CustomLineChartSampleTwoData, options2);
  window.CustomLineChartSampleThree = new Chart(document.getElementById("custom-line-chart-sample-three").getContext("2d")).Line(CustomLineChartSampleThreeData, options2);
});