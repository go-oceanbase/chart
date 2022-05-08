function setTimeAxisData(data) {
  var timeAxisOption = {
    tooltip: {
      trigger: 'axis',
      position: function (pt) {
        return [pt[0], '50%'];
      }
    },
    title: {
      left: 'center',
      text: 'chart'
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
      type: 'category',
      boundaryGap: false
    },
    yAxis: {
      type: 'value',
      boundaryGap: [0, '10%'],
      splitLine: {
        show: false
      }
    },
    dataZoom: [
      {
        type: 'inside',
        start: 0,
        end: 100
      },
      {
        start: 0,
        end: 100
      }
    ],
    series: [
      {
        name: 'Data',
        type: 'line',
        // smooth: true,
        showSymbol: false,
        data: data
      }
    ]
  };
  return timeAxisOption;
}
