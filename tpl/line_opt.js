function setLineData(data) {
    return {
        title: {
            text: "lineData.title"
        },
        grid: {
            left: '40',
            right: '40',
            show: true,
        },
        tooltip: {},
        legend: {
            data: ["lineData.title"]
        },
        xAxis: {
            data: data
        },
        yAxis: {
            "scale": true,
        },
        series: [{
            name: '-',
            // type: 'line',
            // type: 'scatter',
            type: 'line',
            barWidth: 2,
            data: data,
            markPoint: {
                data: [{
                    type: 'max'
                },
                {
                    type: 'min'
                }]
            },
            markLine: {
                symbol: 'none',
                data: [
                    { type: 'average', name: '平均值' }
                ],
                lineStyle: {
                    color: "green"
                }
            }
        }]
    };
}

