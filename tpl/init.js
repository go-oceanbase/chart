// 基于准备好的dom，初始化echarts实例
var myChart = echarts.init(document.getElementById('chartDom'));

const render = async () => {
    var curved = await window.curve();
    var obj = jQuery.parseJSON(curved)
    if (obj.type == 'line') {
        SetCurveDataOption(myChart, obj);
    } else {
        SetTimeCurveDataOption(myChart, obj);
    }
};

function SetCurveDataOption(chart, obj) {
    option = setLineData(obj.data)

    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
}


function SetTimeCurveDataOption(chart, obj) {
    option = setTimeAxisData(obj.data)

    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
}

render();