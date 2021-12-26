function setTimeAxisData(data) {
  var timeAxisOption = {
    tooltip: {
      trigger: 'axis',
      position: function (pt) {
        return [pt[0], '10%'];
      }
    },
    title: {
      left: 'center',
      text: 'data chart'
    },
    toolbox: {
      feature: {
        dataZoom: {
          yAxisIndex: 'none'
        },
        restore: {},
        saveAsImage: {}
      }
    },
    xAxis: {
      type: 'time',
      boundaryGap: false
    },
    yAxis: {
      type: 'value',
      boundaryGap: [0, '100%']
    },
    dataZoom: [
      {
        type: 'inside',
        start: 0,
        end: 20
      },
      {
        start: 0,
        end: 20
      }
    ],
    series: [
      {
        name: 'Data',
        type: 'line',
        smooth: true,
        symbol: 'none',
        areaStyle: {},
        data: data
      }
    ]
  };
return timeAxisOption;
}

function setLineData(lineData) {
  return {
      title: {
          text: lineData.title
      },
      grid:{
        left: '40',
        right: '40',
        show: true,
      },
      tooltip:{},
      legend: {
          data:[lineData.title]
      },
      xAxis: {
          data: xa
      },
      yAxis: {
        "scale": true,
      },
      series: [{
          name: '-',
          // type: 'line',
          // type: 'scatter',
          type: 'bar',
          barWidth: 2,
          data: lineData.data,
          markPoint:{
            data: [{
              type:'max'
              },
              {
              type:'min'
              }]
          },
          markLine : {
              symbol: 'none',
  　　　　　　　data : [
  　　　　　　　　  {type : 'average', name: '平均值'}
  　　　　　　],
        lineStyle:{
          color:"green"
        }
  　　　　}
      }]
  };
}

