/*
* ChartJS - Chart
*/

//Sampel Line Chart 
 var LineChartSampleData = {
  labels: ["January", "February", "March", "April", "May", "June", "July"],
  datasets: [{
   label: "My First dataset",
   fillColor: "rgba(0, 191, 165,0.2)",
   strokeColor: "rgba(0, 191, 165,1)",
   pointColor: "rgba(0, 191, 165,1)",
   pointStrokeColor: "#fff",
   pointHighlightFill: "#fff",
   pointHighlightStroke: "rgba(0, 191, 165,1)",
   data: [65, 30, 50, 81, 56, 55, 40]
  }, {
   label: "My Second dataset",
   fillColor: "rgba(255, 110, 64,0.2)",
   strokeColor: "rgba(255, 110, 64,1)",
   pointColor: "rgba(255, 110, 64,1)",
   pointStrokeColor: "#fff",
   pointHighlightFill: "#fff",
   pointHighlightStroke: "rgba(255, 110, 64,1)",
   data: [28, 90, 40, 19, 86, 27, 90]
  }]
 };


 
 //Sampel Bar Chart
 var BarChartSampleData = {
    labels: ["January", "February", "March", "April", "May", "June", "July"],
    datasets: [
        {
            label: "My First dataset",
            fillColor: "rgba(0, 191, 165,0.5)",
            strokeColor: "rgba(0, 191, 165,0.8)",
            highlightFill: "rgba(0, 191, 165,0.75)",
            highlightStroke: "rgba(0, 191, 165,1)",
            data: [65, 59, 80, 81, 56, 55, 40]
        },
        {
            label: "My Second dataset",
            fillColor: "rgba(255, 110, 64,0.5)",
            strokeColor: "rgba(255, 110, 64,0.8)",
            highlightFill: "rgba(255, 110, 64,0.75)",
            highlightStroke: "rgba(255, 110, 64,1)",
            data: [28, 48, 40, 19, 86, 27, 90]
        }
    ]
};

//Sampel Radar Chart
var RadarChartSampleData = {
    labels: ["Eating", "Drinking", "Sleeping", "Designing", "Coding", "Cycling", "Running"],
    datasets: [
        {
            label: "My First dataset",
            fillColor: "rgba(0, 191, 165,0.2)",
            strokeColor: "rgba(0, 191, 165,1)",
            pointColor: "rgba(0, 191, 165,1)",
            pointStrokeColor: "#fff",
            pointHighlightFill: "#fff",
            pointHighlightStroke: "rgba(0, 191, 165,1)",
            data: [65, 59, 90, 81, 56, 55, 40]
        },
        {
            label: "My Second dataset",
            fillColor: "rgba(255, 110, 64,0.2)",
            strokeColor: "rgba(255, 110, 64,1)",
            pointColor: "rgba(255, 110, 64,1)",
            pointStrokeColor: "#fff",
            pointHighlightFill: "#fff",
            pointHighlightStroke: "rgba(255, 110, 64,1)",
            data: [28, 48, 40, 19, 96, 27, 100]
        }
    ]
};

//Sampel Polor Chart
var PolarChartSampleData = [
    {
        value: 300,
        color:"#FF5252",
        highlight: "#E53935",
        label: "Red"
    },
    {
        value: 50,
        color: "#00BFA5",
        highlight: "#00897b",
        label: "Green"
    },
    {
        value: 100,
        color: "#FF6E40",
        highlight: "#F4511E",
        label: "Orange"
    },
    {
        value: 40,
        color: "#949FB1",
        highlight: "#A8B3C5",
        label: "Grey"
    },
    {
        value: 120,
        color: "#00BCD4",
        highlight: "#00ACC1",
        label: "Cyan"
    }

];

//Sampel Pie Doughnut Chart
var PieDoughnutChartSampleData = [
    {
        value: 300,
        color:"#FF5252",
        highlight: "#E53935",
        label: "Red"
    },
    {
        value: 50,
        color: "#00BFA5",
        highlight: "#00897b",
        label: "Green"
    },
    {
        value: 100,
        color: "#FF6E40",
        highlight: "#F4511E",
        label: "Orange"
    }
]

 window.onload = function() {

  window.LineChartSample = new Chart(document.getElementById("line-chart-sample").getContext("2d")).Line(LineChartSampleData,{
   responsive:true
  });
  
  window.BarChartSample = new Chart(document.getElementById("bar-chart-sample").getContext("2d")).Bar(BarChartSampleData,{
   responsive:true
  });
  
  window.RadarChartSample = new Chart(document.getElementById("radar-chart-sample").getContext("2d")).Radar(RadarChartSampleData,{
   responsive:true
  });
  
  window.PolarChartSample = new Chart(document.getElementById("polar-chart-sample").getContext("2d")).PolarArea(PolarChartSampleData,{
   responsive:true
  });
  
  window.PieChartSample = new Chart(document.getElementById("pie-chart-sample").getContext("2d")).Pie(PieDoughnutChartSampleData,{
   responsive:true
  });
  window.DoughnutChartSample = new Chart(document.getElementById("doughnut-chart-sample").getContext("2d")).Pie(PieDoughnutChartSampleData,{
   responsive:true
  });

 };
 